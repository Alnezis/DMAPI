package deals

import (
	"DMAPI/controllers/api"
	"DMAPI/models"
	"github.com/CossackPyra/pyraconv"
	"github.com/gin-gonic/gin"
	"net/http"
)

type _cd_ struct {
	DealID    int64  `json:"deal_id"`
	StatusKey string `json:"status_key"`
}

func SetStatusDeal(c *gin.Context) {

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

	if !models.ExistDeal(req.DealID) {
		c.JSON(http.StatusBadRequest, api.ResponseError("deal_id не существует", 4))
		return
	}

	if !models.ExistStatus(req.StatusKey) {
		c.JSON(http.StatusBadRequest, api.ResponseError("status_key не существует, полный список по роуту /static/statuses", 4))
		return
	}

	s := models.DealStatus{
		DealID: req.DealID,
		UserID: userID,
		Status: req.StatusKey,
	}
	sID := s.Create()

	d := models.GetDealStatus(sID)

	c.JSON(http.StatusOK, api.Response{Result: d})

}
