package router

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/mag30/project-backend/api/middleware"
	"github.com/mag30/project-backend/auth"
	"github.com/mag30/project-backend/cmd/api/controller"
	"github.com/mag30/project-backend/common"
	"github.com/mag30/project-backend/config"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type Router struct {
	config config.Config
}

func NewRouter(config config.Config) *Router {
	return &Router{
		config: config,
	}
}

func (h *Router) InitRouter(
	controllerContainer *controller.Container,
	JWTManager *auth.JWTManager,
) *gin.Engine {

	router := gin.Default()
	router.Use(middleware.SetTracingContext())
	//router.Use(middleware.SetAccessControl(h.config.Server, *logger))
	router.Use(cors.New(common.DefaultCorsConfig()))

	router.GET("api/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	baseRouter := router.Group("/api")
	user := baseRouter.Group("user")
	{
		user.POST("register", controllerContainer.AuthController.Register)
		user.POST("login", controllerContainer.AuthController.Login)
		user.POST("refresh", controllerContainer.AuthController.RecreateJWT)
		user.POST(
			"logout",
			middleware.SetAuthorizationCheck(JWTManager),
			controllerContainer.AuthController.Logout)

		user.GET(
			"get",
			middleware.SetAuthorizationCheck(JWTManager),
			controllerContainer.UserController.Get)
		user.GET("retrieve", middleware.SetAuthorizationCheck(JWTManager), controllerContainer.UserController.RetrieveUser)
		user.POST(
			"/authorizationFields/update",
			middleware.SetAuthorizationCheck(JWTManager),
			controllerContainer.UserController.UpdateAuthorizationFields)
		user.POST("user/:user-id/update", controllerContainer.UserController.Update)
	}
	entranceTest := baseRouter.Group("entranceTest")
	{
		entranceTest.POST("checking", middleware.SetAuthorizationCheck(JWTManager), controllerContainer.EntranceTestController.Checking)
	}
	test := baseRouter.Group("test")
	{
		test.GET(":quiz-name/get", middleware.SetAuthorizationCheck(JWTManager), controllerContainer.TestController.GetResult)
		test.POST("check", middleware.SetAuthorizationCheck(JWTManager), controllerContainer.TestController.CheckTest)
		test.POST(":quiz-name/restore", middleware.SetAuthorizationCheck(JWTManager), controllerContainer.TestController.RestoreTest)
	}

	return router
}
