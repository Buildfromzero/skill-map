package common

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type requestResponse struct {
	Message string `json:"message"`
	Status  uint   `json:"status"`
}

func SuccessResponse(ctx *gin.Context, msg string) {

	response := requestResponse{
		msg,
		http.StatusOK,
	}

	ctx.JSON(http.StatusOK, response)
}

func BadResponse(ctx *gin.Context, msg string) {

	response := requestResponse{
		msg,
		http.StatusBadRequest,
	}

	ctx.JSON(http.StatusBadRequest, response)
}
