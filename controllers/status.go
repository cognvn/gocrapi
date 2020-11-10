package controllers

import (
	"net/http"

	"github.com/cognvn/ocrviet/services"
	"github.com/otiai10/gosseract/v2"

	"github.com/gin-gonic/gin"
)

// TessServerModel Trạng thái Tesseract-OCR đã cài đặt
type TessServerModel struct {
	Version   string   `json:"version"`
	Languages []string `json:"languages"`
}

// StatusModel Trạng thái server
type StatusModel struct {
	Message   string          `json:"message"`
	Tesseract TessServerModel `json:"tesseract"`
}

// StatusController Controller kiểm tra trạng thái Tesseract trên server
func StatusController(c *gin.Context) {
	langs, err := gosseract.GetAvailableLanguages()
	services.ErrorHandler(c, err)

	client := gosseract.NewClient()
	defer client.Close()
	c.JSON(http.StatusOK, StatusModel{
		Message: "Status OK",
		Tesseract: TessServerModel{
			Version:   client.Version(),
			Languages: langs,
		},
	})
}
