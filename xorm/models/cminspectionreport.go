package models

import (
	"time"
)

type CmInspectionReport struct {
	Locked    int       `json:"locked" xorm:"not null pk default 0 INT(11)"`
	UpdatedAt time.Time `json:"updated_at" xorm:"not null default 'CURRENT_TIMESTAMP' TIMESTAMP"`
}
