package main

import (
	"net/http"
	"time"

	"github.com/cognvn/gocrapi/controllers"
	"github.com/cognvn/gocrapi/swagger"
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
	r.Static("/.well-known", "./.well-known")
	// Routers
	r.GET("/favicon.ico", controllers.FaviconController)
	r.GET("/status", controllers.StatusController)
	r.POST("/recognize", controllers.RecognizeController)

	// Swagger
	swgGroup := r.Group("/swagger")
	{
		swgGroup.GET("/", controllers.SwaggerController)
		swgGroup.StaticFS("/assets", swagger.AssetFS())
	}
	r.NoRoute(func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "/swagger")
	})

	r.Run()
}
