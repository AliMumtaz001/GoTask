package authentication

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/AliMumtaz001/GoTask/api/config"
	model "github.com/AliMumtaz001/GoTask/models"
	"github.com/AliMumtaz001/GoTask/repositories"
	"github.com/AliMumtaz001/GoTask/services"
	"github.com/gin-gonic/gin"
)

// Signup godoc
// @Summary      Signup user
// @Description  Authenticate user and return JWT token
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        user  body  authentication.users  true  "User Credentials"
// @Success      200
// @Failure      401
// @Router       /signup [post]
func Signup(c *gin.Context, db *sql.DB, envConfig config.DBConfig) {
	var user model.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	authService := services.AuthService{
		UserRepo: &repositories.UserRepository{DB: db},
	}

	err := authService.RegisterUser(user)
	if err != nil {
		c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
		return
	}

	jwtWrapper := services.JwtWrap{
		SecretKey:       "esfsdfkpskodkf24234243243",
		Issued:          "admin",
		ExpirationHours: 24,
	}

	signedToken, jwtErr := jwtWrapper.GenerateToken(user.Email)
	if jwtErr != nil {
		log.Println("JWT error:", jwtErr)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"error":   false,
		"message": "Successfully signed up",
		"token":   signedToken,
	})
}
