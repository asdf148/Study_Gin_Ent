// main
package main

import (
	"net/http"
	DB "study_go/DB"

	"github.com/gin-gonic/gin"
)

func setupRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})
	return r
}

func main() {
	db := DB.GetConnector()
	err := db.Ping()
	if err != nil {
		panic(err)
	}

	r := setupRouter()
	r.Run(":8080")
}
