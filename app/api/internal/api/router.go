package api

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"api/internal/database"
	"api/internal/framework"

	"api/internal/handlers/v1"
	v1_auth "api/internal/handlers/v1/auth"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

var (
	Router *gin.Engine
	start  time.Time
)

func Initialize() {
	start = time.Now()
	if err := godotenv.Load(); err != nil {
		log.Println("Aviso: Arquivo .env não encontrado, utilizando variáveis de ambiente do sistema.")
	}

	if mode := os.Getenv("GIN_MODE"); mode != "" {
		gin.SetMode(mode)
	}

	Router = gin.Default()
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{fmt.Sprintf("http://localhost:%d", 6060)}
	config.AllowMethods = []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD", "OPTIONS"}
	config.AllowHeaders = []string{"Origin", "Content-Length", "Content-Type", "Authorization"}
	Router.Use(cors.New(config))
	Router.HandleMethodNotAllowed = true
	Router.NoRoute(func(ctx *gin.Context) {
		ctx.AbortWithStatus(http.StatusNoContent)
	})

	/*
	db, err := database.New(database.ConnectionConfig{
		DriverType: database.Postgres,
		PostgresHost:     os.Getenv("DB_HOST"),
		PostgresPort:     os.Getenv("DB_PORT"),
		PostgresUser:     os.Getenv("DB_USERNAME"),
		PostgresPassword: os.Getenv("DB_PASSWORD"),
		PostgresDBName:   os.Getenv("DB_NAME"),
		PostgresSSLMode:  os.Getenv("DB_SSLMODE"),
	})
	*/

	db, err := database.New(database.ConnectionConfig{
		DriverType: database.SQLite,
		SQLitePath: "fuck.sql",
	})

	if err != nil {
		log.Panic(err)
	}

	base := &framework.BaseController{Database: db.Driver, Router: Router.Group("")}
	v1 := base.Group("/v1")

	// v1.Use(framework.RateLimit(rate.Every(time.Second), 5)).Register("/submit", submit.New(v1))
	v1.Register("/", handler_v1.New(v1))
	v1.Register("/", v1_auth.New(v1))
}
