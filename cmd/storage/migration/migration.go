package migration

import (
	"errors"
	"github.com/google/uuid"
	"github.com/mag30/project-backend/domain/entity"
	"github.com/mag30/project-backend/domain/enum"
	"gorm.io/gorm"
	"strings"
)

func Migration(
	db *gorm.DB,
	adminID uuid.UUID,
	adminUserName string,
	adminEmail string,
	adminPassword string) error {

	if err := db.AutoMigrate(
		&entity.User{},
		&entity.Session{},
		&entity.Quiz{},
		&entity.Task{},
		&entity.Result{},
	); err != nil {
		//relationship doesn't exist
		if !strings.Contains(err.Error(), "42P07") {
			return err
		}
	}

	if err := adminMigration(db, adminID, adminUserName, adminEmail, adminPassword); err != nil {
		return err
	}

	return nil
}

func adminMigration(db *gorm.DB,
	adminID uuid.UUID,
	adminUserName string,
	adminEmail string,
	adminPassword string) error {

	tx := db.First(&entity.User{}, adminID)
	if tx.Error != nil {
		if !errors.Is(tx.Error, gorm.ErrRecordNotFound) {
			return tx.Error
		}
	}

	if tx.RowsAffected == 0 {
		user := &entity.User{
			Email:    adminEmail,
			Name:     adminUserName,
			Password: adminPassword,
			Role:     enum.Admin,
		}
		user.ID = adminID

		if err := db.Create(user).Error; err != nil {
			return err
		}
	}

	return nil
}
