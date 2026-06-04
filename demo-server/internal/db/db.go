package db

import (
	"fmt"
	"os"
	"path/filepath"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"github.com/houmingqi/lims-demo-server/internal/model"
)

func Open(databasePath string) (*gorm.DB, error) {
	if err := os.MkdirAll(filepath.Dir(databasePath), 0o755); err != nil {
		return nil, fmt.Errorf("create db dir: %w", err)
	}
	db, err := gorm.Open(sqlite.Open(databasePath), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	if err := db.AutoMigrate(&model.User{}, &model.Note{}); err != nil {
		return nil, err
	}
	if err := seedDemoUser(db); err != nil {
		return nil, err
	}
	return db, nil
}

func seedDemoUser(db *gorm.DB) error {
	var count int64
	if err := db.Model(&model.User{}).Count(&count).Error; err != nil {
		return err
	}
	if count > 0 {
		return nil
	}
	hash, err := bcrypt.GenerateFromPassword([]byte("demo123"), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	return db.Create(&model.User{
		Username:     "demo",
		PasswordHash: string(hash),
		DisplayName:  "Demo User",
	}).Error
}
