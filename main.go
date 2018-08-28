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
	http.HandleFunc("/addScanner", addScanner)
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
		var r result

		r.Equipment = e.Equipment
		r.Date = e.Date
		r.Tech = e.Tech
		data = append(data, r)
	}

	var output results

	output.setList(data)

	tpl.ExecuteTemplate(w, "index.html", output)
}


func addEntry(w http.ResponseWriter, req *http.Request) {
	eqp := req.FormValue("equipment")
	date := req.FormValue("date")
	tech := req.FormValue("tech")

	fmt.Printf("Inserting %s\t %s\t %s into audits history.\n", eqp, date, tech)
	models.Insert(eqp, date, tech)

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
		var s scanner

		s.Name = e.Name
		s.Fab = e.Fab
		s.Location = e.Location

		data = append(data, s)
	}

	tpl.ExecuteTemplate(w, "scanners.html", data)
}

func addScanner(w http.ResponseWriter, req *http.Request) {
	name := req.FormValue("equipment")
	fab := req.FormValue("fab")
	loc := req.FormValue("location")

	fmt.Printf("Inserting %s\t %s\t %s into scanners.\n", name, fab, loc)
	models.AddScanner(name, fab, loc)

	http.Redirect(w, req, "/scanners", http.StatusSeeOther)
}