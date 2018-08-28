package logging

import (
	"fmt"
	"log"
	"os"
)

func init() {
	nf, err := os.Create("VSATlog.log")
	if err != nil {
		fmt.Println(err)
	}
	log.SetOutput(nf)
}

func Logger(m string, err error) {
	fmt.Println("Starting logging.")
	if err != nil {
		log.Printf("%s; %s\r\n", m, err)
	} else {
		log.Printf("%s\r\n", m)
	}

}