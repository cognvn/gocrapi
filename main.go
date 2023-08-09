package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
	"path"
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
	swaggerAssets, _ := swagger.Assets()
	swgGroup := r.Group("/swagger")
	{
		swgGroup.GET("/", controllers.SwaggerController)
		swgGroup.StaticFS("/assets", http.FS(swaggerAssets))
	}
	r.NoRoute(func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "/swagger")
	})

	// Get port for run
	envPort, portOk := os.LookupEnv("PORT")
	if !portOk {
		envPort = "8080"
	}
	port := flag.String("port", envPort, "Khởi chạy ứng dụng tại cổng")
	host := flag.String("host", "localhost", "Địa chỉ khởi chạy ứng dụng")
	url := flag.String("url", "", "Địa chỉ khởi chạy ứng dụng")
	flag.Parse()
	if *url == "" {
		*url = fmt.Sprintf("%s:%s", *host, *port)
	}

	_, preOk := os.LookupEnv("TESSDATA_PREFIX")
	if !preOk {
		cwd, _ := os.Getwd()
		os.Setenv("TESSDATA_PREFIX", path.Join(cwd, "tessdata"))
	}

	r.Run(*url)
}
