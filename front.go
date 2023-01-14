package main

import (
	"log"
	"net/http"
	"text/template"
)

func outputHTML(w http.ResponseWriter, filename string, data interface{}) {
	t, err := template.ParseFiles(filename)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	if err := t.Execute(w, data); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
}

func front(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		errorHandler(w, r, http.StatusNotFound)
		return
	}

	var alert string
	var info string

	err := r.ParseForm()
	if err != nil {
		log.Println(err)
	}

	if r.Form["healthswitch"] != nil {
		responder(w, r, true, switchHealth())

	} else {
		responder(w, r, true, "https://www.youtube.com/watch?v=dQw4w9WgXcQ")
	}

	myvar := map[string]interface{}{"HealthStatus": getHealth(), "alert": alert, "info": info}
	outputHTML(w, "static/index.html", myvar)
}
