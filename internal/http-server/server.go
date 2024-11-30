package server

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"

	"Template/internal/config"
	"Template/internal/logger"
)

const (
	DefaultReadTimeout       = 10 * time.Second
	DefaultWriteTimeout      = 0 * time.Second
	DefaultIdleTimeout       = 60 * time.Second
	DefaultMaxHeaderBytes    = 0
	DefaultReadHeaderTimeout = 2 * time.Second
)

type App struct {
	server *http.Server
	config *config.Config
	logger *logrus.Logger
}

func New(cfg config.Config) *App {

	return &App{
		server: nil,
		config: &cfg,
		logger: logger.GetLogger(),
	}
}

func (a *App) Start(ctx context.Context) error {
	ginEngine := gin.New()
	ginLogger := a.logger.WithFields(logrus.Fields{"component": "gin"})
	gin.DefaultWriter = ginLogger.Writer()
	gin.DefaultErrorWriter = ginLogger.Writer()

	// TODO: Middlewares
	ginEngine.Use(gin.Logger())
	ginEngine.Use(gin.Recovery())

	a.routerRegister(&ginEngine.RouterGroup)

	a.server = &http.Server{
		Addr:              a.config.HTTPServer.Address,
		Handler:           ginEngine,
		ReadTimeout:       DefaultReadTimeout,
		WriteTimeout:      DefaultWriteTimeout,
		IdleTimeout:       DefaultIdleTimeout,
		MaxHeaderBytes:    DefaultMaxHeaderBytes,
		ReadHeaderTimeout: DefaultReadHeaderTimeout,
		ConnState:         nil,
		BaseContext:       nil,
		ConnContext:       nil,
		TLSNextProto:      nil,
		TLSConfig:         nil,
		ErrorLog:          nil,
	}

	err := a.server.ListenAndServe()
	if err != nil && !errors.Is(err, http.ErrServerClosed) {
		return fmt.Errorf("%w", err)
	}

	return nil
}

func (a *App) Stop(ctx context.Context) error {
	if err := a.server.Shutdown(ctx); err != nil {
		return fmt.Errorf("%w", err)
	}

	return nil
}
