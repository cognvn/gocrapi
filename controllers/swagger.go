package controllers

import (
	"bytes"
	"html/template"
	"net/http"

	"github.com/cognvn/gocrapi/services"
	"github.com/cognvn/gocrapi/swagger"
	"github.com/gin-gonic/gin"
)

// SwaggerController controller for SwaggerUI
func SwaggerController(c *gin.Context) {
	c.Writer.Header().Add("Content-Type", "text/html; charset=utf-8")

	assets, _ := swagger.Assets()

	tmpl, err := template.ParseFS(assets, "swagger.tmpl")
	services.ErrorHandler(c, err)
	var outHtml bytes.Buffer
	tmpl.Execute(&outHtml, gin.H{
		"name":        "Ocr server API",
		"url":         "/swagger/assets/swagger.json",
		"deepLinking": true,
		"theme":       "monokai",
	})
	c.String(http.StatusOK, outHtml.String())
}
