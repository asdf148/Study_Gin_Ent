// connectDB
package DB

import (
	"database/sql"
	"time"

	"github.com/go-sql-driver/mysql"
)

func GetConnector() *sql.DB {
	cfg := mysql.Config{
		User:                 "Gin",
		Passwd:               "Gin",
		Net:                  "tcp",
		Addr:                 "127.0.0.1:3306",
		Collation:            "utf8mb4_general_ci",
		Loc:                  time.UTC,
		MaxAllowedPacket:     4 << 20.,
		AllowNativePasswords: true,
		CheckConnLiveness:    true,
		DBName:               "gin_ent",
	}
	connector, err := mysql.NewConnector(&cfg)
	if err != nil {
		panic(err)
	}
	db := sql.OpenDB(connector)
	return db
}
