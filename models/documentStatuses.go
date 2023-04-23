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
	UserID      int64      `json:"user_id,omitempty" db:"user_id,omitempty"`
	TimeCreated *time.Time `json:"time_created,omitempty" db:"time_created,omitempty"`

	Name       string `json:"name,omitempty" db:"name,omitempty"`
	StatusName string `json:"status_name,omitempty" db:"status_name,omitempty"`
}

func (m *DocumentStatuses) TableName() string {
	return "document_statuses"
}

func GetDocumentStatuses(docID int64) (s []DocumentStatuses) {
	err := app.DB.Select(&s, `select s.*, u.name, s2.name as status_name from 
                                          document_statuses s join users u on u.id = s.user_id
                                      join statuses s2 on s.status = s2.key

                                      where document_id=$1 order by time_created desc`, docID)
	if err != nil {
		logger.Error.Println(err)
	}
	return s
}

func GetDocStatus(statusID int64) (s DocumentStatuses) {
	err := app.DB.Get(&s, `select s.*, u.name, s2.name as status_name from 
                                          document_statuses s join users u on u.id = s.user_id
                                      join statuses s2 on s.status = s2.key

                                      where s.id=$1`, statusID)
	if err != nil {
		logger.Error.Println(err)
	}
	return
}

func (d *DocumentStatuses) Create() (i int64) {

	stmt, err := app.DB.PrepareNamed(`INSERT INTO document_statuses
    (document_id, status,user_id)
VALUES (:document_id, :status, :user_id) returning id;`)
	if err != nil {
		logger.Error.Println(err)
	}

	err = stmt.Get(&i, &d)

	if err != nil {
		logger.Error.Println(err)
	}
	return
}
