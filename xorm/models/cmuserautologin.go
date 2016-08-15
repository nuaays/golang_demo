package models

import (
	"time"
)

type CmUserAutologin struct {
	Id            int       `json:"id" xorm:"not null pk autoincr INT(11)"`
	Token         string    `json:"token" xorm:"unique VARCHAR(45)"`
	Phone         string    `json:"phone" xorm:"VARCHAR(45)"`
	Os            string    `json:"os" xorm:"VARCHAR(45)"`
	Timestamp     string    `json:"timestamp" xorm:"VARCHAR(45)"`
	TokenDisabled int       `json:"token_disabled" xorm:"default 0 INT(11)"`
	LastIp        string    `json:"last_ip" xorm:"VARCHAR(45)"`
	UserId        int       `json:"user_id" xorm:"INT(11)"`
	CreatedAt     time.Time `json:"created_at" xorm:"DATETIME"`
	UpdatedAt     time.Time `json:"updated_at" xorm:"DATETIME"`
}
