package contractors

import (
	"DMAPI/controllers/api"
	"DMAPI/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

//type req struct {
//	ContractorID string `form:"contractor_id"`
//}

func GetContractorsInfo(c *gin.Context) {

	//var req req
	//if err := c.ShouldBind(&req); err != nil {
	//	c.JSON(http.StatusBadRequest, api.ResponseError(err.Error(), 4))
	//	return
	//}

	d := models.GetContractors()
	c.JSON(http.StatusOK, api.Response{Result: d})

}
