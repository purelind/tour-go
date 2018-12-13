package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"

	_ "github.com/go-sql-driver/mysql"
)

const (
	DBHost  = "127.0.0.1"
	DBPort  = ":3308"
	DBUser  = "root"
	DBPass  = "123456"
	DBDbase = "http-web"
	PORT    = ":8080"
)

var database *sql.DB

type Page struct {
	Title   string
	Content string
	Date    string
}

func ServePage(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	pageID := vars["id"]
	thisPage := Page{}
	fmt.Println(pageID)
	err := database.QueryRow("SELECT page_title, page_content, page_date FROM pages WHERE id=?", pageID).Scan(&thisPage.Title, &thisPage.Content, &thisPage.Date)
	if err != nil {
		log.Println("Couldn't get page:" + pageID)
		log.Println(err)
	}
	html := `<html><head><title>` + thisPage.Title +
		`</title></head><body><h1>` + thisPage.Title + `</h1><div>` +
		thisPage.Content + `</div></body></html>`
	fmt.Fprintln(w, html)
}

func main() {
	dbConn := fmt.Sprintf("%s:%s@tcp(%s%s)/%s", DBUser, DBPass, DBHost, DBPort, DBDbase)
	fmt.Println(dbConn)
	db, err := sql.Open("mysql", dbConn)
	if err != nil {
		log.Println("Couldn't connect!" + DBDbase)
		log.Println(err)
	}
	database = db

	routes := mux.NewRouter()
	routes.HandleFunc("/page/{id:[0-9]+}", ServePage)
	http.Handle("/", routes)
	http.ListenAndServe(PORT, nil)
}
