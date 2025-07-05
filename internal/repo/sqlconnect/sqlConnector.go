package sqlconnect

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func SqlConnector() *sql.DB {
	fmt.Println("connection to DB")
	db,err:=sql.Open("mysql","root:root@tcp(127.0.0.1:3306)/school")
	if err!=nil{
		panic(err)
	}
 return db
}