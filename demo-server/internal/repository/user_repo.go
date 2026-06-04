package repository

import (
	"context"

	"gorm.io/gorm"

	"github.com/houmingqi/lims-demo-server/internal/model"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) FindByUsername(ctx context.Context, username string) (*model.User, error) {
	var row model.User
	err := r.db.WithContext(ctx).Where("username = ?", username).First(&row).Error
	if err != nil {
		return nil, err
	}
	return &row, nil
}
