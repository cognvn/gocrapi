package controllers

import (
	"image"
	"io"
	"io/ioutil"
	"log"
	"os"

	"github.com/cognvn/ocrviet/services"
	"github.com/gin-gonic/gin"
)

// ImageController controller xử lý test chuyển đổi ảnh đen trắng
func ImageController(c *gin.Context) {
	file, err := c.FormFile("image")
	services.ErrorHandler(c, err)
	src, err := file.Open()
	services.ErrorHandler(c, err)
	defer src.Close()

	// Create physical file
	tempfile, err := ioutil.TempFile("", "ocrserver"+"-")
	services.ErrorHandler(c, err)
	defer func() {
		tempfile.Close()
		os.Remove(tempfile.Name())
	}()
	_, err = io.Copy(tempfile, src)
	services.ErrorHandler(c, err)
	c.File(tempfile.Name())
}

func grayscaleImage(c *gin.Context, img *os.File) {
	imgSrc, _, err := image.Decode(img)
	services.ErrorHandler(c, err)
	log.Println(imgSrc)
}
