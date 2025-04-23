package authentication

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/AliMumtaz001/GoTask/database"
	"github.com/gin-gonic/gin"
)

func Signup(c *gin.Context) {
	var u users
	err := c.BindJSON(&u)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid input",
			"email": u.Email,
		})
		return
	}

	// Validate email
	if u.Email == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Email cannot be empty",
			"email": u.Email,
		})
		return
	}

	db := database.Connect()
	defer db.Close()

	// check if email already exists
	var existingEmail string
	checkQuery := `SELECT email FROM employee WHERE email = $1`
	err = db.QueryRow(checkQuery, u.Email).Scan(&existingEmail)
	if err == nil {
		// if email already exists
		c.JSON(http.StatusConflict, gin.H{
			"error": "User already exists with this email",
			"email": u.Email,
		})
		return
	} else if err.Error() != "sql: no rows in result set" {
		// Handle other database errors
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Database error",
			"details": err.Error(),
			"email":   u.Email,
		})
		fmt.Println("Database error:", err)
		return
	}

	// iif email doesn't exist, proceed with insertion
	query := `INSERT INTO employee (email, password) VALUES ($1, $2) RETURNING email`
	err = db.QueryRow(query, u.Email, u.Password).Scan(&u.Email)
	if err != nil {
		if strings.Contains(err.Error(), "duplicate key value violates unique constraint") {
			c.JSON(http.StatusConflict, gin.H{
				"error": "User already exists with this email",
				"email": u.Email,
			})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error":   "Failed to create user",
				"details": err.Error(),
				"email":   u.Email,
			})
		}
		fmt.Println("Database error:", err)
		return
	}

	// generate JWT token
	jwtWrapper := JwtWrap{
		SecretKey:       "esfsdfkpskodkf24234243243",
		Issued:          "admin",
		ExpirationHours: 24,
	}
	signedToken, jwtErr := jwtWrapper.GenerateToken(u.Email)
	if jwtErr != nil {
		log.Println("JWT error:", jwtErr)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to generate token",
			"email": u.Email,
		})
		return
	}

	// return success response
	c.JSON(http.StatusOK, gin.H{
		"error":   false,
		"message": "Successfully Signed Up",
		"email":   u.Email,
		"token":   signedToken,
	})
}