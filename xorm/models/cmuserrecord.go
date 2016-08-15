package models

import (
	"time"
)

type CmUserRecord struct {
	Id             int       `json:"id" xorm:"not null pk autoincr INT(11)"`
	UserId         int       `json:"user_id" xorm:"index INT(11)"`
	RecordType     string    `json:"record_type" xorm:"VARCHAR(45)"`
	RecordData     string    `json:"record_data" xorm:"index VARCHAR(500)"`
	RecordTime     time.Time `json:"record_time" xorm:"default 'CURRENT_TIMESTAMP' index DATETIME"`
	Disabled       int       `json:"disabled" xorm:"default 0 index INT(11)"`
	Judge          string    `json:"judge" xorm:"not null VARCHAR(45)"`
	LastSubmitTime string    `json:"last_submit_time" xorm:"VARCHAR(45)"`
	DeviceType     string    `json:"device_type" xorm:"VARCHAR(20)"`
	CreatedAt      time.Time `json:"created_at" xorm:"DATETIME"`
	UpdatedAt      time.Time `json:"updated_at" xorm:"index DATETIME"`
	Ep             string    `json:"ep" xorm:"VARCHAR(4096)"`
}
