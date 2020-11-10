package main

import (
	"net/http"

	"github.com/cognvn/ocrviet/controllers"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.Use(cors.Default())
	r.StaticFile("favicon.ico", "assets/favicon.ico")
	r.StaticFile("swagger.json", "assets/swagger.json")
	r.GET("/status", controllers.StatusController)
	r.POST("/recognize", controllers.RecognizeController)
	r.NoRoute(func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "/status")
	})

	r.Run()
}
