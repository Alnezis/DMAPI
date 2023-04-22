package models

import (
	"DMAPI/app"
	"DMAPI/logger"
	"time"
)

type Deal struct {
	ID          int64     `json:"id,omitempty" db:"id,omitempty"`
	UserID      int64     `json:"user_id,omitempty" db:"user_id,omitempty"`
	ToUserID    int64     `json:"to_user_id,omitempty" db:"to_user_id,omitempty"`
	Name        string    `json:"name,omitempty" db:"name,omitempty"`
	TimeCreated time.Time `json:"time_created,omitempty" db:"time_created,omitempty"`

	DealStatuses []DealStatus `json:"deal_statuses"`
	Documents    []Document   `json:"documents"`
	Photos       []Photo      `json:"photos"`
}

func (d *Deal) TableName() string {
	return "deals"
}

func (d *Deal) CreateDeal() (i int64) {

	stmt, err := app.DB.PrepareNamed(`INSERT INTO deals
	(user_id, to_user_id, name)
	VALUES (:user_id, :to_user_id, :name) returning id;`)
	if err != nil {
		logger.Error.Println(err)
	}

	err = stmt.Get(&i, &d)

	if err != nil {
		logger.Error.Println(err)
	}
	return
}

func GetDeals(userID int64) (ci []Deal) {
	var r Deal

	rows, err := app.DB.Queryx(`select id,to_user_id,time_created from deals where user_id=$1 order by time_created desc`, userID)
	if err != nil {
		logger.Error.Println(err)
	} else {
		defer rows.Close()
	}

	for rows.Next() {
		err = rows.StructScan(&r)
		if err != nil {
			logger.Error.Println(err)
		}

		r.DealStatuses = GetDealStatuses(r.ID)
		r.Documents = GetDocuments(r.ID)
		r.Photos = GetPhotosDeal(r.ID)

		ci = append(ci, r)
	}

	return ci

}

func ExistDEal(dealID int64) (exist bool) {
	err := app.DB.Get(&exist, `select exists(select * from deals where id=$1)`, dealID)
	if err != nil {
		logger.Error.Println(err)
	}
	return
}

func GetDeal(dealID int64) (ci Deal) {

	err := app.DB.Get(&ci, `select * from deals where id=$1 order by time_created desc`, dealID)
	if err != nil {
		logger.Error.Println(err)
	}

	ci.DealStatuses = GetDealStatuses(ci.ID)
	ci.Documents = GetDocuments(ci.ID)
	ci.Photos = GetPhotosDeal(ci.ID)

	return ci

}
