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
			"test" : {"hello1", "hello2"},
		}
		renderTemplate(w, "index.gohtml", &test)
	}
}

func main() {
	http.HandleFunc("/",indexHandler())
	//Start listening
	http.ListenAndServe(":8080",nil)
}


