package service

import (
	"context"
	"github.com/google/uuid"
	"github.com/mag30/project-backend/cmd/storage/dao"
	"github.com/mag30/project-backend/domain/base"
	"github.com/mag30/project-backend/domain/entity"
)

type TaskService struct {
	storage dao.TaskStorage
}

func NewTaskService(storage dao.TaskStorage) *TaskService {
	return &TaskService{
		storage: storage,
	}
}

func (s *TaskService) CreateTask(ctx context.Context, name string, quizID *uuid.UUID, correctAnswer string) (*entity.Task, *base.ServiceError) {
	task := &entity.Task{
		Name:          name,
		QuizID:        quizID,
		CorrectAnswer: correctAnswer,
	}

	err := s.storage.Create(task, ctx)
	if err != nil {
		return nil, base.NewPostgresWriteError(err)
	}

	return task, nil
}

func (s *TaskService) GetTasksByQuizID(ctx context.Context, quizID uuid.UUID) ([]entity.Task, *base.ServiceError) {
	tasks, err := s.storage.GetByQuizID(quizID, ctx)
	if err != nil {
		return nil, base.NewPostgresReadError(err)
	}

	return tasks, nil
}
