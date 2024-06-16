package service

import (
	"context"
	"errors"
	"github.com/mag30/project-backend/cmd/storage/dao"
	"net/http"

	"github.com/google/uuid"
	authComponent "github.com/mag30/project-backend/auth"
	"github.com/mag30/project-backend/cmd/api/model"
	"github.com/mag30/project-backend/domain/base"
)

// UserService implements the UserService.IUserService interface for CRUD users.
type UserService struct {
	userStorage *dao.UserStorage
	authService *AuthService
	hasher      *authComponent.Hasher
}

// NewService creates a new instance of UserService.
func NewUserService(
	userStorage *dao.UserStorage,
	authService *AuthService,
	hasher *authComponent.Hasher) *UserService {
	return &UserService{
		userStorage: userStorage,
		authService: authService,
		hasher:      hasher,
	}
}

// Get retrieves a list of users.
func (s UserService) Get(ctx context.Context) ([]model.UserObject, *base.ServiceError) {
	users, err := s.userStorage.Get(ctx)
	if err != nil {
		return nil, base.NewPostgresReadError(err)
	}

	result := make([]model.UserObject, 0, len(users))
	for _, userEntity := range users {
		result = append(result, model.UserObject{
			ID:        userEntity.ID,
			CreatedAt: userEntity.CreatedAt,
			UpdatedAt: userEntity.UpdatedAt,
			Name:      userEntity.Name,
			Email:     userEntity.Email,
			Level:     userEntity.Level,
			Role:      userEntity.Role,
		})
	}

	return result, nil
}

// RetrieveUser retrieves a specific user by ID.
func (s UserService) RetrieveUser(id uuid.UUID, ctx context.Context) (*model.UserObject, *base.ServiceError) {
	userEntity, err := s.userStorage.Retrieve(id, ctx)
	if err != nil {
		return nil, base.NewPostgresReadError(err)
	}

	return &model.UserObject{
		ID:        userEntity.ID,
		CreatedAt: userEntity.CreatedAt,
		UpdatedAt: userEntity.UpdatedAt,
		Name:      userEntity.Name,
		Email:     userEntity.Email,
		Level:     userEntity.Level,
		Role:      userEntity.Role,
	}, nil
}

// Update updates user data.
func (s UserService) Update(id uuid.UUID, request model.UpdateUserAllFieldRequest, ctx context.Context) (mainErr *base.ServiceError) {
	userEntity, err := s.userStorage.Retrieve(id, ctx)
	if err != nil {
		return base.NewPostgresReadError(err)
	}

	userEntity.Email = request.Email
	if request.Password != nil {
		userEntity.Password, err = s.hasher.Hash(*request.Password)
	}

	if err != nil {
		return base.NewUnauthorizedError(err)
	}

	userEntity.Name = request.FullName
	userEntity.Level = request.Level

	if err := s.userStorage.Update(userEntity, ctx); err != nil {
		return base.NewPostgresWriteError(err)
	}

	return nil
}

// UpdateAuthorizationFields updates user authorization data.
func (s UserService) UpdateAuthorizationFields(id uuid.UUID, request model.UpdateUserAuthorizationFieldsRequest, ctx context.Context) (mainErr *base.ServiceError) {
	hashOldPassword, err := s.hasher.Hash(*request.OldPassword)
	if err != nil {
		return base.NewUnauthorizedError(err)
	}

	userEntity, err := s.userStorage.Retrieve(id, ctx)
	if err != nil {
		return base.NewPostgresReadError(err)
	}

	if userEntity.Password != hashOldPassword {
		return &base.ServiceError{
			Err:     errors.New("authentication failed. Please provide valid credentials"),
			Message: "authentication failed. Please provide valid credentials",
			Blame:   base.BlameUser,
			Code:    http.StatusUnauthorized,
		}
	}

	hashNewPassword, err := s.hasher.Hash(request.NewPassword)
	if err != nil {
		return base.NewUnauthorizedError(err)
	}

	userEntity.Email = request.Email
	userEntity.Password = hashNewPassword

	if err := s.userStorage.Update(userEntity, ctx); err != nil {
		return base.NewPostgresWriteError(err)
	}

	if serviceErr := s.authService.SignOutAllSession(userEntity.ID, ctx); serviceErr != nil {
		return serviceErr
	}

	return nil
}

// GetUsersById retrieves a list of users by their IDs.
func (s UserService) GetUsersById(ids []uuid.UUID, ctx context.Context) ([]model.UserObject, *base.ServiceError) {
	users, err := s.userStorage.GetByIDList(ids, ctx)
	if err != nil {
		return nil, base.NewPostgresReadError(err)
	}

	var result []model.UserObject
	for _, u := range users {
		result = append(result, model.UserObject{
			ID:        u.ID,
			CreatedAt: u.CreatedAt,
			UpdatedAt: u.UpdatedAt,
			Name:      u.Name,
			Email:     u.Email,
			Level:     u.Level,
			Role:      u.Role,
		})
	}

	return result, nil
}
