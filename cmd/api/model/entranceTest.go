package model

import (
	"github.com/mag30/project-backend/domain/base"
	"github.com/mag30/project-backend/domain/enum"
)

type (
	EntranceTestCheckingRequest struct {
		Question1 string `json:"question1"`
		Question2 string `json:"question2"`
		Question3 string `json:"question3"`
		Question4 string `json:"question4"`
		Question5 string `json:"question5"`
	}

	CheckingResponse struct {
		base.ResponseOK
		Level enum.Level `json:"level"`
	}
)
