package models

import (
	"time"
)

type CmAlterPriceLog struct {
	Id          int       `json:"id" xorm:"not null pk autoincr INT(11)"`
	CarId       int       `json:"car_id" xorm:"INT(11)"`
	Price       float32   `json:"price" xorm:"FLOAT"`
	AdjustPrice float32   `json:"adjust_price" xorm:"FLOAT"`
	Source      string    `json:"source" xorm:"CHAR(32)"`
	BusiStatus  int       `json:"busi_status" xorm:"INT(11)"`
	Timestamp   time.Time `json:"timestamp" xorm:"default 'CURRENT_TIMESTAMP' DATETIME"`
	UpdatedAt   time.Time `json:"updated_at" xorm:"not null default 'CURRENT_TIMESTAMP' TIMESTAMP"`
}
