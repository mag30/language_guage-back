package service

import (
	"context"
	"github.com/google/uuid"
	"github.com/mag30/project-backend/cmd/api/model"
	"github.com/mag30/project-backend/cmd/storage/dao"
	"github.com/mag30/project-backend/domain/base"
	"github.com/mag30/project-backend/domain/enum"
)

type EntranceTestService struct {
	userStorage *dao.UserStorage
}

func NewEntranceTestService(userStorage *dao.UserStorage) *EntranceTestService {
	return &EntranceTestService{
		userStorage: userStorage,
	}
}

func (s EntranceTestService) Checking(userID uuid.UUID, request model.EntranceTestCheckingRequest, ctx context.Context) (enum.Level, *base.ServiceError) {
	user, err := s.userStorage.Retrieve(userID, ctx)
	if err != nil {
		return enum.None, base.NewPostgresReadError(err)
	}

	newLevel := enum.None

	counter := 0

	if request.Question1 == "I’m from France." {
		counter++
	}

	if request.Question2 == "The largest island is Great Britain." {
		counter++
	}

	if request.Question3 == "am getting used" {
		counter++
	}

	if request.Question4 == "is the manager's office" {
		counter++
	}

	if request.Question5 == "heated" {
		counter++
	}

	if counter <= 2 {

	}

	if counter <= 2 {
		newLevel = enum.Beginner
	} else if counter <= 4 {
		newLevel = enum.Elementary
	} else {
		newLevel = enum.Intermediate
	}

	if user.Level != enum.None {
		switch user.Level {
		case enum.Beginner:
			if newLevel == enum.Beginner {
				return enum.None, nil
			}

		case enum.Elementary:
			if newLevel == enum.Beginner || newLevel == enum.Elementary {
				return enum.None, nil
			}

		case enum.Intermediate:
			if newLevel == enum.Beginner || newLevel == enum.Elementary || newLevel == enum.Intermediate {
				return enum.None, nil
			}
		}

	}

	user.Level = newLevel

	if err := s.userStorage.Update(user, ctx); err != nil {
		return enum.None, base.NewPostgresWriteError(err)
	}

	return newLevel, nil
}
