package v1_auth
//cadastro né
import (
	"api/internal/models"
	"net/http"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"encoding/json"
)

func SignupHandler(c *gin.Context) {
	var creds models.Credentials

	if err := c.ShouldBindJSON(&creds); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Payload inválido"})
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(creds.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro interno do servidor"})
		return
	}
	_ = hashedPassword

	c.JSON(http.StatusCreated, gin.H{"message": "Usuário criado com sucesso"})
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