package models

import "gorm.io/datatypes"

type ReportType string

const (
	ReportSuccess ReportType = "success"
	ReportWarn    ReportType = "warn"
	ReportError   ReportType = "error"
)

type Report struct {
	Base

	Enabled bool           `gorm:"not null;default:true"`
	Status  ReportType     `gorm:"not null" json:"status"`
	Metrics datatypes.JSON `gorm:"type:text" json:"metrics"`

	Service Service
}
