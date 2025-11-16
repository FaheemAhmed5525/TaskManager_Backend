package utils

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// MARK: - ToDo
var jwtSecretKey = []byte("your_secret_key")

type JWTClaim struct {
	UserId int    `json:"user_id"`
	Email  string `json:"email"`
	jwt.RegisteredClaims
}

func GenerateJwtToken(userId int, email string) (string, error) {
	expirationTime := time.Now().Add(24 * time.Hour)

	jwtClaim := &JWTClaim{
		UserId: userId,
		Email:  email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Subject:   email,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwtClaim)
	return token.SignedString(jwtSecretKey)
}

func ValidateJWTToken(tokenStr string) (*JWTClaim, error) {
	jwtClaim := &JWTClaim{}

	token, error := jwt.ParseWithClaims(tokenStr, jwtClaim, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return jwtSecretKey, nil
	})

	if error != nil {
		return nil, error
	}

	if !token.Valid {
		return nil, fmt.Errorf("invalid token")
	}

	return jwtClaim, nil
}
