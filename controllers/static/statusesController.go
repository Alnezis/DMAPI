package static

import (
	"DMAPI/controllers/api"
	"DMAPI/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetStatuses(c *gin.Context) {
	d := models.GetStatuses()
	c.JSON(http.StatusOK, api.Response{Result: d})
}
