package main

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"go.uber.org/zap"

	"github.com/blattaria7/go-template/config"
	"github.com/blattaria7/go-template/internal/handlers"
	"github.com/blattaria7/go-template/internal/repository"
	"github.com/blattaria7/go-template/internal/service"
	"github.com/blattaria7/go-template/pkg/logger"
)

func main() {
	var cfg config.Config

	if err := cfg.Parse(); err != nil {
		fmt.Errorf("error parsing config: %w", err)
	}

	log, err := logger.InitLogger(&cfg.Logger)
	if err != nil {
		fmt.Errorf("error initializing logger: %w", err)
	}

	// initialize your storage and storage's variables
	items := make(map[string]string, 1)
	items["a"] = "apple"
	repo := repository.NewRepository(items, log)

	// initialize your service
	svc := service.NewService(repo, log)

	// initialize your handlers
	handler := handlers.NewHandler(svc, log)

	// initialize ad describe your routers
	r := mux.NewRouter()
	r.HandleFunc("/healthcheck", handler.Healthcheck).Methods(http.MethodGet)
	r.HandleFunc("/items/{id}", handler.Get).Methods(http.MethodGet)

	server := &http.Server{
		Handler:      r,
		Addr:         fmt.Sprintf(":%d", cfg.HTTPConfig.Port),
		WriteTimeout: cfg.HTTPConfig.WriteTimeout,
		ReadTimeout:  cfg.HTTPConfig.ReadTimeout,
	}

	log.Debug("service started on:", zap.Uint("port", cfg.HTTPConfig.Port))

	// TODO: graceful shutdown
	if err = server.ListenAndServe(); err != nil {
		if err := server.Shutdown(context.Background()); err != nil {
			log.Info("service shutting down at", zap.Time("time", time.Now()))
			log.Error("server shutdown error", zap.Error(err))
		}
	}
}
