package files

import (
	"DMAPI/app"
	"DMAPI/controllers/api"
	"DMAPI/logger"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"os"
	"strings"
)

func HandlePost(c *gin.Context) {

	// Parse request body as multipart form data with 32MB max memory
	err := c.Request.ParseMultipartForm(32 << 20)
	if err != nil {
		logger.Error.Println(err)
	}

	// Get file from Form
	file, handler, err := c.Request.FormFile("file")
	if err != nil {
		logger.Error.Println(err)
		c.JSON(http.StatusInternalServerError, api.ResponseError(".errors.upload_image.cannot_read_file", 3))
		return
	}
	defer file.Close()
	fn := api.RandString(5) + "_" + strings.Replace(handler.Filename, " ", "", -1)
	// Add file locally
	dst, err := os.Create(app.CFG.PathFiles + fn)
	if err != nil {
		logger.Error.Println(err)
		c.JSON(http.StatusInternalServerError, api.ResponseError(".errors.upload_image.cannot_create_local_file", 3))
		return
	}
	defer dst.Close()

	// Copy the uploaded file data to the newly created file on the filesystem
	if _, err := io.Copy(dst, file); err != nil {
		logger.Error.Println(err)
		c.JSON(http.StatusInternalServerError, api.ResponseError(".errors.upload_image.cannot_copy_to_file", 3))
		return
	}
	c.JSON(http.StatusOK, api.Response{Result: app.CFG.FilesUrl + fn})

}
