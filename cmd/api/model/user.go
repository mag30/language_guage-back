package model

import (
	"github.com/google/uuid"
	"github.com/mag30/project-backend/domain/base"
	"github.com/mag30/project-backend/domain/enum"
	"time"
)

type RegisterRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UpdateUserAllFieldRequest struct {
	Email    string     `json:"email"`
	Password *string    `json:"password"`
	FullName string     `json:"fullName"`
	Level    enum.Level `json:"level"`
}

type UpdateUserAuthorizationFieldsRequest struct {
	Email       string  `json:"email"`
	OldPassword *string `json:"old_password"`
	NewPassword string  `json:"new_password"`
}

type RecreateJWTRequest struct {
	RefreshToken uuid.UUID `json:"refreshToken"`
}

type LoginResponse struct {
	base.ResponseOK
	JWT          string    `json:"token"`
	RefreshToken uuid.UUID `json:"refreshToken"`
}

type UsersByIdListRequest struct {
	Ids []uuid.UUID `json:"ids"`
}

type UserObject struct {
	ID        uuid.UUID  `json:"id"`
	CreatedAt time.Time  `json:"createdAt"`
	UpdatedAt time.Time  `json:"updatedAt"`
	Name      string     `json:"name"`
	Email     string     `json:"email"`
	Level     enum.Level `json:"level"`
	Role      enum.Role  `json:"role"`
}

type GetUserResponse struct {
	base.ResponseOK
	User UserObject `json:"user"`
}

type GetUsersResponse struct {
	base.ResponseOK
	Users []UserObject `json:"users"`
}
