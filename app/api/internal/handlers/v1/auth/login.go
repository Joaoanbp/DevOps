package v1_auth

import (
	"api/internal/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"github.com/golang-jwt/jwt/v5"
)


func LoginHandler(c *gin.Context) {
	var creds models.Credentials
	if err := c.ShouldBindJSON(&creds); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Payload inválido"})
		return
	}

	expectedHash, _ := bcrypt.GenerateFromPassword([]byte("password123"), bcrypt.DefaultCost)
	
	err := bcrypt.CompareHashAndPassword(expectedHash, []byte(creds.Password))
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Credenciais inválidas"})
		return
	}

	expirationTime := time.Now().Add(24 * time.Hour)
	claims := &jwt.RegisteredClaims{
		Subject:   creds.Email,
		ExpiresAt: jwt.NewNumericDate(expirationTime),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao gerar token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": tokenString})
}


//func (this *AuthController) Post(ctx *gin.Context) {	
//	email := ctx.Query("email")
//	password := ctx.Query("password")
//	
//	result := this.Database.Create(&models.User{
//		Email: email,
//		Password: password,
//	})
//
//	if result.Error != nil {
//		ctx.JSON(http.StatusInternalServerError, gin.H{
//			"error": result.Error.Error(),
//		})
//		return
//	}
//	
//	ctx.JSON(http.StatusOK, gin.H{
//		"mogged": true,
//	})
//}