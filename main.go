package main

// dummy import to preserve swaggo/swag
// see https://github.com/golang/go/issues/37352
import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/mag30/project-backend/auth"
	"github.com/mag30/project-backend/cmd/api/controller"
	"github.com/mag30/project-backend/cmd/api/router"
	"github.com/mag30/project-backend/cmd/service"
	"github.com/mag30/project-backend/cmd/storage/dao"
	"github.com/mag30/project-backend/cmd/storage/migration"
	"github.com/mag30/project-backend/common"
	"github.com/mag30/project-backend/config"
	"github.com/mag30/project-backend/docs"
	"github.com/mag30/project-backend/server"
	"github.com/spf13/viper"
	_ "github.com/swaggo/swag"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
	"os/signal"
	"syscall"
)

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func main() {
	appCli := common.InitAppCli()
	if err := appCli.Run(os.Args); err != nil {
		panic(err.Error())
	}

	// read config
	var cfg config.Config
	if err := viper.MergeInConfig(); err != nil {
		panic(fmt.Sprintf("error reading config file: %v", err))
	}

	err := viper.Unmarshal(&cfg)
	if err != nil {
		panic(fmt.Sprintf("unable to decode into struct: %v", err))
	}

	// configure swagger
	swaggerConfig := common.NewSwaggerConfig("User API", "TBD")

	docs.SwaggerInfo.Title = swaggerConfig.Title
	docs.SwaggerInfo.Description = swaggerConfig.Description
	docs.SwaggerInfo.Version = swaggerConfig.Version
	docs.SwaggerInfo.BasePath = swaggerConfig.BasePath
	docs.SwaggerInfo.Schemes = swaggerConfig.Schemes

	// init connections
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable",
		cfg.DataBase.Host, cfg.DataBase.User, cfg.DataBase.Password, cfg.DataBase.Name, cfg.DataBase.Port)

	db, err := gorm.Open(postgres.Open(dsn))
	if err != nil {
		panic(fmt.Sprintf("can't connect to database: %v", err))
	}

	hasher := auth.NewHasher(cfg.Auth.Salt)

	adminID, err := uuid.Parse(cfg.AdminMigration.AdminID)
	if err != nil {
		panic(fmt.Sprintf("failed parse uuid admin: %v", err))
	}

	adminPassword, err := hasher.Hash(cfg.AdminMigration.AdminPassword)
	if err != nil {
		panic(fmt.Sprintf("failed hash admin password: %v", err))
	}

	if err := migration.Migration(
		db,
		adminID,
		cfg.AdminMigration.AdminUserName,
		cfg.AdminMigration.AdminEmail,
		adminPassword); err != nil {
		panic(fmt.Sprintf("failed to migrate database: %v", err))
	}

	fmt.Println("database migrated successfully")

	jwtManager, err := auth.NewJWTManager(cfg.Auth.SigningKey, cfg.Auth.TimeToLive)
	if err != nil {
		panic(fmt.Sprintf("failed to create JWTManager: %v", err))
	}

	//init storage
	userStorage := dao.NewUserStorage(db)
	sessionStorage := dao.NewSessionStorage(db)
	resultStorage := dao.NewResultStorage(db)
	taskStorage := dao.NewTaskStorage(db)
	quizStorage := dao.NewQuizStorage(db)

	//init service
	authService := service.NewAuthService(userStorage, sessionStorage, hasher, jwtManager)
	userService := service.NewUserService(userStorage, authService, hasher)
	entranceTestService := service.NewEntranceTestService(userStorage)
	testService := service.NewTestService(quizStorage, taskStorage, resultStorage)

	//init controller
	controllers := controller.NewContainer(authService, userService, entranceTestService, testService)

	handler := router.NewRouter(cfg)
	srv := new(server.Server)

	go func() {
		if err := srv.Run(cfg.Server.Host, cfg.Server.Port, handler.InitRouter(
			controllers,
			jwtManager)); err != nil {
			panic(fmt.Sprintf("error accured while running http server: %s", err.Error()))
		}
	}()

	fmt.Printf("listening server on %s:%s", cfg.Server.Host, cfg.Server.Port)

	// handle signals
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	fmt.Println("shutting down gracefully...")
	defer func() { fmt.Println("shutdown complete") }()

	// perform shutdown
	if err := srv.Shutdown(context.Background()); err != nil {
		panic(fmt.Sprintf("error occured on public server shutting down: %s", err.Error()))
	}
}
