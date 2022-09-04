package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/gps/server/http"
)

func RegisterAssetsRoutes(r *gin.Engine) {
	authRoute := r.Group("/cdn")

	authRoute.GET(":sessionId/:textureCode", http.TextureAsset)
}