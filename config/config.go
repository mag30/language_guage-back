package config

import (
	"github.com/mag30/project-backend/common"
)

type Config struct {
	DataBase       common.DatabaseConfig
	Server         common.ServerConfig
	Auth           common.AuthConfig
	AdminMigration common.AdminMigrationConfig
}
