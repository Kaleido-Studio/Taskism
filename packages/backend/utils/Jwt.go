package utils

import (
	"encoding/base64"
	"log"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type Claims struct {
	UserId    string `json:"jti"`
	Authority string `json:"authorities"`
	jwt.RegisteredClaims
}

func ParseToken(token string) (*Claims, error) {
	secret, _ := base64.RawStdEncoding.DecodeString("JWTSecret")
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return secret, nil
	})
	if err != nil {
		log.Printf("fuck jwt error")
		return nil, err
	}

	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}

	return nil, err
}

func GenerateToken(userid string) (string, error) {
	secret, _ := base64.RawStdEncoding.DecodeString("JWTSecret")
	nowTime := time.Now()
	expireTime := nowTime.Add(300 * 30000 * time.Second)
	issuer := "maiquer"
	claims := Claims{
		UserId: userid,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: &jwt.NumericDate{Time: expireTime},
			Issuer:    issuer,
		},
	}

	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString(secret)
	return "metric-" + token, err
}
