package models

type User struct {
	Base

	Email string           `gorm:"not null;unique" json:"email"`
	Password string         `gorm:"not null" json:"password"`
}
