package model

import (
	"github.com/mag30/project-backend/domain/base"
	"github.com/mag30/project-backend/domain/entity"
)

type GetResultResponse struct {
	base.ResponseOK
	Result *entity.Result `json:"result"`
}

type CheckTestRequest struct {
	QuizName string `json:"quizId"`
	Answers  struct {
		Q1 string `json:"question1"`
		Q2 string `json:"question2"`
		Q3 string `json:"question3"`
		Q4 string `json:"question4"`
		Q5 string `json:"question5"`
	} `json:"answers"`
}

type CheckTestResponse struct {
	base.ResponseOK
	Q1 TestObject `json:"q1"`
	Q2 TestObject `json:"q2"`
	Q3 TestObject `json:"q3"`
	Q4 TestObject `json:"q4"`
	Q5 TestObject `json:"q5"`
}

type TestObject struct {
	IsRight       bool   `json:"isRight"`
	UserAnswer    string `json:"userAnswer"`
	CorrectAnswer string `json:"correctAnswer"`
}
