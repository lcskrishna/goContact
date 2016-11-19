package main

import (
	"html/template"
	"net/http"
	"fmt"
)

var csvPath string = ""

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

func listHandler() http.HandlerFunc {
	return func (w http.ResponseWriter, r *http.Request) {	
		test := map[string][]string{
			"contacts" : {"1;Abhishek Chandratre;+17049068013;abhishek.chandratre@gmail.com;516 barton creek drive, APT E;0", "2;Tejas Konduri;+17049068013;tejas.konduri@gmail.com;516 barton creek drive, APT E;0"},
		}
		
		renderTemplate(w, "list.gohtml", &test)
	}
}

func addContactHandler() http.HandlerFunc {
	return func (w http.ResponseWriter, r *http.Request) {	
		
		fmt.Println("Inside addContactHandler method...");
		
		if r.Method == "GET"{
			fmt.Println("Inside GET method...");
			renderTemplate(w, "addContact.gohtml", nil)
		}
		if r.Method == "POST"{
			fmt.Println("Inside POST method...");
			r.ParseForm()
			//Get input value of csv path
			name := r.Form["name"][0]
			phoneNo := r.Form["phoneNo"][0]
			email := r.Form["email"][0]
			address := r.Form["address"][0]
			
			fmt.Println(name)
			fmt.Println(phoneNo)			
			fmt.Println(email)
			fmt.Println(address)

			var record string = name + ";" + phoneNo + ";" + email + ";" + address  
			fmt.Println(record)
			
			renderTemplate(w, "list.gohtml", nil)		
		}
	}
}

func pathHandler() http.HandlerFunc {
	return func (w http.ResponseWriter, r *http.Request) {	
		
		fmt.Println("Inside pathHandler method...");
		
		if r.Method == "GET"{
			fmt.Println("Inside GET method...");
			renderTemplate(w, "inputCsvFile.gohtml", nil)		
		}
		
		if r.Method == "POST"{
			fmt.Println("Inside POST method...");
			r.ParseForm()
			//Get input value of csv path
			csvPath := r.Form["csvPath"][0]
			
			fmt.Println(csvPath)			
			
			renderTemplate(w, "list.gohtml", nil)		
		}
	}
}

func main() {
	fmt.Println("Starting Application...");
	http.HandleFunc("/addPath",pathHandler())
	http.HandleFunc("/addContact",addContactHandler())
	http.HandleFunc("/list", listHandler())
	http.HandleFunc("/",indexHandler())
	//Start listening
	http.ListenAndServe(":8080",nil)
}


