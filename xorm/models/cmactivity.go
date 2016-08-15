package models

import (
	"time"
)

type CmActivity struct {
	Id         int       `json:"id" xorm:"not null pk autoincr INT(11)"`
	Phone      string    `json:"phone" xorm:"index(activity_id) VARCHAR(45)"`
	ActivityId string    `json:"activity_id" xorm:"index(activity_id) VARCHAR(45)"`
	Datetime   time.Time `json:"datetime" xorm:"default 'CURRENT_TIMESTAMP' DATETIME"`
	CreatedAt  time.Time `json:"created_at" xorm:"default 'CURRENT_TIMESTAMP' DATETIME"`
	UpdatedAt  time.Time `json:"updated_at" xorm:"default 'CURRENT_TIMESTAMP' index DATETIME"`
}
