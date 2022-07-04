package jwt

import "github.com/dgrijalva/jwt-go"

type JwtClaim struct {
	UserId int
	jwt.StandardClaims
}