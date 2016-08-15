package models

import (
	"time"
)

type CmUserBargain struct {
	Id            int       `json:"id" xorm:"not null pk autoincr INT(11)"`
	CarId         string    `json:"car_id" xorm:"VARCHAR(45)"`
	Price         string    `json:"price" xorm:"VARCHAR(45)"`
	Channel       string    `json:"channel" xorm:"VARCHAR(45)"`
	City          string    `json:"city" xorm:"VARCHAR(45)"`
	ProductType   string    `json:"product_type" xorm:"VARCHAR(45)"`
	TrafficSource string    `json:"traffic_source" xorm:"VARCHAR(45)"`
	Intent        string    `json:"intent" xorm:"VARCHAR(45)"`
	Date          time.Time `json:"date" xorm:"default 'CURRENT_TIMESTAMP' DATETIME"`
	Phone         string    `json:"phone" xorm:"VARCHAR(45)"`
	Entry         string    `json:"entry" xorm:"VARCHAR(45)"`
	UpdatedAt     time.Time `json:"updated_at" xorm:"not null default 'CURRENT_TIMESTAMP' TIMESTAMP"`
}
