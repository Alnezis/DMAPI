package messages

import (
	"DMAPI/controllers/api"
	"DMAPI/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

type _req struct {
	DealID int64 `form:"deal_id"`
}

func MessagesList(c *gin.Context) {

	var req _req
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusBadRequest, api.ResponseError(err.Error(), 4))
		return
	}

	if !models.ExistDEal(req.DealID) {
		c.JSON(http.StatusBadRequest, api.ResponseError("deal_id не существует", 4))
		return
	}

	d := models.GetDialogAllMessages(req.DealID)
	c.JSON(http.StatusOK, api.Response{Result: d})

}
