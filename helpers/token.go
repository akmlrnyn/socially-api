package helpers

import (
	"errors"
	"gin_social/entities"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var mySigningKey = []byte("mysecretkey")

type JWTClaims struct{
	id int `json:"id"`
	jwt.RegisteredClaims
}

func GenerateToken(user *entities.User) (string, error) {
	claims := JWTClaims{
		user.ID,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(60 * time.Minute)),
			IssuedAt: jwt.NewNumericDate(time.Now().Add(60 * time.Minute)),
			NotBefore: jwt.NewNumericDate(time.Now().Add(60 * time.Minute)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString(mySigningKey)

	return ss, err
}

func ValidateToken(tokenString string) (*int, error) {
	token, err := jwt.ParseWithClaims(tokenString, JWTClaims{}, func (token *jwt.Token)(interface{}, error) {
		return mySigningKey, nil
	})


	if err != nil{
		if err == jwt.ErrSignatureInvalid{
			return nil, errors.New("Invalid token signature")
		}

		return nil, errors.New("your token was expired")
	}

	claims, ok := token.Claims.(*JWTClaims)

	if !ok || !token.Valid{
		return nil, errors.New("your token was expired")
	}

	return &claims.id, nil
}
