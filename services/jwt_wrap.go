package services

import (
    "time"

    "github.com/golang-jwt/jwt"
)

type JwtWrap struct {
    SecretKey       string
    Issued          string
    ExpirationHours int
}

type JwtClaims struct {
    Email string `json:"email"`
    jwt.StandardClaims
}

func (j *JwtWrap) GenerateToken(email string) (string, error) {
    claims := &JwtClaims{
        Email: email,
        StandardClaims: jwt.StandardClaims{
            ExpiresAt: time.Now().Add(time.Hour * time.Duration(j.ExpirationHours)).Unix(),
            Issuer:    j.Issued,
        },
    }

    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    return token.SignedString([]byte(j.SecretKey))
}