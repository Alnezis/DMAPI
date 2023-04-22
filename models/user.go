package models

import (
	"DMAPI/app"
	"DMAPI/logger"
	"fmt"
	"time"
)

type User struct {
	ID       int64  `json:"id,omitempty" db:"id,omitempty"`
	Email    string `json:"email,omitempty" db:"email,omitempty"`
	Password string `json:"password,omitempty" db:"password,omitempty"`
	Role     string `json:"role,omitempty" db:"user_role,omitempty"`
	Name     string `json:"name,omitempty" db:"name,omitempty"`

	TimeCreated *time.Time `json:"time_created,omitempty" db:"time_created,omitempty"`
}

func (m *User) TableName() string {
	return "users"
}

func Auth(l, p string) *User {

	var user User
	var exist bool

	err := app.DB.Get(&exist, `select exists(select * from users where email=$1 and password=$2)`, l, p)
	if err != nil {
		logger.Error.Println(err)
	}
	if !exist {
		return nil
	}
	fmt.Println(l, p)
	err = app.DB.Get(&user, `select id,user_role,name from users where (email=$1 and password=$2)`, l, p)
	if err != nil {
		logger.Error.Println(err)
	}
	return &user
}

func ExistUser(userID int64) (exist bool) {
	err := app.DB.Get(&exist, `select exists(select * from users where id=$1)`, userID)
	if err != nil {
		logger.Error.Println(err)
	}
	return
}

func ExistUserRole(userID int64, role string) (exist bool) {
	err := app.DB.Get(&exist, "select exists(select * from users where (user_role=$1 and id=$2))", role, userID)
	if err != nil {
		logger.Error.Println(err)
	}
	return
}
