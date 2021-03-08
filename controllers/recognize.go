package controllers

import (
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strings"

	"github.com/cognvn/gocrapi/services"
	"github.com/gin-gonic/gin"
	"github.com/otiai10/gosseract/v2"
)

// ResultModel Kết quả nhận diện
type ResultModel struct {
	Result string `json:"result"`
}

// RecognizeController Controller xử lý nhận diện ảnh thành chữ
func RecognizeController(c *gin.Context) {
	file, err := c.FormFile("file")
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
	// Handle recognize image to text
	client := gosseract.NewClient()
	defer client.Close()

	client.SetImage(tempfile.Name())
	client.Languages = []string{os.Getenv("DEFAULT_LANG")}
	if langs := c.PostForm("languages"); langs != "" {
		client.Languages = strings.Split(langs, ",")
	}

	var out string
	out, err = client.Text()
	services.ErrorHandler(c, err)

	c.JSON(http.StatusOK, ResultModel{
		Result: strings.Trim(out, " "),
	})
}
