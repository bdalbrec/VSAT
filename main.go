package main

import ( 
	"fmt"
	"github.com/bdalbrec/VSAT/models"
	"html/template"
	"net/http"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("*.html"))
}

func main() {
	models.InitDB("sqlserver://VSATuser:P@55word@localhost:1433/SQLEXPRESS?database=VSAT&connection+timeout=30")

	http.HandleFunc("/", index)
	http.HandleFunc("/add", addEntry)
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

	for _, e := range ents {
		fmt.Printf("%s\t %s\t %s\t %s\n", e.Location, e.Equipment, e.Date, e.Tech)
	}

	tpl.ExecuteTemplate(w, "index.html", nil)
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

