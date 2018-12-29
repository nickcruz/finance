package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/upload/", uploadHandler)
	http.HandleFunc("/report/", reportHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
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
	templateFile := fmt.Sprintf("static/%s.html", tmpl)
	t, err := template.ParseFiles(templateFile)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = t.Execute(w, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
