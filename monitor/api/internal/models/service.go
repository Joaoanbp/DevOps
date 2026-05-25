package models

import "gorm.io/datatypes"

type ServiceType string

const (
	ServiceWeb      ServiceType = "web"
	ServiceDatabase ServiceType = "database"
	ServiceSMTP     ServiceType = "smtp"
	ServiceDNS      ServiceType = "dns"
)

type Service struct {
	Base

	Enabled bool           `gorm:"not null;default:true"`
	Type    ServiceType    `gorm:"not null" json:"type"`
	Cron    string         `gorm:"not null" json:"cron"`
	Config  datatypes.JSON `gorm:"not null" json:"config"`

	Reports []Report
	Routes  []DynamicRoute
}
