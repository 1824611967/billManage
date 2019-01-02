package dbslite

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-oci8"
)

var Db *sql.DB

func init() {
	var err error
	Db, err = sql.Open("oci8", "wang/1824611967")

	if err != nil {
		fmt.Print(err.Error())
		log.Fatal(err)
	}

	err = Db.Ping()
	if err != nil {
		fmt.Print("未连接")
	}
}
