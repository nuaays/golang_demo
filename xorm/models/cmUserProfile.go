package models

import (
	"time"
)

type CmUserProfile struct {
	Id                 int       `json:"id" xorm:"not null pk autoincr INT(10)"`
	Phone              string    `json:"phone" xorm:"not null unique VARCHAR(45)"`
	Name               string    `json:"name" xorm:"VARCHAR(45)"`
	NickName           string    `json:"nick_name" xorm:"VARCHAR(45)"`
	Gender             string    `json:"gender" xorm:"VARCHAR(45)"`
	Age                int       `json:"age" xorm:"INT(11)"`
	Region             string    `json:"region" xorm:"VARCHAR(45)"`
	Webchat            string    `json:"webchat" xorm:"VARCHAR(45)"`
	Qq                 string    `json:"qq" xorm:"VARCHAR(45)"`
	Email              string    `json:"email" xorm:"VARCHAR(45)"`
	Budget             string    `json:"budget" xorm:"not null VARCHAR(10)"`
	TimeLimit          string    `json:"time_limit" xorm:"not null VARCHAR(45)"`
	CarQuota           string    `json:"car_quota" xorm:"not null VARCHAR(45)"`
	BrandPreference    string    `json:"brand_preference" xorm:"VARCHAR(45)"`
	SeriesPreference   string    `json:"series_preference" xorm:"VARCHAR(45)"`
	AgePreference      string    `json:"age_preference" xorm:"VARCHAR(45)"`
	Occupation         string    `json:"occupation" xorm:"VARCHAR(45)"`
	Income             string    `json:"income" xorm:"VARCHAR(45)"`
	Insurance          string    `json:"insurance" xorm:"VARCHAR(45)"`
	House              string    `json:"house" xorm:"VARCHAR(45)"`
	PreferenceDescribe string    `json:"preference_describe" xorm:"VARCHAR(512)"`
	PortraitSource     string    `json:"portrait_source" xorm:"not null VARCHAR(45)"`
	CreatedAt          time.Time `json:"created_at" xorm:"default 'CURRENT_TIMESTAMP' TIMESTAMP"`
	UpdatedAt          time.Time `json:"updated_at" xorm:"default 'CURRENT_TIMESTAMP' TIMESTAMP"`
	ExtraUpdatedAt     time.Time `json:"extra_updated_at" xorm:"TIMESTAMP"`
}
