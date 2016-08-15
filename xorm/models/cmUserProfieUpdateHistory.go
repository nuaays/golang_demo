package models

import (
	"time"
)

type CmUserProfieUpdateHistory struct {
	Id             int       `json:"id" xorm:"not null pk autoincr unique INT(11)"`
	Phone          string    `json:"phone" xorm:"not null unique VARCHAR(45)"`
	UserId         int       `json:"user_id" xorm:"not null unique INT(11)"`
	Occupation     string    `json:"occupation" xorm:"VARCHAR(45)"`
	House          string    `json:"house" xorm:"VARCHAR(45)"`
	PortraitSource string    `json:"portrait_source" xorm:"not null VARCHAR(45)"`
	UpdatedAt      time.Time `json:"updated_at" xorm:"default 'CURRENT_TIMESTAMP' TIMESTAMP"`
	CreatedAt      time.Time `json:"created_at" xorm:"default 'CURRENT_TIMESTAMP' TIMESTAMP"`
}
