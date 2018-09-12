package main

import ( 
	"fmt"
	"github.com/bdalbrec/VSAT/models"
	"html/template"
	"net/http"
)

type results struct {
	List []result
	Choices []string
}


type result struct {
	Building string
	Location string
	Equipment string
	Date string
	Tech string
}


func (l *results) setList(list []result) {
	l.List = list
}

func (l * results) setChoices(cl []string) {
	l.Choices = cl
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*.html"))
}

func main() {
	models.InitDB("sqlserver://VSATuser:P@55word@localhost:1433/SQLEXPRESS?database=VSAT&connection+timeout=30")

	http.HandleFunc("/", index)
	http.HandleFunc("/add", addEntry)
	http.HandleFunc("/scanners", scanners)
	http.HandleFunc("/addScanner", addScanner)
	http.HandleFunc("/history", fullHistory)
	http.HandleFunc("/addChecklist", handleCheckboxes)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.Handle("/public/", http.StripPrefix("/public", http.FileServer(http.Dir("public/"))))
	http.ListenAndServe(":8080", nil)
}



func index(w http.ResponseWriter, req *http.Request) {

	// a output variable to pass into the template
	var output results

	// grab the list of all audits from the db
	ents, err := models.LastEntries()
	if err != nil {
		fmt.Println(err)
		http.Error(w, http.StatusText(500), 500)
		return
	}

	// create a []result variable to pass into the template
	var data []result

	// put the returned audits into the data variable
	for _, e := range ents {
		var r result

		r.Equipment = e.Equipment
		r.Building = e.Building
		r.Date = e.Date
		r.Tech = e.Tech
		r.Location = e.Location
		data = append(data, r)
	}

	// grab the list of scanners from the db
	s, err := models.ScannerList()
	if err != nil {
		fmt.Println(err)
		http.Error(w, http.StatusText(500), 500)
		return
	}
	

	// create a []string variable to hold the list of scanners
	var scanners []string

	// put the scanners into the slice
	for _, v := range s {
		var s string
		s = v.Name 
		scanners = append(scanners, s)
	}



	// build the output struct
	output.setList(data)
	output.setChoices(scanners)

	tpl.ExecuteTemplate(w, "checkboxes.html", output)
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

func handleCheckboxes(w http.ResponseWriter, req *http.Request) {
	req.ParseForm()
	selected := req.Form["selectedTools"]
	date := req.FormValue("date")
	tech := req.FormValue("tech")

	for _, v := range selected {
		fmt.Printf("Inserting %s\t %s\t %s into audits history.\n", v, date, tech)
		models.Insert(v, date, tech)
	}

	http.Redirect(w, req, "/", http.StatusSeeOther)
	return
}

func fullHistory(w http.ResponseWriter, req *http.Request) {
		// a output variable to pass into the template
		var output results

		// grab the list of all audits from the db
		ents, err := models.AllEntries()
		if err != nil {
			fmt.Println(err)
			http.Error(w, http.StatusText(500), 500)
			return
		}
	
		// create a []result variable to pass into the template
		var data []result
	
		// put the returned audits into the data variable
		for _, e := range ents {
			var r result
	
			r.Equipment = e.Equipment
			r.Building = e.Building
			r.Date = e.Date
			r.Tech = e.Tech
			r.Location = e.Location
			data = append(data, r)
		}
	
	
	
		// build the output struct
		output.setList(data)
	
		tpl.ExecuteTemplate(w, "history.html", output)
}