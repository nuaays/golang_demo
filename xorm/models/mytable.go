package models

import (
	"time"
)

type Mytable struct {
	Id      int       `json:"id" xorm:"not null pk autoincr INT(11)"`
	Created time.Time `json:"created" xorm:"not null default '0000-00-00 00:00:00' TIMESTAMP"`
	Updated time.Time `json:"updated" xorm:"not null default 'CURRENT_TIMESTAMP' TIMESTAMP"`
	Myfield string    `json:"myfield" xorm:"VARCHAR(255)"`
}
