package dao

import (
	"context"
	"github.com/google/uuid"
	"github.com/mag30/project-backend/domain/entity"
	"gorm.io/gorm"
)

type SessionStorage struct {
	db *gorm.DB
}

func NewSessionStorage(db *gorm.DB) *SessionStorage {
	return &SessionStorage{
		db: db,
	}
}

func (s SessionStorage) Create(session *entity.Session, ctx context.Context) error {
	return s.db.WithContext(ctx).Create(session).Error
}

func (s SessionStorage) Retrieve(id uuid.UUID, ctx context.Context) (*entity.Session, error) {
	var session *entity.Session
	tx := s.db.WithContext(ctx).First(&session, id)
	if tx.Error != nil {
		return nil, tx.Error
	}
	if tx.RowsAffected == 0 {
		return nil, gorm.ErrRecordNotFound
	}
	return session, nil
}
func (s SessionStorage) GetByUserIDAndUserIP(userID uuid.UUID, userIP string, ctx context.Context) (*entity.Session, error) {
	var session *entity.Session
	tx := s.db.WithContext(ctx).Where("user_id = ?", userID).Where("ip = ?", userIP).First(&session)
	if tx.Error != nil {
		return nil, tx.Error
	}
	if tx.RowsAffected == 0 {
		return nil, gorm.ErrRecordNotFound
	}
	return session, nil
}

func (s SessionStorage) GetByUserID(userID uuid.UUID, ctx context.Context) ([]entity.Session, error) {
	var sessions []entity.Session
	tx := s.db.WithContext(ctx).Where("user_id = ?", userID).Find(&sessions)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return sessions, nil
}

func (s SessionStorage) Delete(sessionID uuid.UUID, ctx context.Context) error {
	return s.db.WithContext(ctx).Unscoped().Delete(entity.Session{}, sessionID).Error
}
