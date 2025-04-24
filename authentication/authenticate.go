package authentication

import (
	"net/http"
	"os"
	"strings"

	"github.com/AliMumtaz001/GoTask/services"
	"github.com/gin-gonic/gin"
)

func Auths() gin.HandlerFunc {
	return func(c *gin.Context) {
		clientToken := c.Request.Header.Get("Authorization")
		if clientToken == "" {
			c.JSON(http.StatusForbidden, gin.H{"error": "No Authorization header provided"})
			c.Abort()
			return
		}

		if !strings.HasPrefix(clientToken, "Bearer ") {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Authorization header must start with 'Bearer '"})
			c.Abort()
			return
		}

		token := strings.TrimSpace(strings.TrimPrefix(clientToken, "Bearer "))

		secretKey := os.Getenv("JWT_SECRET")
		if secretKey == "" {
			secretKey = "esfsdfkpskodkf24234243243" 
		}

		jwtService := services.JwtService{
			SecretKey: secretKey,
			Issuer:    "admin",
		}

		claims, err := jwtService.ValidateToken(token)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			c.Abort()
			return
		}

		c.Set("email", claims.Email)
		c.Next()
	}
}
