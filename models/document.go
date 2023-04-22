package models

import (
	"DMAPI/app"
	"DMAPI/logger"
	"time"
)

type Document struct {
	ID          int64      `json:"id,omitempty" db:"id,omitempty"`
	DealID      int64      `json:"deal_id,omitempty" db:"deal_id,omitempty"`
	UserID      int64      `json:"user_id,omitempty" db:"user_id,omitempty"`
	ToUserID    int64      `json:"to_user_id,omitempty" db:"to_user_id,omitempty"`
	Url         string     `json:"url,omitempty" db:"url,omitempty"`
	Name        string     `json:"name,omitempty" db:"name,omitempty"`
	TimeCreated *time.Time `json:"time_created,omitempty" db:"time_created,omitempty"`

	DocumentStatuses []DocumentStatuses `json:"document_statuses"`
}

func (m *Document) TableName() string {
	return "documents"
}

func (d *Document) CreateDocument() (i int64) {

	stmt, err := app.DB.PrepareNamed(`INSERT INTO documents
    (deal_id, user_id, to_user_id, url, name)
VALUES (:deal_id, :user_id, :to_user_id, :url, :name) returning id;`)
	if err != nil {
		logger.Error.Println(err)
	}

	err = stmt.Get(&i, &d)

	if err != nil {
		logger.Error.Println(err)
	}
	return
}

func GetDocuments(dealID int64) (ci []Document) {
	var r Document

	rows, err := app.DB.Queryx(`select id,user_id, to_user_id, url, name from documents where deal_id=$1`, dealID)
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

		r.DocumentStatuses = GetDocumentStatuses(r.ID)

		ci = append(ci, r)

	}

	return ci

}

func GetDocumentsIncoming(userID int64) (ci []Document) {
	var r Document

	rows, err := app.DB.Queryx(`select * from documents where to_user_id=$1 order by time_created desc `, userID)
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

		r.DocumentStatuses = GetDocumentStatuses(r.ID)

		ci = append(ci, r)

	}

	return ci

}

func GetDocumentsOutgoing(userID int64) (ci []Document) {
	var r Document

	rows, err := app.DB.Queryx(`select * from documents where user_id=$1 order by time_created desc `, userID)
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

		r.DocumentStatuses = GetDocumentStatuses(r.ID)

		ci = append(ci, r)

	}

	return ci

}
