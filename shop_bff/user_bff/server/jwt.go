package server

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/pkg/errors"
)

// @secretKey: JWT 加解密密钥
// @iat: 时间戳
// @seconds: 过期时间，单位秒
// @payload: 数据载体

func SetJwtToken(secretKey string, iat, seconds int64, payload string) (string, error) {
	claims := make(jwt.MapClaims)
	claims["exp"] = iat + seconds
	claims["iat"] = iat
	claims["payload"] = payload
	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = claims
	return token.SignedString([]byte(secretKey))
}

func ParseToken(secretKey string, tokenString string) (string, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(secretKey), nil
	})

	if err != nil {
		return "", errors.New("无效的载荷")
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		return claims["payload"].(string), nil
	}

	return "", nil
}
