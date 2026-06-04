package response

import (
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"

	"github.com/houmingqi/lims-demo-server/internal/pkg/apperror"
)

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Detail  string      `json:"detail,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

func Success(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, Response{Code: 0, Message: "ok", Data: data})
}

func SuccessWithPage(c *gin.Context, list interface{}, total int64, page, pageSize int) {
	c.JSON(http.StatusOK, Response{
		Code:    0,
		Message: "ok",
		Data: gin.H{
			"list":      list,
			"total":     total,
			"page":      page,
			"page_size": pageSize,
		},
	})
}

func Fail(c *gin.Context, err error) {
	if appErr, ok := err.(*apperror.AppError); ok {
		c.JSON(appErr.HTTPStatus, Response{
			Code:    appErr.Code,
			Message: appErr.Message,
			Detail:  appErr.Detail,
		})
		return
	}
	resp := Response{Code: 50000, Message: "服务器内部错误"}
	if strings.ToLower(os.Getenv("APP_ENV")) == "development" {
		resp.Detail = err.Error()
	}
	c.JSON(http.StatusInternalServerError, resp)
}
