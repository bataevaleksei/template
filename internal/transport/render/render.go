package render

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type ApiError struct {
	ResultCode uint32 `json:"result_code"`
	Data       any    `json:"data,omitempty"`
}

func WriteResponse(ctx *gin.Context, data any) {
	ctx.JSON(http.StatusOK, data)
}
func WriteError(ctx *gin.Context, code int, apiError ApiError) {
	ctx.JSON(code, apiError)
}
