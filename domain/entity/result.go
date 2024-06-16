package entity

import (
	"github.com/google/uuid"
	"github.com/mag30/project-backend/domain/base"
)

type Result struct {
	base.EntityWithIdKey
	UserID uuid.UUID
	User   User
	QuizID uuid.UUID
	Quiz   Quiz
	Passed *bool `gorm:"default:false"`
	Answer string
}
