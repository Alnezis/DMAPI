package photos

import (
	"DMAPI/controllers/api"
	"DMAPI/models"
	"github.com/CossackPyra/pyraconv"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type _cd struct {
	DealID  int64  `json:"deal_id"`
	URL     string `json:"url"`
	Caption string `json:"caption"`
}

func AddPhoto(c *gin.Context) {

	var req _cd

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

	if !models.ExistDEal(req.DealID) {
		c.JSON(http.StatusBadRequest, api.ResponseError("deal_id не существует", 4))
		return
	}

	p := models.Photo{
		DealID: req.DealID,
		UserID: userID,
		Url:    req.URL,

		TimeCreated: time.Time{},
	}

	p.Add()

	c.JSON(http.StatusOK, api.Response{Result: p})

}
