package models

import (
	"time"
)

type CmFeedbackMsg struct {
	Id        int       `json:"id" xorm:"not null pk autoincr INT(11)"`
	Uuid      string    `json:"uuid" xorm:"index VARCHAR(60)"`
	Content   string    `json:"content" xorm:"VARCHAR(500)"`
	Sender    int       `json:"sender" xorm:"default 0 INT(11)"`
	Timestamp string    `json:"timestamp" xorm:"VARCHAR(45)"`
	CarId     string    `json:"car_id" xorm:"CHAR(16)"`
	Title     string    `json:"title" xorm:"VARCHAR(45)"`
	CreatedAt time.Time `json:"created_at" xorm:"default 'CURRENT_TIMESTAMP' DATETIME"`
	UpdatedAt time.Time `json:"updated_at" xorm:"default 'CURRENT_TIMESTAMP' DATETIME"`
}
