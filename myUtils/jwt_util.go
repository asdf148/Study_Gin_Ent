package myutils

import (
	"errors"
	dto "study_go/dto/jwt"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type JwtUtil interface {
	CreateAccessToken(int) (string, error)
	ParseTokenWithSecretKey(string) (int, error)
}

type jwtUtil struct {
}

func NewJwtUtil() JwtUtil {
	return &jwtUtil{}
}

func (c *jwtUtil) CreateAccessToken(userId int) (string, error) {

	claims := &dto.JwtClaim{
		UserId: userId,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Hour * 24).Unix(),
			Issuer:    "me",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedToken, err := token.SignedString([]byte("testClaims"))
	if err != nil {
		return "error", err
	}

	return signedToken, nil
}

func (c *jwtUtil) ParseTokenWithSecretKey(signedToken string) (int, error) {
	token, err := jwt.ParseWithClaims(
		signedToken,
		&dto.JwtClaim{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte("testClaims"), nil
		},
	)
	if err != nil {
		return 0, err
	}

	claims, ok := token.Claims.(*dto.JwtClaim)
	if !ok {
		err = errors.New("couldn't parse claims")
		return 0, err
	}

	if claims.ExpiresAt < time.Now().Local().Unix() {
		err = errors.New("JWT is expired")
		return 0, err
	}

	return claims.UserId, nil
}