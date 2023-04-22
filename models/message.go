package models

import (
	"DMAPI/app"
	"DMAPI/logger"
	"time"
)

type Message struct {
	ID          int64     `json:"id,omitempty" db:"id,omitempty"`
	DealID      int64     `json:"deal_id,omitempty" db:"deal_id,omitempty"`
	UserID      int64     `json:"user_id,omitempty" db:"user_id,omitempty"`
	Text        string    `json:"text,omitempty" db:"text,omitempty"`
	TimeCreated time.Time `json:"time_created,omitempty" db:"time_created,omitempty"`

	Name string `json:"name" db:"name"`
}

func (m *Message) TableName() string {
	return "messages"
}

func (d *Message) Send() (i int64) {

	stmt, err := app.DB.PrepareNamed(`INSERT INTO messages
    (deal_id, user_id, text) 
VALUES (:deal_id, :user_id, :text) returning id;`)
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

func GetDialogAllMessages(dealID int64) (s []Message) {
	err := app.DB.Select(&s, `select m.text,m.id,m.user_id,m.time_created,u.name from messages m join users u on u.id = m.user_id where m.deal_id=$1 order by m.time_created desc`, dealID)
	if err != nil {
		logger.Error.Println(err)
	}
	return s
}
