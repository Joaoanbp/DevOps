package main

import (
	"api/internal/database"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)
func main() {
  r := gin.Default()

	db, err := database.New("fuck.sql", gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.Migration()

	r.GET("/ping", func(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{
      "message": "pong",
    })
  })

  if err := r.Run("0.0.0.0:4000"); err != nil {
    log.Fatalf("failed to run server: %v", err)
  }
}