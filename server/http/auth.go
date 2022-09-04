package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gps/pkg/auth"
)

func GetTextureSample(ctx *gin.Context) {
	authService := auth.NewAuth()
	textures := authService.RegisterPWD()
	ctx.JSON(http.StatusOK, gin.H{
		"data": map[string]interface{}{
			"session": authService.Session(),
			"images": textures,
			"width": authService.TextureWidth,
			"height": authService.TextureHeight,
		}, 
		"success": false,
	})
}