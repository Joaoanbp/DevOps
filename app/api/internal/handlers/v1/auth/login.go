package v1_auth

import (
	"api/internal/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (this *AuthController) Post(ctx *gin.Context) {
	ctx.SetRouter("/login")
	
	email := ctx.Query("email")
	password := ctx.Query("password")
	
	user := this.Database.Create(&models.User{
		Email: email,
		Password: password,
	})
	ctx.JSON(http.StatusOK, gin.H{
		"mogged": true,
	})
}