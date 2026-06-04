package apperror

import "net/http"

type AppError struct {
	Code       int
	HTTPStatus int
	Message    string
	Detail     string
}

func (e *AppError) Error() string {
	if e.Detail != "" {
		return e.Message + ": " + e.Detail
	}
	return e.Message
}

func New(code, httpStatus int, message, detail string) *AppError {
	return &AppError{Code: code, HTTPStatus: httpStatus, Message: message, Detail: detail}
}

var (
	ErrUnauthorized = New(40100, http.StatusUnauthorized, "未登录或登录已失效", "")
	ErrForbidden    = New(40300, http.StatusForbidden, "没有操作权限", "")
	ErrNotFound     = New(40400, http.StatusNotFound, "资源不存在", "")
	ErrConflict     = New(40900, http.StatusConflict, "数据已被他人修改，请刷新后重试", "")
	ErrBadRequest   = New(40000, http.StatusBadRequest, "请求参数无效", "")
	ErrInternal     = New(50000, http.StatusInternalServerError, "服务器内部错误", "")
)
