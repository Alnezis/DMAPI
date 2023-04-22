package files

import (
	"DMAPI/app"
	"DMAPI/controllers/api"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
)

type getFilename struct {
	Filename string `uri:"filename" binding:"required"`
}

func Files(c *gin.Context) {

	var req getFilename
	if err := c.ShouldBindUri(&req); err != nil {
		c.JSON(http.StatusBadRequest, api.ResponseError(err.Error(), 4))
		return
	}

	path := app.CFG.PathFiles + req.Filename

	if _, err := os.Stat(path); err != nil {
		if os.IsNotExist(err) {
			c.JSON(http.StatusBadRequest, api.ResponseError(err.Error(), 3))
			return
		}
	} else {
		c.Writer.Header().Add("Content-Disposition", fmt.Sprintf("attachment; filename=%s", req.Filename))
		c.File(path)
	}
}
