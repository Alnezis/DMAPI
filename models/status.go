package models

import (
	"DMAPI/app"
	"DMAPI/logger"
)

type Status struct {
	Key   string `json:"key,omitempty" db:"key,omitempty"`
	Name  string `json:"name,omitempty" db:"name,omitempty"`
	Color string `json:"color,omitempty" db:"color,omitempty"`
}

func (m *Status) TableName() string {
	return "statuses"
}

func GetStatuses() map[string]Status {
	var ss []Status
	err := app.DB.Select(&ss, `select * from statuses s order by s.key desc`)
	if err != nil {
		logger.Error.Println(err)
	}

	var sss = map[string]Status{}
	for _, status := range ss {
		sss[status.Key] = status
	}
	return sss
}

func ExistStatus(key string) (exist bool) {
	err := app.DB.Get(&exist, `select exists(select * from statuses s where s.key=$1)`, key)
	if err != nil {
		logger.Error.Println(err)
	}
	return
}
