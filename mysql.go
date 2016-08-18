package test_sql


import (
	"fmt"
	"database/sql"
)

import _ "github.com/go-sql-driver/mysql"

func main() {
	db, err := sql.Open("mysql", "root:password@/amazon")

	if err != nil {
		fmt.Println( err )
	}

	defer db.Close()

	row := db.QueryRow( "select * from product limit 1")

	fmt.Println( row )
}
