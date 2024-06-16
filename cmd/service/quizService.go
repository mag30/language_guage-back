package service

import (
	"context"
	"github.com/mag30/project-backend/cmd/storage/dao"
	"github.com/mag30/project-backend/domain/base"
	"github.com/mag30/project-backend/domain/entity"
)

type QuizService struct {
	storage dao.QuizStorage
}

func NewQuizService(storage dao.QuizStorage) *QuizService {
	return &QuizService{
		storage: storage,
	}
}

func (s *QuizService) CreateQuiz(ctx context.Context, name string) (*entity.Quiz, *base.ServiceError) {
	quiz := &entity.Quiz{
		Name: name,
	}

	err := s.storage.Create(quiz, ctx)
	if err != nil {
		return nil, base.NewPostgresWriteError(err)
	}

	return quiz, nil
}

func (s *QuizService) GetQuizByName(ctx context.Context, name string) (*entity.Quiz, *base.ServiceError) {
	quiz, err := s.storage.GetByName(name, ctx)
	if err != nil {
		return nil, base.NewPostgresReadError(err)
	}

	return quiz, nil
}
