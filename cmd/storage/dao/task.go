package dao

import (
	"context"
	"github.com/google/uuid"
	"github.com/mag30/project-backend/domain/entity"
	"gorm.io/gorm"
)

type TaskStorage struct {
	db *gorm.DB
}

func NewTaskStorage(db *gorm.DB) *TaskStorage {
	return &TaskStorage{
		db: db,
	}
}

func (s TaskStorage) Create(task *entity.Task, ctx context.Context) error {
	return s.db.WithContext(ctx).Create(task).Error
}

func (s TaskStorage) Retrieve(id uuid.UUID, ctx context.Context) (*entity.Task, error) {
	var task entity.Task
	err := s.db.WithContext(ctx).Preload("Quiz").First(&task, id).Error
	return &task, err
}

func (s TaskStorage) Update(task *entity.Task, ctx context.Context) error {
	return s.db.WithContext(ctx).Updates(task).Error
}

func (s TaskStorage) GetByQuizID(quizID uuid.UUID, ctx context.Context) ([]entity.Task, error) {
	var tasks []entity.Task
	tx := s.db.WithContext(ctx).Model(&entity.Task{}).Where("quiz_id = ?", quizID).Find(&tasks)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return tasks, nil
}

func (s TaskStorage) GetByNameAndQuizID(quizID uuid.UUID, name string, ctx context.Context) (*entity.Task, error) {
	var task entity.Task
	tx := s.db.WithContext(ctx).First(&task, "name = ? AND quiz_id = ?", name, quizID)
	if tx.Error != nil {
		return nil, tx.Error
	}

	if tx.RowsAffected == 0 {
		return nil, gorm.ErrRecordNotFound
	}
	return &task, nil
}
