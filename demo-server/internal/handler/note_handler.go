package handler

import (
	"github.com/gin-gonic/gin"

	"github.com/houmingqi/lims-demo-server/internal/dto"
	"github.com/houmingqi/lims-demo-server/internal/pkg/ginx"
	"github.com/houmingqi/lims-demo-server/internal/pkg/response"
	"github.com/houmingqi/lims-demo-server/internal/service"
)

type NoteHandler struct {
	svc *service.NoteService
}

func NewNoteHandler(svc *service.NoteService) *NoteHandler {
	return &NoteHandler{svc: svc}
}

func userIDFromCtx(c *gin.Context) uint64 {
	v, _ := c.Get("user_id")
	id, _ := v.(uint64)
	return id
}

func (h *NoteHandler) Create(c *gin.Context) {
	req, ok := ginx.BindJSON[dto.CreateNoteReq](c)
	if !ok {
		return
	}
	row, err := h.svc.Create(c.Request.Context(), userIDFromCtx(c), req)
	if err != nil {
		response.Fail(c, err)
		return
	}
	response.Success(c, row)
}

func (h *NoteHandler) List(c *gin.Context) {
	req, ok := ginx.BindQuery[dto.NoteListReq](c)
	if !ok {
		return
	}
	req.Normalize()
	list, total, err := h.svc.List(c.Request.Context(), userIDFromCtx(c), req)
	if err != nil {
		response.Fail(c, err)
		return
	}
	response.SuccessWithPage(c, list, total, req.Page, req.PageSize)
}

func (h *NoteHandler) Get(c *gin.Context) {
	id, ok := ginx.ParamID(c)
	if !ok {
		return
	}
	row, err := h.svc.GetByID(c.Request.Context(), userIDFromCtx(c), id)
	if err != nil {
		response.Fail(c, err)
		return
	}
	response.Success(c, row)
}

func (h *NoteHandler) Update(c *gin.Context) {
	id, ok := ginx.ParamID(c)
	if !ok {
		return
	}
	req, ok := ginx.BindJSON[dto.UpdateNoteReq](c)
	if !ok {
		return
	}
	if err := h.svc.Update(c.Request.Context(), userIDFromCtx(c), id, req); err != nil {
		response.Fail(c, err)
		return
	}
	response.Success(c, nil)
}

func (h *NoteHandler) Delete(c *gin.Context) {
	id, ok := ginx.ParamID(c)
	if !ok {
		return
	}
	updatedAt, ok := ginx.UpdatedAtQuery(c)
	if !ok {
		return
	}
	if err := h.svc.Delete(c.Request.Context(), userIDFromCtx(c), id, *updatedAt); err != nil {
		response.Fail(c, err)
		return
	}
	response.Success(c, nil)
}
