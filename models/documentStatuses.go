package models

import (
	"DMAPI/app"
	"DMAPI/logger"
	"time"
)

type DocumentStatuses struct {
	ID          int64      `json:"id,omitempty" db:"id,omitempty"`
	DocumentID  int64      `json:"document_id,omitempty" db:"document_id,omitempty"`
	Status      string     `json:"status,omitempty" db:"status,omitempty"`
	TimeCreated *time.Time `json:"time_created,omitempty" db:"time_created,omitempty"`
}

func (m *DocumentStatuses) TableName() string {
	return "document_statuses"
}

func GetDocumentStatuses(docID int64) (s []DocumentStatuses) {
	err := app.DB.Select(&s, `select id,status,time_created from document_statuses where document_id=$1 order by time_created desc`, docID)
	if err != nil {
		logger.Error.Println(err)
	}
	return s
}

func (d *DocumentStatuses) Create() (i int64) {

	stmt, err := app.DB.PrepareNamed(`INSERT INTO document_statuses
    (document_id, status)
VALUES (:document_id, :status) returning id;`)
	if err != nil {
		logger.Error.Println(err)
	}

	err = stmt.Get(&i, &d)

	if err != nil {
		logger.Error.Println(err)
	}
	return
}
