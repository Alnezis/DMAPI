package models

import (
	"DMAPI/app"
	"DMAPI/logger"
	"time"
)

type ContractorInfo struct {
	UserID          int64      `json:"user_id,omitempty" db:"user_id,omitempty"`
	ExperienceYears int64      `json:"experience_years,omitempty" db:"experience_years,omitempty"`
	ExistTechnic    bool       `json:"exist_technic,omitempty" db:"exist_technic,omitempty"`
	Specification   string     `json:"specification,omitempty" db:"specification,omitempty"`
	FeaturesList    string     `json:"features_list,omitempty" db:"features_list,omitempty"`
	TimeCreated     *time.Time `json:"time_created,omitempty" db:"time_created,omitempty"`

	Name string `json:"name,omitempty" db:"name,omitempty"`

	Technics []ContractorTechnics `json:"technics"`
}

func (m *ContractorInfo) TableName() string {
	return "contractor_info"
}

func GetContractors() (ci []ContractorInfo) {
	var r ContractorInfo

	rows, err := app.DB.Queryx(`select contractor_info.*, u.name from contractor_info inner join users u on u.id = contractor_info.user_id order by time_created desc`)
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

		if r.ExistTechnic {
			r.Technics = GetContractorTechnics(r.UserID)
		}

		ci = append(ci, r)
	}

	return ci

}
