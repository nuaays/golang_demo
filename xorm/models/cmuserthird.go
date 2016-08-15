package models

import (
	"time"
)

type CmUserThird struct {
	Id        int       `json:"id" xorm:"not null pk autoincr INT(11)"`
	UserId    int       `json:"user_id" xorm:"index INT(11)"`
	Type      int       `json:"type" xorm:"not null unique(open_id) TINYINT(3)"`
	OpenId    string    `json:"open_id" xorm:"not null unique(open_id) CHAR(45)"`
	UnionId   string    `json:"union_id" xorm:"not null CHAR(45)"`
	Nickname  string    `json:"nickname" xorm:"CHAR(32)"`
	HeadUrl   string    `json:"head_url" xorm:"VARCHAR(512)"`
	Extends   string    `json:"extends" xorm:"not null VARCHAR(1024)"`
	Time      int       `json:"time" xorm:"INT(11)"`
	UpdatedAt time.Time `json:"updated_at" xorm:"not null default 'CURRENT_TIMESTAMP' TIMESTAMP"`
}
