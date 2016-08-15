package models

import (
	"time"
)

type DtsIncrementTrx struct {
	JobId      string    `json:"job_id" xorm:"not null pk CHAR(32)"`
	Partition  int       `json:"partition" xorm:"not null pk INT(11)"`
	Checkpoint string    `json:"checkpoint" xorm:"VARCHAR(256)"`
	UpdatedAt  time.Time `json:"updated_at" xorm:"not null default 'CURRENT_TIMESTAMP' TIMESTAMP"`
}
