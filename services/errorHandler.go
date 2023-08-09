package services

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// ErrorHandler Xử lý lỗi cho server
func ErrorHandler(c *gin.Context, err error) {
	if err != nil {
		log.Println(err.Error())
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}
}
