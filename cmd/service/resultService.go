package service

import (
	"context"
	"github.com/google/uuid"
	"github.com/mag30/project-backend/cmd/storage/dao"
	"github.com/mag30/project-backend/domain/base"
	"github.com/mag30/project-backend/domain/entity"
)

type ResultService struct {
	storage     *dao.ResultStorage
	quizStorage *dao.QuizStorage
}

func NewResultService(storage *dao.ResultStorage, quizStorage *dao.QuizStorage) *ResultService {
	return &ResultService{
		storage:     storage,
		quizStorage: quizStorage,
	}
}

func (s *ResultService) CreateResult(ctx context.Context, userID uuid.UUID, quizName string, answer string, passed *bool) (*entity.Result, *base.ServiceError) {
	quiz, err := s.quizStorage.GetByName(quizName, ctx)
	if err != nil {
		return nil, base.NewPostgresReadError(err)
	}

	result := &entity.Result{
		UserID: userID,
		QuizID: quiz.ID,
		Answer: answer,
		Passed: passed,
	}

	err = s.storage.Create(result, ctx)
	if err != nil {
		return nil, base.NewPostgresWriteError(err)
	}

	return result, nil
}

func (s *ResultService) GetResultByID(ctx context.Context, id uuid.UUID) (*entity.Result, *base.ServiceError) {
	result, err := s.storage.Retrieve(id, ctx)
	if err != nil {
		return nil, base.NewPostgresReadError(err)
	}

	return result, nil
}
