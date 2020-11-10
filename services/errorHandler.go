package services

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// ErrorHandler Xử lý lỗi cho server
func ErrorHandler(c *gin.Context, err error) {
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}
}
