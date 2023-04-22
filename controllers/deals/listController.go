package deals

import (
	"DMAPI/controllers/api"
	"DMAPI/models"
	"github.com/CossackPyra/pyraconv"
	"github.com/gin-gonic/gin"
	"net/http"
)

type req struct {
	What string `form:"what"`
}

func GetDeals(c *gin.Context) {

	//	incoming outgoing

	//var req req
	//if err := c.ShouldBind(&req); err != nil {
	//	c.JSON(http.StatusBadRequest, api.ResponseError(err.Error(), 4))
	//	return
	//}

	userID := pyraconv.ToInt64(c.GetHeader("USER-ID"))
	if userID == 0 {
		c.JSON(http.StatusBadRequest, api.ResponseError("Не передан USER-ID в заголовке", 4))
		return
	}

	d := models.GetDeals(userID)

	c.JSON(http.StatusOK, api.Response{Result: d})

}
