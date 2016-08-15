package models

import (
	"time"
)

type Migrations struct {
	Migration string    `json:"migration" xorm:"not null VARCHAR(255)"`
	Batch     int       `json:"batch" xorm:"not null INT(11)"`
	CreatedAt time.Time `json:"created_at" xorm:"default 'CURRENT_TIMESTAMP' DATETIME"`
	UpdatedAt time.Time `json:"updated_at" xorm:"not null default 'CURRENT_TIMESTAMP' TIMESTAMP"`
}
