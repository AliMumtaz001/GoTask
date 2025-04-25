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

// Login godoc
// @Summary      Login user
// @Description  Authenticate user and return JWT token
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        user  body  authentication.users  true  "User Credentials"
// @Success      200
// @Failure      401
// @Router       /login [post]
func Login(c *gin.Context, db *sql.DB, envConfig config.DBConfig) {
	var credentials model.UserCredentials
	if err := c.ShouldBindJSON(&credentials); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	authService := services.AuthService{
		UserRepo: &repositories.UserRepository{DB: db},
	}

	user, err := authService.AuthenticateUser(credentials)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	// Generate JWT toke
	jwtWrapper := services.JwtWrap{
		SecretKey:       "esfsdfkpskodkf24234243243",
		Issued:          "admin",
		ExpirationHours: 48,
	}

	signedToken, jwtErr := jwtWrapper.GenerateToken(user.Email)
	if jwtErr != nil {
		log.Println("JWT error:", jwtErr)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"error":   false,
		"message": "Successfully logged in",
		"token":   signedToken,
	})
}
