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

type EntranceTestController struct {
	entranceTestService *service.EntranceTestService
}

func NewEntranceTestController(entranceTestService *service.EntranceTestService) *EntranceTestController {
	return &EntranceTestController{entranceTestService: entranceTestService}
}

// Checking entrance-api
// @Summary     Start checking the first text
// @Description  Start checking the first text
// @Tags         Entrance
// @Accept       json
// @Produce      json
// @Security ApiKeyAuth
// @param Authorization header string true "Authorization"
// @Success      200  {object}  model.EntranceTestCheckingRequest "OK"
// @Failure      400  {object}  base.ResponseFailure "Bad request"
// @Failure      500  {object}  base.ResponseFailure "Internal error (server fault)"
// @Router       /entranceTest/checking [get]
func (a *EntranceTestController) Checking(c *gin.Context) {
	userID, _ := c.Get(middleware.UserIDKey)

	var payload model.EntranceTestCheckingRequest
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, api.GeneralParsingError(middleware.GetTrackingId(c)))
		return
	}

	newLevel, serviceErr := a.entranceTestService.Checking(userID.(uuid.UUID), payload, api.GetExtendedContext(c))
	if serviceErr != nil {
		c.JSON(serviceErr.Code, api.ResponseFromServiceError(*serviceErr, middleware.GetTrackingId(c)))
		return
	}

	c.JSON(http.StatusOK, model.CheckingResponse{
		ResponseOK: base.ResponseOK{
			Status:     http.StatusText(http.StatusOK),
			TrackingID: middleware.GetTrackingId(c),
		},
		Level: newLevel,
	})
}
