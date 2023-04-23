package deals

import (
	"DMAPI/controllers/api"
	"DMAPI/models"
	"github.com/CossackPyra/pyraconv"
	"github.com/davecgh/go-spew/spew"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

type _cd struct {
	DealName string `json:"name"`

	ContractorID int64  `json:"contractor_id"`
	DocName      string `json:"doc_name"`
	DocURL       string `json:"doc_url"`
}

func CreateDeal(c *gin.Context) {

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

	if !models.ExistUserRole(req.ContractorID, "contractor") {
		c.JSON(http.StatusBadRequest, api.ResponseError("contractor_id не существует или не имеет роль contractor", 4))
		return
	}
	if req.DocURL == "" {
		c.JSON(http.StatusBadRequest, api.ResponseError("doc_url не существует", 4))
		return
	}

	if req.DealName == "" {
		req.DealName = "наименование сделки"
	}

	deal := models.Deal{
		UserID:   userID,
		ToUserID: req.ContractorID,
		Name:     req.DealName,
	}

	dealID := deal.CreateDeal()

	ds := models.DealStatus{
		DealID: dealID,
		Status: "pending",
		UserID: userID,
	}
	spew.Dump(ds)
	ds.Create()

	if req.DocName == "" {
		l := strings.Split(req.DocURL, "/")
		req.DocName = l[len(l)-1]
	}

	doc := models.Document{
		DealID:   dealID,
		UserID:   userID,
		ToUserID: req.ContractorID,
		Url:      req.DocURL,
		Name:     req.DocName,
	}

	docID := doc.CreateDocument()
	//spew.Dump(doc)
	//fmt.Println(docID)
	docStatus := models.DocumentStatuses{
		DocumentID: docID,
		Status:     "on_review",
		UserID:     userID,
	}

	docStatus.Create()

	d, _ := models.GetDeal(dealID)
	c.JSON(http.StatusOK, api.Response{Result: d})

}
