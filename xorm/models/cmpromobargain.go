package models

import (
	"time"
)

type CmPromoBargain struct {
	UserId    int       `json:"user_id" xorm:"not null pk default 0 INT(11)"`
	Value     int       `json:"value" xorm:"INT(11)"`
	Source    string    `json:"source" xorm:"CHAR(32)"`
	Status    int       `json:"status" xorm:"TINYINT(8)"`
	Time      int       `json:"time" xorm:"INT(11)"`
	UpdatedAt time.Time `json:"updated_at" xorm:"not null default 'CURRENT_TIMESTAMP' TIMESTAMP"`
}
