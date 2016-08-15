package models

import (
	"time"
)

type CmUserVerify struct {
	UserId          int       `json:"user_id" xorm:"not null pk INT(11)"`
	Phone           string    `json:"phone" xorm:"index VARCHAR(45)"`
	Verifycode      string    `json:"verifycode" xorm:"VARCHAR(45)"`
	Timestamp       string    `json:"timestamp" xorm:"VARCHAR(45)"`
	Used            int       `json:"used" xorm:"default 0 INT(11)"`
	WrongTime       int       `json:"wrong_time" xorm:"default 0 INT(11)"`
	PhoneVerifycode string    `json:"phone_verifycode" xorm:"VARCHAR(45)"`
	CreatedAt       time.Time `json:"created_at" xorm:"DATETIME"`
	UpdatedAt       time.Time `json:"updated_at" xorm:"DATETIME"`
}
