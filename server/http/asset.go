package http

import "github.com/gin-gonic/gin"


/*
Sample URL /:sessionId/:textureCode
*/
func TextureAsset(ctx *gin.Context) {
	sessionId := ctx.Param("sessionId")
	textureCode := ctx.Param("textureCode")

	ctx.File("storage/sp/" + sessionId + "/" + textureCode + ".png")
}