package models

import (
	"time"
)

type CmUserSubsys struct {
	Id         int       `json:"id" xorm:"not null pk autoincr INT(11)"`
	Title      string    `json:"title" xorm:"not null VARCHAR(32)"`
	Token      string    `json:"token" xorm:"not null VARCHAR(32)"`
	Status     int       `json:"status" xorm:"not null default 0 TINYINT(1)"`
	Createtime int       `json:"createtime" xorm:"not null INT(10)"`
	Updatetime int       `json:"updatetime" xorm:"not null INT(10)"`
	UpdatedAt  time.Time `json:"updated_at" xorm:"not null default 'CURRENT_TIMESTAMP' TIMESTAMP"`
}
