package main

import (
	"html/template"
	"net/http"
	//"fmt"
)

//Render gohtml
func renderTemplate(w http.ResponseWriter, tmpl_str string, node *map[string][]string) {
	tmpl := make(map[string]*template.Template)
	tmpl[tmpl_str] = template.Must(template.ParseFiles(tmpl_str, "base.html"))
	tmpl[tmpl_str].ExecuteTemplate(w, "base", node)
}

func indexHandler() http.HandlerFunc {
	return func (w http.ResponseWriter, r *http.Request) {
		test := map[string][]string{
			"contacts" : {"1;Abhishek Chandratre;+17049068013;abhishek.chandratre@gmail.com;516 barton creek drive, APT E;0", "2;Tejas Konduri;+17049068013;tejas.konduri@gmail.com;516 barton creek drive, APT E;0"},
		}
		renderTemplate(w, "index.gohtml", &test)
	}
}

func pathHandler() http.HandlerFunc {
	return func (w http.ResponseWriter, r *http.Request) {	
		renderTemplate(w, "inputCsvFile.gohtml", nil)
	}
}

func main() {
	http.HandleFunc("/addPath",pathHandler())
	http.HandleFunc("/",indexHandler())
	//Start listening
	http.ListenAndServe(":8080",nil)
}


