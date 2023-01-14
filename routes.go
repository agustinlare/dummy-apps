package main

import (
	"net/http"
)

func health(w http.ResponseWriter, r *http.Request) {
	if getHealth() {
		responder(w, r, true, "true")
	} else {
		responder(w, r, false, "false")
	}
}

func unhealthy(w http.ResponseWriter, r *http.Request) {
	responder(w, r, false, "cof...cof")
}

func ping(w http.ResponseWriter, r *http.Request) {
	responder(w, r, true, "pong")
}
