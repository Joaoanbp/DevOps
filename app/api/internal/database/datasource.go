package database

import (
	"api/internal/models"
	"fmt"
	"log"
	"sort"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Driver string

const (
	SQLite   Driver = "sqlite"
	Postgres Driver = "postgres"
)

type ConnectionConfig struct {
	DriverType       Driver
	SQLitePath       string
	PostgresHost     string
	PostgresPort     string
	PostgresUser     string
	PostgresPassword string
	PostgresDBName   string
	PostgresSSLMode  string
}
type Datasource struct {
	Driver *gorm.DB
}

func New(config ConnectionConfig, gormConfig gorm.Config) (*Datasource, error) {
	var database *gorm.DB
	var connectionError error

	switch config.DriverType {
	case "sqlite":
		database, connectionError = gorm.Open(sqlite.Open(config.SQLitePath), &gormConfig)
	case "postgres":
		if config.PostgresSSLMode == "" {
			config.PostgresSSLMode = "disable"
		}

		connectionString := fmt.Sprintf(
			"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
			config.PostgresHost,
			config.PostgresPort,
			config.PostgresUser,
			config.PostgresPassword,
			config.PostgresDBName,
			config.PostgresSSLMode,
		)

		database, connectionError = gorm.Open(postgres.Open(connectionString), &gormConfig)
	default:
		return nil, fmt.Errorf("tipo de driver não suportado: %s", config.DriverType)
	}

	if connectionError != nil {
		return nil, connectionError
	}

	return &Datasource{
		Driver: database,
	}, nil
}

func (datasource *Datasource) AutoMigrate() error {
	log.Println("[Database] Executando AutoMigrate declarativo")

	return datasource.Driver.AutoMigrate(
		&models.User{},
	)
}