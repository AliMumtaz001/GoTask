// package authentication

// import (
// 	"encoding/json"
// 	"fmt"
// 	"log"
// 	"net/http"

// 	"github.com/AliMumtaz001/GoTask/database"
// 	"github.com/gin-gonic/gin"
// )

// func Signup(c *gin.Context) {
// 	var u users
// 	err := json.NewDecoder(c.Request.Body).Decode(&u)
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
// 		return
// 	}
// 	db := database.Connect()

// 	query := `INSERT INTO Employeedata (email, password) VALUES ($1, $2) RETURNING email`
// 	err = db.QueryRow(query, u.Email, u.Password).Scan(&u.Email, &u.Password)
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": "User already existed with this email"})
// 		fmt.Println(err)
// 		return
// 	}

// 	c.JSON(200, gin.H{
// 		"message": "User created successfully",
// 		"email":   u.Email,
// 	})

// 	//got error:= no new variables on left side of :=compiler
// 	//so instead of err i use e
// 	e := c.BindJSON(&u)
// 	if e != nil {
// 		log.Println(e)
// 		return
// 	}
// 	_, ok := userss[u.Email]
// 	if ok {
// 		c.JSON(http.StatusConflict, gin.H{"message": "you already created an account"})
// 		return
// 	}

// 	userss[u.Email] = u

// 	jwtWrapper := JwtWrap{
// 		SecretKey:       "esfsdfkpskodkf24234243243",
// 		Issued:          "admin",
// 		ExpirationHours: 12,
// 	}
// 	signedToken, jwtErr := jwtWrapper.GenerateToken(u.Email)
// 	if jwtErr != nil {
// 		log.Println(jwtErr)
// 		return
// 	}
// 	c.JSON(http.StatusOK, gin.H{"error": false, "message": "Successfully Signed Up", "token": signedToken})
// }

package authentication

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/AliMumtaz001/GoTask/database"
	"github.com/gin-gonic/gin"
)

func Signup(c *gin.Context) {
	var u users
	err := json.NewDecoder(c.Request.Body).Decode(&u)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	db := database.Connect()

	query := `INSERT INTO employeedata (email, password) VALUES ($1, $2) RETURNING email`
	err = db.QueryRow(query, u.Email, u.Password).Scan(&u.Email)
	if err != nil {
		if strings.Contains(err.Error(), "duplicate key value violates unique constraint") {
			c.JSON(http.StatusConflict, gin.H{"error": "User already exists with this email"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user", "details": err.Error()})
		}
		fmt.Println("Database error:", err)
		return
	}

	// userss[u.Email] = u

	jwtWrapper := JwtWrap{
		SecretKey:       "esfsdfkpskodkf24234243243",
		Issued:          "admin",
		ExpirationHours: 12,
	}
	signedToken, jwtErr := jwtWrapper.GenerateToken(u.Email)
	if jwtErr != nil {
		log.Println("JWT error:", jwtErr)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"error":   false,
		"message": "Successfully Signed Up",
		"email":   u.Email,
		"token":   signedToken,
	})
}
