package main

import (
	"html/template"
	"log"
	"net/http"
)

var templates = template.Must(template.ParseFiles(
	"static/index.html",
	"static/upload.html",
	"static/report.html",
))

func main() {
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/upload/", uploadHandler)
	http.HandleFunc("/report/", reportHandler)

	addr := ":8080"
	log.Printf("Starting server at %s\n", addr)
	log.Fatal(http.ListenAndServe(addr, nil))
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	renderHtml(w, "index")
}

func uploadHandler(w http.ResponseWriter, r *http.Request) {
	renderHtml(w, "upload")
}

func reportHandler(w http.ResponseWriter, r *http.Request) {
	renderHtml(w, "report")
}

func renderHtml(w http.ResponseWriter, tmpl string) {
	err := templates.ExecuteTemplate(w, tmpl + ".html", nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
