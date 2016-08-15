package models

import (
	"time"
)

type CmFeedbackUserRule struct {
	Id        int       `json:"id" xorm:"not null pk autoincr INT(11)"`
	Keywords  string    `json:"keywords" xorm:"VARCHAR(1000)"`
	Receivers string    `json:"receivers" xorm:"VARCHAR(1000)"`
	UpdatedAt time.Time `json:"updated_at" xorm:"not null default 'CURRENT_TIMESTAMP' TIMESTAMP"`
}
