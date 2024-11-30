package handlers

import (
	"Template/internal/http-server/models/ping"
	"Template/internal/transport/render"

	"github.com/gin-gonic/gin"
)

func Ping(c *gin.Context) {
	render.WriteResponse(c, ping.Get())
}
