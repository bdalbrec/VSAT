package models

import (
	"database/sql"
	_ "github.com/denisenkom/go-mssqldb"
	"fmt"
	"log"
)

var db *sql.DB

func InitDB(dataSourceName string) {
    var err error
    db, err = sql.Open("mssql", dataSourceName)
    if err != nil {
        log.Panic(err)
    }

	fmt.Println("Attempting to ping database.")
    if err = db.Ping(); err != nil {
        log.Panic(err)
	}
	fmt.Println("Database ping successful.")
}
	

