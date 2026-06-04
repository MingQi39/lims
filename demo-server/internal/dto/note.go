package dto

import "time"

type NoteListReq struct {
	Page     int    `form:"page"`
	PageSize int    `form:"page_size"`
	Keyword  string `form:"keyword"`
}

func (r *NoteListReq) Normalize() {
	if r.Page < 1 {
		r.Page = 1
	}
	if r.PageSize < 1 {
		r.PageSize = 20
	}
	if r.PageSize > 100 {
		r.PageSize = 100
	}
}

func (r NoteListReq) Offset() int {
	return (r.Page - 1) * r.PageSize
}

type CreateNoteReq struct {
	Title   string `json:"title" binding:"required,max=200"`
	Content string `json:"content"`
}

type UpdateNoteReq struct {
	Title     string    `json:"title" binding:"required,max=200"`
	Content   string    `json:"content"`
	UpdatedAt time.Time `json:"updated_at" binding:"required"`
}
