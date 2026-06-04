package service

import (
	"context"
	"errors"
	"strings"
	"time"

	"gorm.io/gorm"

	"github.com/houmingqi/lims-demo-server/internal/dto"
	"github.com/houmingqi/lims-demo-server/internal/model"
	"github.com/houmingqi/lims-demo-server/internal/pkg/apperror"
	"github.com/houmingqi/lims-demo-server/internal/repository"
)

type NoteService struct {
	notes *repository.NoteRepository
}

func NewNoteService(notes *repository.NoteRepository) *NoteService {
	return &NoteService{notes: notes}
}

func (s *NoteService) Create(ctx context.Context, userID uint64, req dto.CreateNoteReq) (*model.Note, error) {
	row := &model.Note{
		UserID:  userID,
		Title:   strings.TrimSpace(req.Title),
		Content: req.Content,
	}
	if row.Title == "" {
		return nil, apperror.ErrBadRequest
	}
	if err := s.notes.Create(ctx, row); err != nil {
		return nil, err
	}
	return row, nil
}

func (s *NoteService) List(ctx context.Context, userID uint64, req dto.NoteListReq) ([]model.Note, int64, error) {
	return s.notes.List(ctx, userID, req.Offset(), req.PageSize, strings.TrimSpace(req.Keyword))
}

func (s *NoteService) GetByID(ctx context.Context, userID, id uint64) (*model.Note, error) {
	row, err := s.notes.GetByID(ctx, id, userID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, apperror.ErrNotFound
		}
		return nil, err
	}
	return row, nil
}

func (s *NoteService) Update(ctx context.Context, userID, id uint64, req dto.UpdateNoteReq) error {
	title := strings.TrimSpace(req.Title)
	if title == "" {
		return apperror.ErrBadRequest
	}
	ok, err := s.notes.UpdateCAS(ctx, id, userID, title, req.Content, req.UpdatedAt)
	if err != nil {
		return err
	}
	if !ok {
		return apperror.ErrConflict
	}
	return nil
}

func (s *NoteService) Delete(ctx context.Context, userID, id uint64, updatedAt time.Time) error {
	ok, err := s.notes.DeleteCAS(ctx, id, userID, updatedAt)
	if err != nil {
		return err
	}
	if !ok {
		return apperror.ErrConflict
	}
	return nil
}
