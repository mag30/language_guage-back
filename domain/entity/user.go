package entity

import (
	"github.com/google/uuid"
	"github.com/mag30/project-backend/domain/base"
	"github.com/mag30/project-backend/domain/enum"
)

type User struct {
	base.EntityWithIdKey
	Name     string
	Email    string `gorm:"uniqueIndex"`
	Password string
	Level    enum.Level `gorm:"default:None"`
	Role     enum.Role  `gorm:"default:Guest"`
	Sessions []Session
}

type Session struct {
	base.EntityWithIdKey
	User   *User     `json:"user"`
	UserID uuid.UUID `json:"userID"`
}
