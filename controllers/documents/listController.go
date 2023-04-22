package documents

import (
	"DMAPI/controllers/api"
	"DMAPI/models"
	"github.com/CossackPyra/pyraconv"
	"github.com/gin-gonic/gin"
	"net/http"
)

type req struct {
	What string `uri:"what" binding:"required"`
}

func GetDocuments(c *gin.Context) {

	//	incoming outgoing

	var req req
	if err := c.ShouldBindUri(&req); err != nil {
		c.JSON(http.StatusBadRequest, api.ResponseError(err.Error(), 4))
		return
	}

	userID := pyraconv.ToInt64(c.GetHeader("USER-ID"))
	if userID == 0 {
		c.JSON(http.StatusBadRequest, api.ResponseError("Не передан USER-ID в заголовке", 4))
		return
	}

	var i []models.Document
	switch req.What {
	case "incoming":
		i = models.GetDocumentsIncoming(userID)

	case "outgoing":
		i = models.GetDocumentsOutgoing(userID)

	default:
		c.JSON(http.StatusBadRequest, api.ResponseError("status invalid: need incoming, outgoing", 4))
		return
	}

	c.JSON(http.StatusOK, api.Response{Result: i})

}
