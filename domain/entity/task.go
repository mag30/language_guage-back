package entity

import (
	"github.com/google/uuid"
	"github.com/mag30/project-backend/domain/base"
)

type Task struct {
	base.EntityWithIdKey
	Name   string
	QuizID *uuid.UUID
	Quiz   Quiz

	CorrectAnswer string
}
