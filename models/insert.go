package models

import "log"
import "time"

func Insert(loc string, eqp string, date string, tech string) {

	timestamp := time.Now()

	res, err := db.Exec("INSERT INTO audits VALUES($1, $2, $3, $4, $5)", timestamp, loc, eqp, date, tech)
	if err != nil{
		log.Fatal(err)
	}

	rowCnt, err := res.RowsAffected()
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Rows affected = %d\n", rowCnt)
}