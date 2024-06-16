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

type UserController struct {
	userService *service.UserService
}

func NewUserController(
	userService *service.UserService) *UserController {
	return &UserController{
		userService: userService,
	}
}

// Get user-api
// @Summary      Get all users
// @Description  Get all users
// @Tags         User
// @Accept       json
// @Produce      json
// @Security ApiKeyAuth
// @param Authorization header string true "Authorization"
// @Success      200  {object}  model.GetUsersResponse "OK"
// @Failure      400  {object}  base.ResponseFailure "Bad request"
// @Failure      500  {object}  base.ResponseFailure "Internal error (server fault)"
// @Router       /user/get [get]
func (a *UserController) Get(c *gin.Context) {

	users, serviceErr := a.userService.Get(api.GetExtendedContext(c))
	if serviceErr != nil {
		c.JSON(serviceErr.Code, api.ResponseFromServiceError(*serviceErr, middleware.GetTrackingId(c)))
		return

	}

	c.JSON(http.StatusOK, model.GetUsersResponse{
		ResponseOK: base.ResponseOK{
			Status:     http.StatusText(http.StatusOK),
			TrackingID: middleware.GetTrackingId(c),
		},
		Users: users,
	})
}

// Update
// @Summary      Update User All Fields
// @Description  Update User All Fields
// @Tags         User
// @Accept       json
// @Produce      json
// @Security ApiKeyAuth
// @param Authorization header string true "Authorization"
// @Param        payload body   model.UpdateUserAllFieldRequest true "User data"
// @Success      200  {object}  base.ResponseOK "OK"
// @Failure      400  {object}  base.ResponseFailure "Bad request"
// @Failure      500  {object}  base.ResponseFailure "Internal error (server fault)"
// @Router       /user/:user-id/update [post]
func (a *UserController) Update(c *gin.Context) {
	userID, err := uuid.Parse(c.Params.ByName("user-id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, api.GeneralParsingError(middleware.GetTrackingId(c)))
		return
	}

	var payload model.UpdateUserAllFieldRequest
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, base.ResponseFailure{
			Status:     http.StatusText(http.StatusBadRequest),
			Blame:      base.BlameUser,
			Message:    "failed to parse json",
			TrackingID: middleware.GetTrackingId(c),
		})
		return
	}

	if serviceErr := a.userService.Update(userID, payload, api.GetExtendedContext(c)); serviceErr != nil {
		c.JSON(serviceErr.Code, api.ResponseFromServiceError(*serviceErr, middleware.GetTrackingId(c)))
		return
	}

	c.JSON(http.StatusOK, base.ResponseOK{
		Status:     http.StatusText(http.StatusOK),
		TrackingID: middleware.GetTrackingId(c),
	})
}

// UpdateAuthorizationFields
// @Summary      Update User Authorization Fields
// @Description  Update User Authorization Fields
// @Tags         User
// @Accept       json
// @Produce      json
// @Security ApiKeyAuth
// @param Authorization header string true "Authorization"
// @Param        payload body   model.UpdateUserAuthorizationFieldsRequest true "User data"
// @Success      200  {object}  base.ResponseOK "OK"
// @Failure      400  {object}  base.ResponseFailure "Bad request"
// @Failure      500  {object}  base.ResponseFailure "Internal error (server fault)"
// @Router       /user/authorizationFields/update [post]
func (a *UserController) UpdateAuthorizationFields(c *gin.Context) {
	userID, _ := c.Get(middleware.UserIDKey)

	var payload model.UpdateUserAuthorizationFieldsRequest
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, base.ResponseFailure{
			Status:     http.StatusText(http.StatusBadRequest),
			Blame:      base.BlameUser,
			Message:    "failed to parse json",
			TrackingID: middleware.GetTrackingId(c),
		})
		return
	}

	if serviceErr := a.userService.UpdateAuthorizationFields(userID.(uuid.UUID), payload, api.GetExtendedContext(c)); serviceErr != nil {
		c.JSON(serviceErr.Code, api.ResponseFromServiceError(*serviceErr, middleware.GetTrackingId(c)))
		return
	}

	c.JSON(http.StatusOK, base.ResponseOK{
		Status:     http.StatusText(http.StatusOK),
		TrackingID: middleware.GetTrackingId(c),
	})
}

// RetrieveUser user-api
// @Summary     Retrieve data of an authorised user
// @Description  Retrieve data of an authorised user
// @Tags         User
// @Accept       json
// @Produce      json
// @Security ApiKeyAuth
// @param Authorization header string true "Authorization"
// @Success      200  {object}  model.GetUserResponse "OK"
// @Failure      400  {object}  base.ResponseFailure "Bad request"
// @Failure      500  {object}  base.ResponseFailure "Internal error (server fault)"
// @Router       /user/retrieve [get]
func (a *UserController) RetrieveUser(c *gin.Context) {
	userID, _ := c.Get(middleware.UserIDKey)
	userObjectModel, serviceErr := a.userService.RetrieveUser(userID.(uuid.UUID), api.GetExtendedContext(c))
	if serviceErr != nil {
		c.JSON(serviceErr.Code, api.ResponseFromServiceError(*serviceErr, middleware.GetTrackingId(c)))
		return
	}

	c.JSON(http.StatusOK, model.GetUserResponse{
		ResponseOK: base.ResponseOK{
			Status:     http.StatusText(http.StatusOK),
			TrackingID: middleware.GetTrackingId(c),
		},
		User: *userObjectModel,
	})
}

// GetUserByIdList private-user-api
// @Summary      Retrieve user information by id list
// @Description  Retrieve user information by id list
// @Tags         Users
// @Accept       json
// @Produce      json
// @Param        payload body   model.UsersByIdListRequest true "User data"
// @Success      200  {object}  model.GetUsersResponse "OK"
// @Failure      400  {object}  base.ResponseFailure "Bad request"
// @Failure      500  {object}  base.ResponseFailure "Internal error (server fault)"
// @Router       /usersByIdList [post]
func (a *UserController) GetUserByIdList(c *gin.Context) {
	var payload model.UsersByIdListRequest
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, base.ResponseFailure{
			Status:     http.StatusText(http.StatusBadRequest),
			Blame:      base.BlameUser,
			Message:    "failed to parse json",
			TrackingID: middleware.GetTrackingId(c),
		})
		return
	}
	if len(payload.Ids) == 0 {
		c.JSON(http.StatusOK, model.GetUsersResponse{
			ResponseOK: base.ResponseOK{
				Status:     http.StatusText(http.StatusOK),
				TrackingID: middleware.GetTrackingId(c),
			},
			Users: []model.UserObject{},
		})
	}

	res, serviceErr := a.userService.GetUsersById(payload.Ids, api.GetExtendedContext(c))
	if serviceErr != nil {
		c.JSON(serviceErr.Code, api.ResponseFromServiceError(*serviceErr, middleware.GetTrackingId(c)))
		return
	}
	c.JSON(http.StatusOK, model.GetUsersResponse{
		ResponseOK: base.ResponseOK{
			Status:     http.StatusText(http.StatusOK),
			TrackingID: middleware.GetTrackingId(c),
		},
		Users: res,
	})
}

func (a *UserController) GetUserById(c *gin.Context) {
	userId, err := uuid.Parse(c.Params.ByName("user-id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, api.GeneralParsingError(middleware.GetTrackingId(c)))
		return
	}
	userObjectModel, serviceError := a.userService.RetrieveUser(userId, api.GetExtendedContext(c))
	if serviceError != nil {
		if serviceError.Code != 404 {
			c.JSON(serviceError.Code, serviceError)
			return
		}
		c.JSON(serviceError.Code, base.ResponseOK{
			Status:     http.StatusText(http.StatusOK),
			TrackingID: middleware.GetTrackingId(c),
		})
		return
	}
	c.JSON(http.StatusOK, model.GetUserResponse{
		ResponseOK: base.ResponseOK{
			Status:     http.StatusText(http.StatusOK),
			TrackingID: middleware.GetTrackingId(c),
		},
		User: *userObjectModel,
	})
}
