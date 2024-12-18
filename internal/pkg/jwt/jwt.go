package jwt

import (
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/spf13/cast"
)

func JWTGenerate(data map[string]string, expireSeconds int64, signKey string) (string, error) {
	claims := jwt.MapClaims{}
	for k, v := range data {
		claims[k] = v
	}

	claims["iat"] = time.Now().Unix()
	if expireSeconds > 0 {
		claims["exp"] = time.Now().Add(time.Second * time.Duration(expireSeconds)).Unix()
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(signKey))
}

func JWTExtract(tokenStr string, signKey string) (map[string]string, error) {
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		return []byte(signKey), nil
	})
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !(ok && token.Valid) {
		err = fmt.Errorf("INVALID TOKEN")
		return nil, err
	}

	res := map[string]string{}
	for k, v := range claims {
		res[k] = cast.ToString(v)
	}

	return res, nil
}
