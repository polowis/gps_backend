package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	//"github.com/gps/server/models"
	"github.com/gps/conf"
	"github.com/gps/server/routes"
	//"gorm.io/driver/sqlite"
	//"gorm.io/gorm"
)

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")
  
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
  
		c.Next()
	}
}

func init() {
	fmt.Println("Setting up...")
	conf.Setup("conf/app.ini")
	fmt.Println("Set up completed!")
}

func main() {
	/*db, err := gorm.Open(sqlite.Open("main.db"), &gorm.Config{})
	if err != nil {
	  panic("failed to connect database")
	}
  
	// Migrate the schema
	db.AutoMigrate(&models.User{})*/

	r := gin.Default()
	r.Use(CORSMiddleware())

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
		  "data": map[string]interface{}  {
			"source": "GSP",
			"version": "1.0",
		  },
		})    
	})

	routes.RegisterAuthRoutes(r)
	routes.RegisterAssetsRoutes(r)
	r.Run(":8090")
}