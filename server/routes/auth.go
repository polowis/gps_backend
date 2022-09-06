package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/gps/server/http"
)

func RegisterAuthRoutes(r *gin.Engine) {
	authRoute := r.Group("/auth")

	authRoute.POST("texture", http.GetTextureSample)
	authRoute.POST("register", http.Register)
	authRoute.POST("verify", http.GetLoginPhotos)
}