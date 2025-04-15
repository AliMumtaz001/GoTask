package main

import (
	"errors"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

type JwtWrap struct {
	SecretKey       string	//secret key to sign token .
	Issued          string	
	ExpirationHours int64	//jhow much time token valid
}

type JwtClaim struct {
	Email          string
	StandardClaims jwt.StandardClaims
}


func (j *JwtClaim) Valid() error {

	if err := j.StandardClaims.Valid(); err != nil {
		return err
	}
	return nil
}
//a tokenn is generated from the email.the token is signed using a secret key the token also has an expiry time set.
func (j *JwtWrap) GenerateToken(email string) (string, error) {
	claims := &JwtClaim{
		Email: email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Hour * time.Duration(j.ExpirationHours)).Unix(),
			Issuer:    j.Issued,
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte(j.SecretKey))
	if err != nil {
		return "", err
	}
	return signedToken, nil
}

func (j *JwtWrap) ValidateToken(signedToken string) (*JwtClaim, error) {
	token, err := jwt.ParseWithClaims(
		signedToken,
		&JwtClaim{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(j.SecretKey), nil
		},
	)
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*JwtClaim)
	if !ok {
		return nil, errors.New("couldn't parse claims")
	}

	if claims.StandardClaims.ExpiresAt < time.Now().Unix() {
		return nil, errors.New("JWT is expired")
	}
	return claims, nil
}

func Auths() gin.HandlerFunc {
	return func(c *gin.Context) {
		clintToken := c.Request.Header.Get("Authorization")
		if clintToken == "" {
			c.JSON(403, "no Authorization header provided")
			c.Abort()
			return
		}

		extractedToken := strings.Split(clintToken, "Bearer")

		if len(extractedToken) == 2 {
			clintToken = strings.TrimSpace(extractedToken[1])
		} else {
			c.JSON(400, "Incorrect Format of Authorization Token")
			c.Abort()
			return
		}

		jwtWrapper := JwtWrap{
			SecretKey: "esfsdfkpskodkf24234243243",
			Issued:    "admin",
		}
		claims, err := jwtWrapper.ValidateToken(clintToken)
		if err != nil {
			c.JSON(401, err.Error())
			c.Abort()
			return
		}
		c.Set("email", claims.Email)
		c.Next()
	}
}

type users struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

var userss map[string]users

func signup(c *gin.Context) {
	var u users
	err := c.BindJSON(&u)
	if err != nil {
		log.Println(err)
		return
	}
	_, ok := userss[u.Email]
	if ok {
		c.JSON(http.StatusConflict, gin.H{"message": "you already created an account"})
		return
	}

	userss[u.Email] = u

	jwtWrapper := JwtWrap{
		SecretKey:       "esfsdfkpskodkf24234243243",
		Issued:          "admin",
		ExpirationHours: 12,
	}
	signedToken, jwtErr := jwtWrapper.GenerateToken(u.Email)
	if jwtErr != nil {
		log.Println(jwtErr)
		return
	}
	c.JSON(http.StatusOK, gin.H{"error": false, "message": "Successfully Signed Up", "token": signedToken})
}

func login(c *gin.Context) {
	var u users
	err := c.BindJSON(&u)
	if err != nil {
		log.Println(err)
		return
	}
	val, ok := userss[u.Email]
	if !ok {
		c.JSON(http.StatusConflict, gin.H{"error": true, "message": "Please Sign Up"})
		return
	}
	jwtWrapper := JwtWrap{
		SecretKey:       "esfsdfkpskodkf24234243243",
		Issued:          "admin",
		ExpirationHours: 48,
	}
	signedToken, jwtErr := jwtWrapper.GenerateToken(u.Email)
	if jwtErr != nil {
		log.Println(jwtErr)
		return
	}
	c.JSON(http.StatusOK, gin.H{"error": false, "message": "successfully logged in", "data": val, "token": signedToken})
}

func getData(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"data": userss})
}

func main() {
	userss = make(map[string]users)

	r := gin.Default()

	r.POST("/login", login)
	r.POST("/signup", signup)
	r.Use(Auths())
	r.GET("/data", getData)

	r.Run(":9090")
}
