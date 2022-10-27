package main

import (
	"log"
	"net/http"
	"path/filepath"
	"text/template"
)

func returnError(w http.ResponseWriter, err error) {
	w.WriteHeader(500)
	w.Write([]byte(err.Error()))
}

func serveMain(w http.ResponseWriter, r *http.Request) {
	p := filepath.Join("./", "index.html")

	tmpl, err := template.ParseFiles(p)
	if err != nil {
		returnError(w, err)
		return
	}

	data, err := getData()
	if err != nil {
		returnError(w, err)
		return
	}

	err = tmpl.ExecuteTemplate(w, "index", data)
	if err != nil {
		returnError(w, err)
		return
	}
}

func main() {
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.HandleFunc("/", serveMain)

	log.Println("Listening on :80...")
	err := http.ListenAndServe("0.0.0.0:80", nil)
	if err != nil {
		log.Fatal(err)
	}
}
