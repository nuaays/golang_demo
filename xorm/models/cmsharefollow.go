package models

import (
	"time"
)

type CmShareFollow struct {
	Id        int       `json:"id" xorm:"not null pk autoincr INT(11)"`
	UserId    int       `json:"user_id" xorm:"unique(unique) index(user_share_type) INT(11)"`
	ShareId   int       `json:"share_id" xorm:"unique(unique) index(user_share_type) INT(11)"`
	Type      int       `json:"type" xorm:"unique(unique) index(user_share_type) INT(11)"`
	ThirdId   int       `json:"third_id" xorm:"unique(unique) INT(11)"`
	Time      int       `json:"time" xorm:"INT(11)"`
	Option    string    `json:"option" xorm:"CHAR(255)"`
	UpdatedAt time.Time `json:"updated_at" xorm:"not null default 'CURRENT_TIMESTAMP' TIMESTAMP"`
}
