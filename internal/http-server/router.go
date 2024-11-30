package server

import (
	"Template/internal/transport/handlers"

	"github.com/gin-gonic/gin"
)

func (a *App) routerRegister(ginGroup *gin.RouterGroup) {

	ginGroup.GET("/ping", handlers.Ping)
}
