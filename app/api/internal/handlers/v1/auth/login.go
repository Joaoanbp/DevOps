package v1_auth

import (
	"api/internal/models"
	"net/http"
	"github.com/gin-gonic/gin"
)

func (this *AuthController) Post(ctx *gin.Context) {	
	email := ctx.Query("email")
	password := ctx.Query("password")
	
	result := this.Database.Create(&models.User{
		Email: email,
		Password: password,
	})

	if result.Error != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": result.Error.Error(),
		})
		return
	}
	
	ctx.JSON(http.StatusOK, gin.H{
		"mogged": true,
	})
}