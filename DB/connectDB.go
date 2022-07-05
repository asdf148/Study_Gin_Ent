// connectDB
package DB

import (
	"log"
	"study_go/ent"

	_ "github.com/go-sql-driver/mysql"
)

func GetConnector() *ent.Client {
	client, err := ent.Open("mysql", "Gin:Gin@tcp(localhost:3306)/gin_ent?parseTime=True")
	if err != nil {
		log.Fatal(err)
	}
	return client
}
