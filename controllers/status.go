package controllers

import (
	"net/http"

	"github.com/cognvn/ocrviet/services"
	"github.com/otiai10/gosseract/v2"

	"github.com/gin-gonic/gin"
)

// StatusController Controller kiểm tra trạng thái Tesseract trên server
func StatusController(c *gin.Context) {
	langs, err := gosseract.GetAvailableLanguages()
	services.ErrorHandler(c, err)

	client := gosseract.NewClient()
	defer client.Close()
	c.JSON(http.StatusOK, gin.H{
		"message": "Status OK",
		"tesseract": gin.H{
			"verion":    client.Version(),
			"languages": langs,
		},
	})
}
