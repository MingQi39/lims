package router

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"

	"github.com/houmingqi/lims-demo-server/internal/config"
	"github.com/houmingqi/lims-demo-server/internal/handler"
	"github.com/houmingqi/lims-demo-server/internal/middleware"
)

type Handlers struct {
	Auth *handler.AuthHandler
	Note *handler.NoteHandler
}

func New(cfg config.Config, h Handlers) *gin.Engine {
	if cfg.AppEnv != "development" {
		gin.SetMode(gin.ReleaseMode)
	}
	r := gin.New()
	r.Use(gin.Recovery(), cors(cfg.CORSOrigins))

	r.GET("/healthz", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	})

	v1 := r.Group("/api/v1")
	v1.POST("/auth/login", h.Auth.Login)

	protected := v1.Group("")
	protected.Use(middleware.Auth(cfg.JWTSecret))
	protected.GET("/notes", h.Note.List)
	protected.POST("/notes", h.Note.Create)
	protected.GET("/notes/:id", h.Note.Get)
	protected.PUT("/notes/:id", h.Note.Update)
	protected.DELETE("/notes/:id", h.Note.Delete)

	return r
}

func cors(origins []string) gin.HandlerFunc {
	allowAll := len(origins) == 0
	originSet := make(map[string]struct{}, len(origins))
	for _, o := range origins {
		originSet[o] = struct{}{}
	}
	return func(c *gin.Context) {
		origin := c.GetHeader("Origin")
		if allowAll || (origin != "" && containsOrigin(originSet, origin)) {
			if origin != "" {
				c.Header("Access-Control-Allow-Origin", origin)
			}
			c.Header("Access-Control-Allow-Credentials", "true")
			c.Header("Access-Control-Allow-Headers", "Authorization, Content-Type")
			c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		}
		if c.Request.Method == http.MethodOptions {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}
		c.Next()
	}
}

func containsOrigin(set map[string]struct{}, origin string) bool {
	_, ok := set[origin]
	return ok || strings.Contains(origin, "localhost")
}
