package main

import (
	"log"
	"net/http"

	"github.com/houmingqi/lims-demo-server/internal/config"
	"github.com/houmingqi/lims-demo-server/internal/db"
	"github.com/houmingqi/lims-demo-server/internal/handler"
	"github.com/houmingqi/lims-demo-server/internal/repository"
	"github.com/houmingqi/lims-demo-server/internal/router"
	"github.com/houmingqi/lims-demo-server/internal/service"
)

func main() {
	cfg := config.Load()
	if err := cfg.Validate(); err != nil {
		log.Fatalf("config: %v", err)
	}

	gormDB, err := db.Open(cfg.DatabasePath)
	if err != nil {
		log.Fatalf("database: %v", err)
	}

	userRepo := repository.NewUserRepository(gormDB)
	noteRepo := repository.NewNoteRepository(gormDB)
	authSvc := service.NewAuthService(userRepo, cfg.JWTSecret)
	noteSvc := service.NewNoteService(noteRepo)

	r := router.New(cfg, router.Handlers{
		Auth: handler.NewAuthHandler(authSvc),
		Note: handler.NewNoteHandler(noteSvc),
	})

	log.Printf("demo-server listening on %s (APP_ENV=%s)", cfg.HTTPAddr, cfg.AppEnv)
	if err := http.ListenAndServe(cfg.HTTPAddr, r); err != nil {
		log.Fatal(err)
	}
}
