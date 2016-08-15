package models

import (
	"time"
)

type CmUsers struct {
	Id            int       `json:"id" xorm:"not null pk autoincr INT(11)"`
	Phone         string    `json:"phone" xorm:"not null unique VARCHAR(45)"`
	RegistTime    time.Time `json:"regist_time" xorm:"default 'CURRENT_TIMESTAMP' DATETIME"`
	Origin        string    `json:"origin" xorm:"VARCHAR(45)"`
	MessageAuth   string    `json:"message_auth" xorm:"VARCHAR(45)"`
	AuthTimestamp string    `json:"auth_timestamp" xorm:"VARCHAR(45)"`
	Fail          int       `json:"fail" xorm:"default 0 INT(11)"`
	City          string    `json:"city" xorm:"VARCHAR(45)"`
	Channel       string    `json:"channel" xorm:"VARCHAR(45)"`
	Sv            string    `json:"sv" xorm:"VARCHAR(45)"`
	UniqueSym     string    `json:"unique_sym" xorm:"index VARCHAR(80)"`
	CurrentCity   string    `json:"current_city" xorm:"VARCHAR(45)"`
	Source        string    `json:"source" xorm:"VARCHAR(16)"`
	Action        string    `json:"action" xorm:"VARCHAR(16)"`
	CreatedAt     time.Time `json:"created_at" xorm:"DATETIME"`
	UpdatedAt     time.Time `json:"updated_at" xorm:"DATETIME"`
	UpdatedAt2    time.Time `json:"updated_at2" xorm:"default 'CURRENT_TIMESTAMP' DATETIME"`
}
