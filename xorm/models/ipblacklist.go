package models

import (
	"time"
)

type IpBlacklist struct {
	Id             int       `json:"id" xorm:"not null pk autoincr INT(11)"`
	Ip             string    `json:"ip" xorm:"unique VARCHAR(45)"`
	Timestamp      string    `json:"timestamp" xorm:"VARCHAR(45)"`
	IpBlacklistcol string    `json:"ip_blacklistcol" xorm:"VARCHAR(45)"`
	CreatedAt      time.Time `json:"created_at" xorm:"default 'CURRENT_TIMESTAMP' DATETIME"`
	UpdatedAt      time.Time `json:"updated_at" xorm:"default 'CURRENT_TIMESTAMP' DATETIME"`
}
