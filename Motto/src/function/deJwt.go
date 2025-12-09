package function

import (
	"SummerVactionSQL/config"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
)

func DeJwt(token string) (*jwt.MapClaims, error) {
	parse, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(config.JwtKey), nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := parse.Claims.(jwt.MapClaims); ok && parse.Valid {
		return &claims, nil
	}
	return nil, fmt.Errorf("Token invalid")
}
