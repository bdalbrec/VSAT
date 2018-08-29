package models

import "log"
import "time"

func Insert(eqp string, date string, tech string) {

	timestamp := time.Now()

	res, err := db.Exec("INSERT INTO audit2 VALUES($1, $2, $3, $4)", timestamp, eqp, date, tech)
	if err != nil{
		log.Fatal(err)
	}

	rowCnt, err := res.RowsAffected()
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Rows affected = %d\n", rowCnt)
}


func AddScanner(name string, fab string, location string) {
	res, err := db.Exec("INSERT INTO scanners VALUES($1, $2, $3)", name, fab, location)
	if err != nil{
		log.Fatal(err)
	}

	rowCnt, err := res.RowsAffected()
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Rows affected = %d\n", rowCnt)
}