package service

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/google/uuid"
	"github.com/mag30/project-backend/cmd/api/model"
	"github.com/mag30/project-backend/cmd/storage/dao"
	"github.com/mag30/project-backend/domain/base"
	"github.com/mag30/project-backend/domain/entity"
	"gorm.io/gorm"
)

type TestService struct {
	quizStorage   *dao.QuizStorage
	taskStorage   *dao.TaskStorage
	resultStorage *dao.ResultStorage
}

func NewTestService(quizStorage *dao.QuizStorage,
	taskStorage *dao.TaskStorage,
	resultStorage *dao.ResultStorage) *TestService {
	return &TestService{
		quizStorage:   quizStorage,
		taskStorage:   taskStorage,
		resultStorage: resultStorage,
	}
}

func (s TestService) GetResult(userID uuid.UUID, quidName string, ctx context.Context) (*entity.Result, *base.ServiceError) {
	quiz, err := s.quizStorage.GetByName(quidName, ctx)
	if err != nil {
		return nil, base.NewPostgresReadError(err)
	}

	result, err := s.resultStorage.GetByUserIDAndQuizID(userID, quiz.ID, ctx)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			newResult := &entity.Result{
				UserID: userID,
				QuizID: quiz.ID,
			}

			if err := s.resultStorage.Create(newResult, ctx); err != nil {
				return nil, base.NewPostgresWriteError(err)
			}

			return nil, nil
		}
		return nil, base.NewPostgresReadError(err)
	}

	return result, nil
}

func (s TestService) RestoreTest(userID uuid.UUID, quidName string, ctx context.Context) *base.ServiceError {
	quiz, err := s.quizStorage.GetByName(quidName, ctx)
	if err != nil {
		return base.NewPostgresReadError(err)
	}

	result, err := s.resultStorage.GetByUserIDAndQuizID(userID, quiz.ID, ctx)
	if err != nil {
		return base.NewPostgresReadError(err)
	}

	if err := s.resultStorage.Delete(result.ID, ctx); err != nil {
		return base.NewPostgresReadError(err)
	}

	return nil
}

func (s TestService) CheckTest(userID uuid.UUID, request model.CheckTestRequest, ctx context.Context) (*model.CheckTestResponse, *base.ServiceError) {
	quiz, err := s.quizStorage.GetByName(request.QuizName, ctx)
	if err != nil {
		return nil, base.NewPostgresReadError(err)
	}

	test1, err := s.taskStorage.GetByNameAndQuizID(quiz.ID, "question1", ctx)
	if err != nil {
		return nil, base.NewPostgresReadError(err)
	}
	testObject1 := model.TestObject{
		UserAnswer:    request.Answers.Q1,
		CorrectAnswer: test1.CorrectAnswer,
	}

	if request.Answers.Q1 == test1.CorrectAnswer {
		testObject1.IsRight = true
	} else {
		testObject1.IsRight = false
	}

	test2, err := s.taskStorage.GetByNameAndQuizID(quiz.ID, "question2", ctx)
	if err != nil {
		return nil, base.NewPostgresReadError(err)
	}
	testObject2 := model.TestObject{
		UserAnswer:    request.Answers.Q2,
		CorrectAnswer: test2.CorrectAnswer,
	}

	if request.Answers.Q2 == test2.CorrectAnswer {
		testObject2.IsRight = true
	} else {
		testObject2.IsRight = false
	}

	test3, err := s.taskStorage.GetByNameAndQuizID(quiz.ID, "question3", ctx)
	if err != nil {
		return nil, base.NewPostgresReadError(err)
	}
	testObject3 := model.TestObject{
		UserAnswer:    request.Answers.Q3,
		CorrectAnswer: test3.CorrectAnswer,
	}

	if request.Answers.Q3 == test3.CorrectAnswer {
		testObject3.IsRight = true
	} else {
		testObject3.IsRight = false
	}

	test4, err := s.taskStorage.GetByNameAndQuizID(quiz.ID, "question4", ctx)
	if err != nil {
		return nil, base.NewPostgresReadError(err)
	}
	testObject4 := model.TestObject{
		UserAnswer:    request.Answers.Q4,
		CorrectAnswer: test4.CorrectAnswer,
	}

	if request.Answers.Q4 == test4.CorrectAnswer {
		testObject4.IsRight = true
	} else {
		testObject4.IsRight = false
	}

	test5, err := s.taskStorage.GetByNameAndQuizID(quiz.ID, "question5", ctx)
	if err != nil {
		return nil, base.NewPostgresReadError(err)
	}
	testObject5 := model.TestObject{
		UserAnswer:    request.Answers.Q5,
		CorrectAnswer: test5.CorrectAnswer,
	}

	if request.Answers.Q5 == test5.CorrectAnswer {
		testObject5.IsRight = true
	} else {
		testObject5.IsRight = false
	}

	result, err := s.resultStorage.GetByUserIDAndQuizID(userID, quiz.ID, ctx)
	if err != nil {
		return nil, base.NewPostgresReadError(err)
	}

	CheckTestObject := model.CheckTestResponse{
		Q1: testObject1,
		Q2: testObject2,
		Q3: testObject3,
		Q4: testObject4,
		Q5: testObject5,
	}

	b, err := json.Marshal(CheckTestObject)
	if err != nil {
		return nil, base.NewJsonMarshalError(err)
	}

	t := true
	result.Passed = &t
	result.Answer = string(b)

	if err := s.resultStorage.Update(result, ctx); err != nil {
		return nil, base.NewPostgresWriteError(err)
	}

	return &CheckTestObject, nil
}
