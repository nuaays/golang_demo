package models

import (
	"time"
)

type CmUserScanRecord struct {
	Id        int       `json:"id" xorm:"not null pk autoincr INT(11)"`
	Phone     string    `json:"phone" xorm:"index(pd) VARCHAR(45)"`
	CarId     string    `json:"car_id" xorm:"index VARCHAR(45)"`
	Os        string    `json:"os" xorm:"VARCHAR(45)"`
	Ip        string    `json:"ip" xorm:"VARCHAR(45)"`
	ScanTime  time.Time `json:"scan_time" xorm:"default 'CURRENT_TIMESTAMP' index DATETIME"`
	Disabled  int       `json:"disabled" xorm:"default 0 index(pd) index INT(11)"`
	CreatedAt time.Time `json:"created_at" xorm:"DATETIME"`
	UpdatedAt time.Time `json:"updated_at" xorm:"DATETIME"`
}
