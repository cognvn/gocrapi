package main

import (
	"net/http"
	"time"

	"github.com/cognvn/ocrviet/controllers"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	// Cors config
	corsConf := cors.Config{
		AllowAllOrigins: true,
		// AllowOrigins:     []string{"https://gdt-bahnar.gitlab.io", "http://localhost:4005", "https://localhost:5001"},
		AllowMethods:     []string{"GET", "POST"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length", "Content-Type"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}
	r.Use(cors.New(corsConf))
	// Static file
	r.StaticFile("favicon.ico", "assets/favicon.ico")
	r.StaticFile("swagger.json", "assets/swagger.json")
	// Routers
	r.GET("/status", controllers.StatusController)
	r.POST("/recognize", controllers.RecognizeController)
	r.POST("/image", controllers.ImageController)
	r.NoRoute(func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "/status")
	})

	r.Run()
}
