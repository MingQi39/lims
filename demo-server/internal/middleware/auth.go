package middleware

import (
	"strings"

	"github.com/gin-gonic/gin"

	"github.com/houmingqi/lims-demo-server/internal/pkg/apperror"
	"github.com/houmingqi/lims-demo-server/internal/pkg/auth"
	"github.com/houmingqi/lims-demo-server/internal/pkg/response"
)

func Auth(jwtSecret string) gin.HandlerFunc {
	return func(c *gin.Context) {
		header := c.GetHeader("Authorization")
		if header == "" {
			response.Fail(c, apperror.ErrUnauthorized)
			c.Abort()
			return
		}
		parts := strings.SplitN(header, " ", 2)
		if len(parts) != 2 || parts[0] != "Bearer" {
			response.Fail(c, apperror.ErrUnauthorized)
			c.Abort()
			return
		}
		claims, err := auth.ParseAccessToken(parts[1], jwtSecret)
		if err != nil {
			response.Fail(c, apperror.ErrUnauthorized)
			c.Abort()
			return
		}
		c.Set("user_id", claims.UserID)
		c.Set("user_name", claims.Username)
		c.Next()
	}
}
