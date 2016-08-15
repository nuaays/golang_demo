package models

import (
	"time"
)

type CmUserIntent struct {
	Id         int       `json:"id" xorm:"not null pk autoincr INT(10)"`
	Device     string    `json:"device" xorm:"not null VARCHAR(45)"`
	Product    string    `json:"product" xorm:"not null VARCHAR(45)"`
	Phone      string    `json:"phone" xorm:"not null index VARCHAR(45)"`
	Intent     string    `json:"intent" xorm:"not null VARCHAR(45)"`
	RegistTime time.Time `json:"regist_time" xorm:"index DATETIME"`
	City       string    `json:"city" xorm:"VARCHAR(45)"`
	Origin     string    `json:"origin" xorm:"VARCHAR(45)"`
	Ip         string    `json:"ip" xorm:"VARCHAR(45)"`
	CarId      string    `json:"car_id" xorm:"index VARCHAR(45)"`
	BrandId    string    `json:"brand_id" xorm:"VARCHAR(45)"`
	SeriesId   string    `json:"series_id" xorm:"VARCHAR(45)"`
	ModelId    string    `json:"model_id" xorm:"VARCHAR(45)"`
	Extends    string    `json:"extends" xorm:"VARCHAR(2048)"`
	Price      float32   `json:"price" xorm:"FLOAT(10,2)"`
	SupplyFlag int       `json:"supply_flag" xorm:"default 0 INT(11)"`
	UpdatedAt  time.Time `json:"updated_at" xorm:"not null default 'CURRENT_TIMESTAMP' TIMESTAMP"`
}
