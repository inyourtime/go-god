package coreplugins

import (
	"github.com/golang-jwt/jwt/v5"
)

func Token(claims jwt.Claims, secret string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString([]byte(secret))
	return t, err
}

func DecodedToken(token string, secret string) (jwt.MapClaims, error) {
	c := jwt.MapClaims{}
	_, err := jwt.ParseWithClaims(token, &c, func(t *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})
	return c, err
}
