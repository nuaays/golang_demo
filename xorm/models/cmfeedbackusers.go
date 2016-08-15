package models

import (
	"time"
)

type CmFeedbackUsers struct {
	Id         int       `json:"id" xorm:"not null pk autoincr INT(11)"`
	Phone      string    `json:"phone" xorm:"VARCHAR(11)"`
	Uuid       string    `json:"uuid" xorm:"unique VARCHAR(50)"`
	LastSubmit string    `json:"last_submit" xorm:"index(status) VARCHAR(45)"`
	City       string    `json:"city" xorm:"VARCHAR(45)"`
	Os         string    `json:"os" xorm:"VARCHAR(45)"`
	Sv         string    `json:"sv" xorm:"VARCHAR(45)"`
	Ov         string    `json:"ov" xorm:"VARCHAR(45)"`
	Mb         string    `json:"mb" xorm:"VARCHAR(45)"`
	Channel    string    `json:"channel" xorm:"VARCHAR(45)"`
	Status     int       `json:"status" xorm:"default 0 index(status) TINYINT(1)"`
	CreatedAt  time.Time `json:"created_at" xorm:"default 'CURRENT_TIMESTAMP' DATETIME"`
	UpdatedAt  time.Time `json:"updated_at" xorm:"default 'CURRENT_TIMESTAMP' DATETIME"`
}
