package authentication

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/AliMumtaz001/GoTask/database"
	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	var u users
	var dbemail, dbpassword string
	db := database.Connect()

	err := json.NewDecoder(c.Request.Body).Decode(&u)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON token"})
		return
	}

	query := `SELECT email, password FROM employee WHERE email = $1`
	err = db.QueryRow(query, u.Email).Scan(&dbemail, &dbpassword)

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not found"})
		return
	}

	if dbpassword != u.Password {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	val, ok := u.Email, true
	if !ok {
		c.JSON(http.StatusConflict, gin.H{"error": true, "message": "Please Sign Up"})
		return
	}

	// generate JWT token
	jwtWrapper := JwtWrap{
		SecretKey:       "esfsdfkpskodkf24234243243",
		Issued:          "admin",
		ExpirationHours: 48,
	}
	signedToken, jwtErr := jwtWrapper.GenerateToken(u.Email)
	if jwtErr != nil {
		log.Println("JWT error:", jwtErr)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	// return success response
	c.JSON(http.StatusOK, gin.H{
		"error":   false,
		"message": "Successfully logged in",
		"data":    val,
		"token":   signedToken,
	})
}
