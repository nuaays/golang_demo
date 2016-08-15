package models

import (
	"time"
)

type CmC2Cc struct {
	Id        int       `json:"id" xorm:"not null pk autoincr INT(11)"`
	ClueId    int       `json:"clue_id" xorm:"not null unique(IDX_CL_CA) INT(11)"`
	Phone     string    `json:"phone" xorm:"not null index CHAR(24)"`
	CarId     string    `json:"car_id" xorm:"not null unique(IDX_CL_CA) CHAR(16)"`
	Status    string    `json:"status" xorm:"not null CHAR(16)"`
	Source    string    `json:"source" xorm:"not null CHAR(32)"`
	Channel   string    `json:"channel" xorm:"not null VARCHAR(128)"`
	CreatedAt time.Time `json:"created_at" xorm:"not null default 'CURRENT_TIMESTAMP' index DATETIME"`
	UpdatedAt time.Time `json:"updated_at" xorm:"default '0000-00-00 00:00:00' DATETIME"`
}
