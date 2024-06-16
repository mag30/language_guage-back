package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/mag30/project-backend/api"
	"github.com/mag30/project-backend/api/middleware"
	"github.com/mag30/project-backend/cmd/api/model"
	"github.com/mag30/project-backend/cmd/service"
	"github.com/mag30/project-backend/domain/base"
	"net/http"
)

type TestController struct {
	testService *service.TestService
}

func NewTestController(testService *service.TestService) *TestController {
	return &TestController{testService: testService}
}

func (a *TestController) GetResult(c *gin.Context) {
	userID, _ := c.Get(middleware.UserIDKey)

	quizName := c.Params.ByName("quiz-name")

	result, serviceErr := a.testService.GetResult(userID.(uuid.UUID), quizName, api.GetExtendedContext(c))
	if serviceErr != nil {
		c.JSON(serviceErr.Code, base.ResponseOK{
			Status:     http.StatusText(http.StatusOK),
			TrackingID: middleware.GetTrackingId(c),
		})
		return
	}

	c.JSON(http.StatusOK, model.GetResultResponse{
		ResponseOK: base.ResponseOK{
			Status:     http.StatusText(http.StatusOK),
			TrackingID: middleware.GetTrackingId(c),
		},
		Result: result,
	})
}

func (a *TestController) CheckTest(c *gin.Context) {
	userID, _ := c.Get(middleware.UserIDKey)

	var payload model.CheckTestRequest
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, base.ResponseFailure{
			Status:     http.StatusText(http.StatusBadRequest),
			Blame:      base.BlameUser,
			Message:    "failed to parse json",
			TrackingID: middleware.GetTrackingId(c),
		})
		return
	}

	result, serviceErr := a.testService.CheckTest(userID.(uuid.UUID), payload, api.GetExtendedContext(c))
	if serviceErr != nil {
		c.JSON(serviceErr.Code, base.ResponseOK{
			Status:     http.StatusText(http.StatusOK),
			TrackingID: middleware.GetTrackingId(c),
		})
		return
	}

	result.Status = http.StatusText(http.StatusOK)
	result.TrackingID = middleware.GetTrackingId(c)
	c.JSON(http.StatusOK, result)
}

func (a *TestController) RestoreTest(c *gin.Context) {
	userID, _ := c.Get(middleware.UserIDKey)

	quizName := c.Params.ByName("quiz-name")

	if serviceErr := a.testService.RestoreTest(userID.(uuid.UUID), quizName, api.GetExtendedContext(c)); serviceErr != nil {
		c.JSON(serviceErr.Code, base.ResponseOK{
			Status:     http.StatusText(http.StatusOK),
			TrackingID: middleware.GetTrackingId(c),
		})
		return
	}

	c.JSON(http.StatusOK, base.ResponseOK{
		Status:     http.StatusText(http.StatusOK),
		TrackingID: middleware.GetTrackingId(c),
	})
}
