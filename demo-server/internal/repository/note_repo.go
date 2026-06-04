package repository

import (
	"context"
	"time"

	"gorm.io/gorm"

	"github.com/houmingqi/lims-demo-server/internal/model"
)

type NoteRepository struct {
	db *gorm.DB
}

func NewNoteRepository(db *gorm.DB) *NoteRepository {
	return &NoteRepository{db: db}
}

func (r *NoteRepository) Create(ctx context.Context, row *model.Note) error {
	return r.db.WithContext(ctx).Create(row).Error
}

func (r *NoteRepository) GetByID(ctx context.Context, id, userID uint64) (*model.Note, error) {
	var row model.Note
	err := r.db.WithContext(ctx).Where("id = ? AND user_id = ?", id, userID).First(&row).Error
	if err != nil {
		return nil, err
	}
	return &row, nil
}

func (r *NoteRepository) List(ctx context.Context, userID uint64, offset, limit int, keyword string) ([]model.Note, int64, error) {
	q := r.db.WithContext(ctx).Model(&model.Note{}).Where("user_id = ?", userID)
	if keyword != "" {
		like := "%" + keyword + "%"
		q = q.Where("title LIKE ? OR content LIKE ?", like, like)
	}
	var total int64
	if err := q.Count(&total).Error; err != nil {
		return nil, 0, err
	}
	var rows []model.Note
	err := q.Order("updated_at DESC").Offset(offset).Limit(limit).Find(&rows).Error
	return rows, total, err
}

func (r *NoteRepository) UpdateCAS(ctx context.Context, id, userID uint64, title, content string, expectedUpdatedAt time.Time) (bool, error) {
	now := time.Now()
	res := r.db.WithContext(ctx).Model(&model.Note{}).
		Where("id = ? AND user_id = ? AND updated_at = ?", id, userID, expectedUpdatedAt).
		Updates(map[string]interface{}{
			"title":      title,
			"content":    content,
			"updated_at": now,
		})
	if res.Error != nil {
		return false, res.Error
	}
	return res.RowsAffected > 0, nil
}

func (r *NoteRepository) DeleteCAS(ctx context.Context, id, userID uint64, expectedUpdatedAt time.Time) (bool, error) {
	res := r.db.WithContext(ctx).
		Where("id = ? AND user_id = ? AND updated_at = ?", id, userID, expectedUpdatedAt).
		Delete(&model.Note{})
	if res.Error != nil {
		return false, res.Error
	}
	return res.RowsAffected > 0, nil
}
