package common

import (
	"net/http"
	"time"

	"github.com/gin-contrib/cors"
)

// ServerConfig configures gin server.
type ServerConfig struct {
	Host             string
	Port             string
	UseAuthorization bool
	UseMetrics       bool

	GinMode string

	Limits     []string
	Operations map[string]string
}

// DatabaseConfig stores db credentials.
type DatabaseConfig struct {
	Host     string
	Port     int
	User     string
	Password string
	Name     string
}

const (
	defaultHost     = "localhost:8080"
	defaultBasePath = "/api"
)

var defaultSchemes = []string{"http", "https"}

// SwaggerConfig configures swaggo/swag.
type SwaggerConfig struct {
	Title       string
	Description string
	Version     string
	Host        string
	BasePath    string
	Schemes     []string
}

// NewSwaggerConfig returns *SwaggerConfig with preconfigured fields.
func NewSwaggerConfig(title, description string) *SwaggerConfig {
	return &SwaggerConfig{
		Title:       title,
		Description: description,
		Version:     "v1",
		Host:        defaultHost,
		BasePath:    defaultBasePath,
		Schemes:     defaultSchemes,
	}
}

// DefaultCorsConfig returns cors.Config with very permissive policy.
func DefaultCorsConfig() cors.Config {
	return cors.Config{
		AllowOrigins:  []string{"*"},
		AllowMethods:  []string{http.MethodGet, http.MethodPost, http.MethodDelete, http.MethodOptions},
		AllowHeaders:  []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders: []string{"Content-Length"},
		MaxAge:        12 * time.Hour,
	}
}

type AuthConfig struct {
	Salt       string
	SigningKey string
	TimeToLive time.Duration
}

type AdminMigrationConfig struct {
	AdminID       string
	AdminUserName string
	AdminEmail    string
	AdminPassword string
}
