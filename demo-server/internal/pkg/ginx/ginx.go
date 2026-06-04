package ginx

import (
	"strconv"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/houmingqi/lims-demo-server/internal/pkg/apperror"
	"github.com/houmingqi/lims-demo-server/internal/pkg/response"
)

func BindJSON[T any](c *gin.Context) (T, bool) {
	var req T
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Fail(c, apperror.New(apperror.ErrBadRequest.Code, apperror.ErrBadRequest.HTTPStatus, "请求体无效", err.Error()))
		return req, false
	}
	return req, true
}

func BindQuery[T any](c *gin.Context) (T, bool) {
	var req T
	if err := c.ShouldBindQuery(&req); err != nil {
		response.Fail(c, apperror.New(apperror.ErrBadRequest.Code, apperror.ErrBadRequest.HTTPStatus, "查询参数无效", err.Error()))
		return req, false
	}
	return req, true
}

func ParamID(c *gin.Context) (uint64, bool) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil || id == 0 {
		response.Fail(c, apperror.ErrBadRequest)
		return 0, false
	}
	return id, true
}

func UpdatedAtQuery(c *gin.Context) (*time.Time, bool) {
	raw := c.Query("updated_at")
	if raw == "" {
		response.Fail(c, apperror.New(apperror.ErrBadRequest.Code, apperror.ErrBadRequest.HTTPStatus, "缺少 updated_at", ""))
		return nil, false
	}
	t, err := time.Parse(time.RFC3339Nano, raw)
	if err != nil {
		t, err = time.Parse(time.RFC3339, raw)
	}
	if err != nil {
		response.Fail(c, apperror.New(apperror.ErrBadRequest.Code, apperror.ErrBadRequest.HTTPStatus, "updated_at 格式无效", ""))
		return nil, false
	}
	return &t, true
}
