package documents

import (
	"DMAPI/controllers/api"
	"DMAPI/models"
	"github.com/CossackPyra/pyraconv"
	"github.com/gin-gonic/gin"
	"net/http"
)

type _cd_ struct {
	DocID     int64  `json:"doc_id"`
	StatusKey string `json:"status_key"`
}

func SetStatusDoc(c *gin.Context) {

	var req _cd_

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, api.ResponseError(err.Error(), 3))
		return
	}
	userID := pyraconv.ToInt64(c.GetHeader("USER-ID"))
	if userID == 0 {
		c.JSON(http.StatusBadRequest, api.ResponseError("Не передан USER-ID в заголовке", 4))
		return
	}

	if !models.ExistUser(userID) {
		c.JSON(http.StatusBadRequest, api.ResponseError("USER-ID не существует", 4))
		return
	}

	if !models.ExistDoc(req.DocID) {
		c.JSON(http.StatusBadRequest, api.ResponseError("doc_id не существует", 4))
		return
	}

	if !models.ExistStatus(req.StatusKey) {
		c.JSON(http.StatusBadRequest, api.ResponseError("status_key не существует, полный список по роуту /static/statuses", 4))
		return
	}

	s := models.DocumentStatuses{
		DocumentID: req.DocID,
		UserID:     userID,
		Status:     req.StatusKey,
	}
	sID := s.Create()

	d := models.GetDocStatus(sID)

	c.JSON(http.StatusOK, api.Response{Result: d})

}
