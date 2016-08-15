package models

import (
	"time"
)

type CmUserSmsToPhoneRecord struct {
	Id         int       `json:"id" xorm:"not null pk autoincr INT(11)"`
	OwnerPhone string    `json:"owner_phone" xorm:"not null VARCHAR(45)"`
	SmsDate    time.Time `json:"sms_date" xorm:"default 'CURRENT_TIMESTAMP' DATETIME"`
	CarId      string    `json:"car_id" xorm:"VARCHAR(45)"`
	UpdatedAt  time.Time `json:"updated_at" xorm:"not null default 'CURRENT_TIMESTAMP' TIMESTAMP"`
}
