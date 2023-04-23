package models

import (
	"DMAPI/app"
	"DMAPI/logger"
	"time"
)

type ContractorTechnics struct {
	UserID      int64      `json:"user_id,omitempty" db:"user_id,omitempty"`
	Name        string     `json:"name,omitempty" db:"name,omitempty"`
	Count       int64      `json:"count,omitempty" db:"count,omitempty"`
	TimeCreated *time.Time `json:"time_created,omitempty" db:"time_created,omitempty"`

	DealStatuses []DealStatus `json:"deal_statuses"`
}

func (m *ContractorTechnics) TableName() string {
	return "contractor_technics"
}
func GetContractorTechnics(userID int64) (s []ContractorTechnics) {
	err := app.DB.Select(&s, `select * from contractor_technics where user_id=$1`, userID)
	if err != nil {
		logger.Error.Println(err)
	}
	return s
}
