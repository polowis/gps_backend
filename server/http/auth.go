package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gps/pkg/auth"
)

func GetTextureSample(ctx *gin.Context) {
	authService := auth.NewAuth()
	authService.RegisterPWD()
	ctx.JSON(http.StatusOK, gin.H{"data": authService.Session(), "success": false})
}