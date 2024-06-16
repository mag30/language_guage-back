package dao

import (
	"context"
	"github.com/google/uuid"
	"github.com/mag30/project-backend/domain/entity"
	"gorm.io/gorm"
)

type ResultStorage struct {
	db *gorm.DB
}

func NewResultStorage(db *gorm.DB) *ResultStorage {
	return &ResultStorage{
		db: db,
	}
}

func (s ResultStorage) Create(result *entity.Result, ctx context.Context) error {
	return s.db.WithContext(ctx).Create(result).Error
}

func (s ResultStorage) Retrieve(id uuid.UUID, ctx context.Context) (*entity.Result, error) {
	var result entity.Result
	err := s.db.WithContext(ctx).Preload("User").Preload("Quiz").First(&result, id).Error
	return &result, err
}

func (s ResultStorage) Update(result *entity.Result, ctx context.Context) error {
	return s.db.WithContext(ctx).Updates(result).Error
}

func (s ResultStorage) GetByUserIDAndQuizID(userID uuid.UUID, quizID uuid.UUID, ctx context.Context) (*entity.Result, error) {
	var results *entity.Result
	tx := s.db.WithContext(ctx).Model(&entity.Result{}).Where("user_id = ?", userID).Where("quiz_id = ?", quizID).First(&results)
	if tx.Error != nil {
		return nil, tx.Error
	}

	if tx.RowsAffected == 0 {
		return nil, gorm.ErrRecordNotFound
	}

	return results, nil
}

func (s ResultStorage) Delete(id uuid.UUID, ctx context.Context) error {
	return s.db.WithContext(ctx).Unscoped().Delete(&entity.Result{}, id).Error
}
