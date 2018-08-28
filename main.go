package main

import ( 
	"fmt"
	"github.com/bdalbrec/VSAT/models"
	"html/template"
	"net/http"
)

type results struct {
	List []result
}


type result struct {
	Location string
	Equipment string
	Date string
	Tech string
}


func (l *results) setList(list []result) {
	l.List = list
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("*.html"))
}

func main() {
	models.InitDB("sqlserver://VSATuser:P@55word@localhost:1433/SQLEXPRESS?database=VSAT&connection+timeout=30")

	http.HandleFunc("/", index)
	http.HandleFunc("/add", addEntry)
	http.HandleFunc("/scanners", scanners)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.Handle("/public/", http.StripPrefix("/public", http.FileServer(http.Dir("public/"))))
	http.ListenAndServe(":8080", nil)
}



func index(w http.ResponseWriter, req *http.Request) {

	ents, err := models.AllEntries()
	if err != nil {
		fmt.Println(err)
		http.Error(w, http.StatusText(500), 500)
		return
	}

	// create a results variable to pass into the template
	var data []result

	for _, e := range ents {
		fmt.Printf("%s\t %s\t %s\t %s\n", e.Location, e.Equipment, e.Date, e.Tech)

		var r result

		r.Location = e.Location
		r.Equipment = e.Equipment
		r.Date = e.Date
		r.Tech = e.Tech
		data = append(data, r)
	}

	var output results

	output.setList(data)

	fmt.Println(data)

	tpl.ExecuteTemplate(w, "index.html", output)
}


func addEntry(w http.ResponseWriter, req *http.Request) {
	loc := req.FormValue("location")
	eqp := req.FormValue("equipment")
	date := req.FormValue("date")
	tech := req.FormValue("tech")

	fmt.Printf("Inserting %s\t %s\t %s\t %s\n", loc, eqp, date, tech)
	models.Insert(loc, eqp, date, tech)

	http.Redirect(w, req, "/", http.StatusSeeOther)
	return
}

func scanners(w http.ResponseWriter, req *http.Request) {
	
	type scanner struct {
		Name string
		Fab string
		Location string
	}
	
	
	ents, err := models.GetScanners()
	if err != nil {
		fmt.Println(err)
		http.Error(w, http.StatusText(500), 500)
		return
	}

	// create variable to pass into the template
	var data []scanner

	for _, e := range ents {
		fmt.Printf("%s\t %s\t %s\n", e.Name, e.Fab, e.Location)

		var s scanner

		s.Name = e.Name
		s.Fab = e.Fab
		s.Location = e.Location

		data = append(data, s)
	}

	fmt.Println(data)

	tpl.ExecuteTemplate(w, "scanners.html", data)
}