package handler

import (
	"github.com/gin-gonic/gin"

	"github.com/houmingqi/lims-demo-server/internal/dto"
	"github.com/houmingqi/lims-demo-server/internal/pkg/ginx"
	"github.com/houmingqi/lims-demo-server/internal/pkg/response"
	"github.com/houmingqi/lims-demo-server/internal/service"
)

type AuthHandler struct {
	svc *service.AuthService
}

func NewAuthHandler(svc *service.AuthService) *AuthHandler {
	return &AuthHandler{svc: svc}
}

func (h *AuthHandler) Login(c *gin.Context) {
	req, ok := ginx.BindJSON[dto.LoginReq](c)
	if !ok {
		return
	}
	resp, err := h.svc.Login(c.Request.Context(), req)
	if err != nil {
		response.Fail(c, err)
		return
	}
	response.Success(c, resp)
}
