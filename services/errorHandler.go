package services

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// ErrorHandler Xử lý lỗi cho server
func ErrorHandler(c *gin.Context, err error) {
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
	}
}
