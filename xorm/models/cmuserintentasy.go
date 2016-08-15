package models

import (
	"time"
)

type CmUserIntentAsy struct {
	Id         int       `json:"id" xorm:"not null pk autoincr INT(11)"`
	IntentData string    `json:"intent_data" xorm:"TEXT"`
	Status     int       `json:"status" xorm:"default 0 index TINYINT(1)"`
	CreatedAt  time.Time `json:"created_at" xorm:"default 'CURRENT_TIMESTAMP' DATETIME"`
	UpdatedAt  time.Time `json:"updated_at" xorm:"default 'CURRENT_TIMESTAMP' DATETIME"`
}
