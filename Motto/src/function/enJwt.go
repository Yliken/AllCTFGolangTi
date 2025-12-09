package function

import (
	"SummerVactionSQL/config"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

func EnJwt(username, nickname string) (string, error) {
	claims := jwt.MapClaims{
		"username": username,
		"nickname": nickname,
		"iat":      time.Now(),
		"exp":      time.Now().Add(time.Minute * 30).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(config.JwtKey))
}
