package main

import (
	"fmt"
	"net/http"
)

func HealthCheck(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "healthy")
}

func ErrorPage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "404 pogger")
	http.NotFound(w, r)
}

func main() {
	http.HandleFunc("/healthz", HealthCheck)
	http.HandleFunc("/", ErrorPage)
	http.ListenAndServe(":8000", nil)
}