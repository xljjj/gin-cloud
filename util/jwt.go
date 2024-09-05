package util

import (
	"github.com/golang-jwt/jwt/v5"
	"time"
)

var JWTKey []byte

type Claims struct {
	UserName string `json:"user_name"`
	jwt.RegisteredClaims
}

// GenerateToken 生成JWT
func GenerateToken(userName string) (string, error) {
	expirationTime := time.Now().Add(24 * time.Hour) // 一天后过期

	claims := &Claims{
		UserName: userName,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(JWTKey)
}

func ParseToken(tokenString string) (*Claims, error) {
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return JWTKey, nil
	})

	if err != nil || !token.Valid {
		return nil, err
	}

	return claims, nil
}
