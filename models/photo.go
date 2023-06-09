package models

import (
	"DMAPI/app"
	"DMAPI/logger"
	"time"
)

type Photo struct {
	ID      int64  `json:"id,omitempty" db:"id,omitempty"`
	DealID  int64  `json:"deal_id,omitempty" db:"deal_id,omitempty"`
	UserID  int64  `json:"user_id,omitempty" db:"user_id,omitempty"`
	Url     string `json:"url" db:"url,omitempty"`
	Caption string `json:"caption" db:"caption,omitempty"`

	TimeCreated time.Time `json:"time_created,omitempty" db:"time_created,omitempty"`
	Name        string    `json:"name,omitempty" db:"name,omitempty"`
}

func (m *Photo) TableName() string {
	return "photos"
}

func GetPhotosDeal(dealID int64) (s []Photo) {
	err := app.DB.Select(&s, `select p.id,p.user_id,p.time_created,caption,url,u.name from photos p 
                                           join users u on u.id = p.user_id
                                           where deal_id=$1 order by time_created desc`, dealID)
	if err != nil {
		logger.Error.Println(err)
	}
	return s
}

func (d *Photo) Add() (i int64) {

	stmt, err := app.DB.PrepareNamed(`INSERT INTO photos
    (deal_id, user_id, url, caption) 
VALUES (:deal_id, :user_id, :url, :caption) returning id;`)
	if err != nil {
		logger.Error.Println(err)
	}

	err = stmt.Get(&i, &d)

	if err != nil {
		logger.Error.Println(err)
	}

	d.ID = i
	d.TimeCreated = time.Now()
	return
}
