package base

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

// EntityWithIdKey is a base db entity with uuid.UUID as a primary key.
type EntityWithIdKey struct {
	ID        uuid.UUID      `json:"id" gorm:"type:uuid;default:uuid_generate_v1();primaryKey"`
	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `json:"-"`
}

// EntityWithIdKeyUniqueIndex is a base db entity with uuid.UUID as a unique index.
type EntityWithIdKeyUniqueIndex struct {
	ID        uuid.UUID      `json:"id" gorm:"type:uuid;default:uuid_generate_v1();uniqueIndex"`
	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `json:"-"`
}
