package deals

import (
	"DMAPI/controllers/api"
	"DMAPI/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

type _req struct {
	DealID int64 `form:"deal_id"`
}

func GetDeal(c *gin.Context) {

	var req _req
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusBadRequest, api.ResponseError(err.Error(), 4))
		return
	}

	if !models.ExistDeal(req.DealID) {
		c.JSON(http.StatusBadRequest, api.ResponseError("deal_id не существует", 4))
		return
	}

	d, _ := models.GetDeal(req.DealID)
	c.JSON(http.StatusOK, api.Response{Result: d})

}
