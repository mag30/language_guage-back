package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/mag30/project-backend/api"
	"github.com/mag30/project-backend/api/middleware"
	"github.com/mag30/project-backend/cmd/api/model"
	"github.com/mag30/project-backend/cmd/service"
	"github.com/mag30/project-backend/domain/base"
	"net/http"
)

type AuthController struct {
	authService *service.AuthService
}

func NewAuthController(authService *service.AuthService) *AuthController {
	return &AuthController{authService: authService}
}

// Register User registration user-api
// @Summary      User registration
// @Description  User registration
// @Tags         Auth
// @Accept       json
// @Produce      json
// @Param payload body model.RegisterRequest true "User request"
// @Success      200  {object}  base.ResponseOKWithID "OK"
// @Failure      400  {object}  base.ResponseFailure "Bad request"
// @Failure      500  {object}  base.ResponseFailure "Internal error (server fault)"
// @Router       /user/register [post]
func (a *AuthController) Register(c *gin.Context) {
	var payload model.RegisterRequest
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, api.GeneralParsingError(middleware.GetTrackingId(c)))
		return
	}

	id, serviceErr := a.authService.Register(&payload, api.GetExtendedContext(c))
	if serviceErr != nil {
		c.JSON(serviceErr.Code, api.ResponseFromServiceError(*serviceErr, middleware.GetTrackingId(c)))
		return
	}

	c.JSON(http.StatusOK, base.ResponseOKWithID{
		Status:     http.StatusText(http.StatusOK),
		TrackingID: middleware.GetTrackingId(c),
		ID:         *id,
	})
}

// Login User authorisation user-api
// @Summary      User authorisation
// @Description  User authorisation
// @Tags         Auth
// @Accept       json
// @Produce      json
// @Param payload body model.LoginRequest true "User request"
// @Success      200  {object}  model.LoginResponse "OK"
// @Failure      400  {object}  base.ResponseFailure "Bad request"
// @Failure      500  {object}  base.ResponseFailure "Internal error (server fault)"
// @Router       /user/login [post]
func (a *AuthController) Login(c *gin.Context) {
	var payload model.LoginRequest
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, api.GeneralParsingError(middleware.GetTrackingId(c)))
		return
	}

	token, refresh, serviceErr := a.authService.Login(&payload, api.GetExtendedContext(c))
	if serviceErr != nil {
		c.JSON(serviceErr.Code, api.ResponseFromServiceError(*serviceErr, middleware.GetTrackingId(c)))
		return
	}

	c.JSON(http.StatusOK, model.LoginResponse{
		ResponseOK: base.ResponseOK{
			Status:     http.StatusText(http.StatusOK),
			TrackingID: middleware.GetTrackingId(c),
		},
		JWT:          *token,
		RefreshToken: *refresh,
	})
}

// Logout Unauthorized users user-api
// @Summary      Unauthorized users
// @Description  Unauthorized users
// @Tags         Auth
// @Accept       json
// @Produce      json
// @Security ApiKeyAuth
// @param Authorization header string true "Authorization"
// @Param payload body model.RecreateJWTRequest true "User request"
// @Success      200  {object}  base.ResponseOK "OK"
// @Failure      400  {object}  base.ResponseFailure "Bad request"
// @Failure      500  {object}  base.ResponseFailure "Internal error (server fault)"
// @Router       /user/logout [post]
func (a *AuthController) Logout(c *gin.Context) {
	var payload model.RecreateJWTRequest
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, api.GeneralParsingError(middleware.GetTrackingId(c)))
		return
	}

	if serviceErr := a.authService.Logout(payload.RefreshToken, api.GetExtendedContext(c)); serviceErr != nil {
		c.JSON(serviceErr.Code, api.ResponseFromServiceError(*serviceErr, middleware.GetTrackingId(c)))
		return
	}

	c.JSON(http.StatusOK, base.ResponseOK{
		Status:     http.StatusText(http.StatusOK),
		TrackingID: middleware.GetTrackingId(c),
	})
}

// RecreateJWT Re-create refresh token user-api
// @Summary      Re-create refresh token
// @Description  Re-create refresh token
// @Tags         Auth
// @Accept       json
// @Produce      json
// @Param payload body model.RecreateJWTRequest true "User request"
// @Success      200  {object}  model.LoginResponse "OK"
// @Failure      400  {object}  base.ResponseFailure "Bad request"
// @Failure      500  {object}  base.ResponseFailure "Internal error (server fault)"
// @Router       /user/refresh [post]
func (a *AuthController) RecreateJWT(c *gin.Context) {
	var payload model.RecreateJWTRequest
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, api.GeneralParsingError(middleware.GetTrackingId(c)))
		return
	}

	token, newRefresh, _, serviceErr := a.authService.RefreshJWT(payload.RefreshToken, api.GetExtendedContext(c))
	if serviceErr != nil {
		c.JSON(serviceErr.Code, api.ResponseFromServiceError(*serviceErr, middleware.GetTrackingId(c)))
		return
	}

	c.JSON(http.StatusOK, model.LoginResponse{
		ResponseOK: base.ResponseOK{
			Status:     http.StatusText(http.StatusOK),
			TrackingID: middleware.GetTrackingId(c),
		},
		JWT:          *token,
		RefreshToken: *newRefresh,
	})
}
