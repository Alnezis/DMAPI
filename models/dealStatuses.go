package models

import (
	"DMAPI/app"
	"DMAPI/logger"
	"time"
)

type DealStatus struct {
	ID          int64     `json:"id,omitempty" db:"id,omitempty"`
	DealID      int64     `json:"deal_id,omitempty" db:"deal_id,omitempty"`
	UserID      int64     `json:"user_id,omitempty" db:"user_id,omitempty"`
	Status      string    `json:"status,omitempty" db:"status,omitempty"`
	TimeCreated time.Time `json:"time_created,omitempty" db:"time_created,omitempty"`

	Name       string `json:"name,omitempty" db:"name,omitempty"`
	StatusName string `json:"status_name,omitempty" db:"status_name,omitempty"`
}

func (m *DealStatus) TableName() string {
	return "deal_statuses"
}

func GetDealStatuses(dealID int64) (s []DealStatus) {
	err := app.DB.Select(&s, `select s.*, u.name, s2.name as status_name from
		deal_statuses s join users u on u.id = s.user_id
                   join statuses s2 on s.status = s2.key
		 where deal_id=$1 order by time_created desc`, dealID)
	if err != nil {
		logger.Error.Println(err)
	}
	return s
}

func GetDealStatus(statusID int64) (s DealStatus) {
	err := app.DB.Get(&s, `select s.*, u.name, s2.name as status_name from 
                                          deal_statuses s join users u on u.id = s.user_id
                   join statuses s2 on s.status = s2.key

                                      where s.id=$1`, statusID)
	if err != nil {
		logger.Error.Println(err)
	}
	return
}

func (d *DealStatus) Create() (i int64) {

	stmt, err := app.DB.PrepareNamed(`INSERT INTO deal_statuses
    (deal_id, status,user_id)
VALUES (:deal_id, :status, :user_id) returning id;`)

	if err != nil {
		logger.Error.Println(err)
	}

	err = stmt.Get(&i, &d)

	if err != nil {
		logger.Error.Println(err)
	}

	return
}
