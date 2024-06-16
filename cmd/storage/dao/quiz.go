package dao

import (
	"context"
	"github.com/google/uuid"
	"github.com/mag30/project-backend/domain/entity"
	"gorm.io/gorm"
)

type QuizStorage struct {
	db *gorm.DB
}

func NewQuizStorage(db *gorm.DB) *QuizStorage {
	return &QuizStorage{
		db: db,
	}
}

func (s QuizStorage) Create(quiz *entity.Quiz, ctx context.Context) error {
	return s.db.WithContext(ctx).Create(quiz).Error
}

func (s QuizStorage) Retrieve(id uuid.UUID, ctx context.Context) (*entity.Quiz, error) {
	var quiz entity.Quiz
	err := s.db.WithContext(ctx).Preload("Tasks").First(&quiz, id).Error
	return &quiz, err
}

func (s QuizStorage) GetByName(name string, ctx context.Context) (*entity.Quiz, error) {
	var quiz entity.Quiz
	tx := s.db.WithContext(ctx).Preload("Tasks").First(&quiz, "name = ?", name)
	if tx.Error != nil {
		return nil, tx.Error
	}

	if tx.RowsAffected == 0 {
		return nil, gorm.ErrRecordNotFound
	}
	return &quiz, nil
}
