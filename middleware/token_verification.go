package middleware

import (
	"log"
	"net/http"
	"strings"
	myutils "study_go/myUtils"

	"github.com/gin-gonic/gin"
)

var (
	customUtil myutils.JwtUtil = myutils.NewJwtUtil()
)

type TokenVerification interface {
	TokenVerify(*gin.Context)
}

type tokenVerification struct {}

func NewTokenVerification() TokenVerification {
	return &tokenVerification{}
}

func (middleware *tokenVerification) TokenVerify(ctx *gin.Context) {
	bearerToken := ctx.Request.Header.Get("Authorization")

	log.Println(bearerToken)

	if bearerToken == "" {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"message": "No Access Token",
		})
		ctx.Abort()
		return
	}

	token := strings.Split(bearerToken, " ")[1]

	if len(ctx.Keys) == 0 {
		ctx.Keys = make(map[string]interface{})
	}

	userId, err := customUtil.ParseTokenWithSecretKey(token)
	if err != nil {
		panic(err)
	}

	ctx.Set("userId", userId)

	ctx.Next()
}