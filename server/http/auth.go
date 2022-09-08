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
		"success": true,
	})
}

func Register(ctx *gin.Context) {
	var registerRequest auth.RegisterRequest

	if err := ctx.BindJSON(&registerRequest); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"data": map[string]interface{}{
				"error": "Not allowed",
				"success": false,
			},
		})
	}
	authService := auth.NewAuth()
	err := authService.Register(registerRequest)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"data": map[string]interface{}{
				"error": "Not allowed",
				"success": false,
			},
		})
		return
	}

	authService.ClearSessionTexture(registerRequest.Session)

	ctx.JSON(http.StatusOK, gin.H{
		"data": map[string]interface{}{
			"redirect_url": "http://localhost:3000/success",
			"success": true,
		},
	})

}

type RequestPhotoPhase struct {
	Email   string  `json:"email"`
	Order   string  `json:"order"`
	Session string `json:"session"`
}

// verify email and password to get photo signature
func GetLoginPhotos(ctx *gin.Context) {
	var verifyRequest RequestPhotoPhase
	if err := ctx.BindJSON(&verifyRequest); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"data": map[string]interface{}{
				"error": "Not allowed",
				"success": false,
			},
		})
	}
	authService := auth.NewAuth()
	authService.VerifyUser(verifyRequest.Email, verifyRequest.Order, verifyRequest.Session)
	authService.ClearSessionTexture(verifyRequest.Session)
}