package main

import (
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

const (
	PORT = ":8080"
)

func pageHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	pageID := vars["id"]
	fileName := "files/" + pageID + ".html"
	_, err := os.Stat(fileName)
	if err != nil {
		fileName = "files/404.html"
	}
	http.ServeFile(w, r, fileName)
}

func main() {
	rtr := mux.NewRouter()
	"js;html;htm;css"
	rtr.HandleFunc("/pages/{id:[0-9]+}", pageHandler)
	rtr.HandleFunc("/action", actionPageHandler)
	rtr.HandleFunc("/static", staticPageHandler)
	http.Handle("/", rtr)
	http.ListenAndServe(PORT, nil)
}
