package database

import (
	"api/internal/models"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Datasource struct {
	Driver *gorm.DB
}

func New(filepath string, config gorm.Config) (*Datasource, error) {
	db, err := gorm.Open(sqlite.Open(filepath), &config)
	if (err != nil) {
		return nil, err
	} 

	return &Datasource{
		Driver: db,
	}, nil
}

func (datasource *Datasource) Migration() {
	datasource.Driver.AutoMigrate(
		&models.Service{},
		&models.Report{},
		&models.DynamicRoute{},
	)
}