package models

import "gorm.io/datatypes"

type Service struct {
	Base

	Email string           `gorm:"not null;unique" json:"email"`
	Password string         `gorm:"not null" json:"password"`
}
