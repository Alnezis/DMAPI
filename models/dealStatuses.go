package models

import (
	"DMAPI/app"
	"DMAPI/logger"
	"time"
)

type DealStatus struct {
	ID          int64      `json:"id,omitempty" db:"id,omitempty"`
	DealID      int64      `json:"deal_id,omitempty" db:"deal_id,omitempty"`
	Status      string     `json:"status,omitempty" db:"status,omitempty"`
	TimeCreated *time.Time `json:"time_created,omitempty" db:"time_created,omitempty"`
}

func (m *DealStatus) TableName() string {
	return "deal_statuses"
}

func GetDealStatuses(dealID int64) (s []DealStatus) {
	err := app.DB.Select(&s, `select id,status,time_created from deal_statuses where deal_id=$1 order by time_created desc`, dealID)
	if err != nil {
		logger.Error.Println(err)
	}
	return s
}

func (d *DealStatus) Create() (i int64) {

	stmt, err := app.DB.PrepareNamed(`INSERT INTO deal_statuses
    (deal_id, status)
VALUES (:deal_id, :status) returning id;`)

	if err != nil {
		logger.Error.Println(err)
	}

	err = stmt.Get(&i, &d)

	if err != nil {
		logger.Error.Println(err)
	}
	return
}
