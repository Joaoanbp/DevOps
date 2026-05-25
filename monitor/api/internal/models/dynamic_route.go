package models

type RouteType string

const (
	RouteTypeLog = "log"
)

type DynamicRoute struct {
	Base
	
	Type RouteType `gorm:"not null" json:"type"`
}
