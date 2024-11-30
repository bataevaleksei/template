package app

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	_ "github.com/mattn/go-sqlite3"
	"github.com/sirupsen/logrus"

	"Template/internal/config"
	server "Template/internal/http-server"
	"Template/internal/logger"
	"Template/internal/storage/sqlite"
)

func Run(ctx context.Context) error {

	cfg := *config.LoadConfig(ctx)

	if err := logger.InitLogger(cfg.Env); err != nil {
		return fmt.Errorf("failed to init logger: %w", err)
	}

	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	log := logger.GetLogger()
	log.WithFields(logrus.Fields{
		"env":     cfg.Env,
		"address": cfg.HTTPServer.Address,
	}).Info("starting server")

	if _, err := sqlite.New(); err != nil {
		return fmt.Errorf("failed to load storage: %w", err)
	}

	srv := server.New(cfg)
	go func() {
		if err := srv.Start(ctx); err != nil {
			log.Errorf("server has been shut down: %v", err)
			done <- syscall.SIGTERM
		}
	}()

	<-done

	if err := srv.Stop(ctx); err != nil {
		return fmt.Errorf("failed to stop server: %w", err)
	}

	log.Info("server closed")

	return nil
}
