package main

import (
	"log"

	"github.com/cognvn/ocrviet/controllers"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.Use(cors.Default())
	r.StaticFile("favicon.ico", "assets/favicon.ico")
	r.GET("/", controllers.StatusController)
	r.POST("/recognize", controllers.RecognizeController)
	r.Run()
	log.Println("Starting server")
}
