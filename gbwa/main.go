package main

import (
	"fmt"
	"net/http"
	"time"
)

const (
	Port = ":8080"
)

func serverDynamic(w http.ResponseWriter, r *http.Request) {
	response := "The time is now " + time.Now().String()
	fmt.Fprintln(w, response)
}

func serverStatic(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "static.html")
}

func main() {
	http.HandleFunc("/static", serverStatic)
	http.HandleFunc("/", serverDynamic)
	http.ListenAndServe(Port, nil)
}
