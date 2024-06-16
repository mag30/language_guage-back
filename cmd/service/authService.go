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
	"github.com/mag30/project-backend/domain/entity"
	"gorm.io/gorm"
)

// AuthService implements the AuthService.IAuthService interface for user authentication and authorization operations.
type AuthService struct {
	userStorage    *dao.UserStorage
	storageSession *dao.SessionStorage
	hasher         *authComponent.Hasher
	jwtManager     *authComponent.JWTManager
}

// NewService creates a new instance of AuthService.
func NewAuthService(
	userStorage *dao.UserStorage,
	storageSession *dao.SessionStorage,
	hasher *authComponent.Hasher,
	jwtManager *authComponent.JWTManager) *AuthService {
	return &AuthService{
		userStorage:    userStorage,
		storageSession: storageSession,
		hasher:         hasher,
		jwtManager:     jwtManager,
	}
}

// Register handles the user registration process.
func (s AuthService) Register(request *model.RegisterRequest, ctx context.Context) (_uuid *uuid.UUID, mainErr *base.ServiceError) {
	hashPassword, err := s.hasher.Hash(request.Password)
	if err != nil {
		return nil, base.NewReadByteError(err)
	}

	user := &entity.User{
		Name:     request.Name,
		Email:    request.Email,
		Password: hashPassword,
	}

	if err := s.userStorage.Create(user, ctx); err != nil {
		return nil, base.NewPostgresWriteError(err)
	}

	return &user.ID, nil
}

// Login handles the user login process.
func (s AuthService) Login(request *model.LoginRequest, ctx context.Context) (jwt *string, refToken *uuid.UUID, mainErr *base.ServiceError) {
	user, err := s.userStorage.GetUser(request.Email, ctx)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil, base.NewLoginError(err)
		}
		return nil, nil, base.NewPostgresReadError(err)
	}

	hashPassword, err := s.hasher.Hash(request.Password)
	if err != nil {
		return nil, nil, base.NewReadByteError(err)
	}

	if hashPassword != user.Password {
		return nil, nil, base.NewLoginError(errors.New("incorrect password"))
	}

	session := &entity.Session{
		UserID: user.ID,
	}

	if err := s.storageSession.Create(session, ctx); err != nil {
		return nil, nil, base.NewPostgresWriteError(err)
	}

	token, err := s.jwtManager.NewJWT(user.ID)
	if err != nil {
		return nil, nil, base.NewCreateJWTError(err)
	}

	return &token, &session.ID, nil
}

// SignOutAllSession signs out all active sessions for a user.
func (s AuthService) SignOutAllSession(id uuid.UUID, ctx context.Context) (mainErr *base.ServiceError) {
	sessions, err := s.storageSession.GetByUserID(id, ctx)
	if err != nil {
		return base.NewPostgresReadError(err)
	}

	for _, session := range sessions {
		if serviceErr := s.storageSession.Delete(session.ID, ctx); serviceErr != nil {
			return base.NewPostgresReadError(serviceErr)
		}
	}

	return nil
}

// Logout logs out a user session based on the refresh token.
func (s AuthService) Logout(refreshToken uuid.UUID, ctx context.Context) (mainErr *base.ServiceError) {
	session, err := s.storageSession.Retrieve(refreshToken, ctx)
	if err != nil {
		return &base.ServiceError{
			Err:     err,
			Blame:   base.BlameUser,
			Code:    http.StatusUnauthorized,
			Message: "failed get session",
		}
	}

	if err := s.storageSession.Delete(session.ID, ctx); err != nil {
		return base.NewPostgresReadError(err)
	}

	return nil
}

// RefreshJWT refreshes the JWT token for a user session.
func (s AuthService) RefreshJWT(id uuid.UUID, ctx context.Context) (*string, *uuid.UUID, *uuid.UUID, *base.ServiceError) {
	session, err := s.storageSession.Retrieve(id, ctx)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil, nil, base.NewNotSessionError(err)
		}
		return nil, nil, nil, base.NewPostgresReadError(err)
	}

	newSession := &entity.Session{
		UserID: session.UserID,
	}

	if err := s.storageSession.Delete(session.ID, ctx); err != nil {
		return nil, nil, nil, base.NewPostgresReadError(err)
	}

	if err := s.storageSession.Create(newSession, ctx); err != nil {
		return nil, nil, nil, base.NewPostgresWriteError(err)
	}

	token, err := s.jwtManager.NewJWT(session.UserID)
	if err != nil {
		return nil, nil, nil, base.NewCreateJWTError(err)
	}

	return &token, &newSession.ID, &newSession.UserID, nil
}
