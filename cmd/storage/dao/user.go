package dao

import (
	"context"
	"errors"
	"github.com/google/uuid"
	"github.com/mag30/project-backend/domain/entity"
	"gorm.io/gorm"
)

type UserStorage struct {
	db *gorm.DB
}

func NewUserStorage(db *gorm.DB) *UserStorage {
	return &UserStorage{
		db: db,
	}
}

func (s UserStorage) Create(user *entity.User, ctx context.Context) error {
	return s.db.WithContext(ctx).Create(user).Error
}

func (s UserStorage) Retrieve(id uuid.UUID, ctx context.Context) (*entity.User, error) {
	var user entity.User
	err := s.db.WithContext(ctx).Preload("Sessions").First(&user, id).Error
	return &user, err
}

func (s UserStorage) Update(user *entity.User, ctx context.Context) error {
	return s.db.WithContext(ctx).Updates(user).Error
}

func (s UserStorage) Get(ctx context.Context) ([]entity.User, error) {
	var users []entity.User
	tx := s.db.WithContext(ctx).Model(&entity.User{}).Find(&users)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return users, nil
}

func (s UserStorage) GetUser(email string, ctx context.Context) (*entity.User, error) {
	var user entity.User
	tx := s.db.WithContext(ctx).Preload("Sessions").First(&user, "email = ?", email)
	if tx.Error != nil {
		return nil, tx.Error
	}

	if tx.RowsAffected == 0 {
		return nil, gorm.ErrRecordNotFound
	}
	return &user, nil
}

func (s UserStorage) GetByIDList(ids []uuid.UUID, ctx context.Context) ([]entity.User, error) {
	var users []entity.User
	tx := s.db.WithContext(ctx).Model(&entity.User{}).Where(ids)

	var err error

	if err := tx.Find(&users).Error; err != nil {
		if errors.Is(tx.Error, gorm.ErrRecordNotFound) {
			return []entity.User{}, nil
		}
		return []entity.User{}, err
	}

	return users, err
}
