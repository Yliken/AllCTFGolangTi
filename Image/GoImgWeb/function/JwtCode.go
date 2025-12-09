package function

import (
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

func JwtEncode(username, StudentNum string) (string, error) {
	claims := jwt.MapClaims{
		"username": username,
		"stuNum":   StudentNum,
		"ImgUrl":   "", // 题目会访问这里的地址
		"iat":      time.Now(),
		"exp":      time.Now().Add(time.Minute * 30).Unix(),
	}
	signedString, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte("Can not lose RedBean"))
	return signedString, err
}
func JwtDecode(token string) (*jwt.MapClaims, error) {
	parse, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte("Can not lose RedBean"), nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := parse.Claims.(jwt.MapClaims); ok && parse.Valid {
		return &claims, nil
	}
	return nil, fmt.Errorf("Token invalid")
}
