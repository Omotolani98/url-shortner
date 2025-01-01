package main

import (
	"database/sql"
	"github.com/Omotolani98/url-shortner/internal/controllers"
	"github.com/Omotolani98/url-shortner/internal/db"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"net/http"
)

func main() {
	sqlite, err := sql.Open("sqlite3", "db.sqlite")
	if err != nil {
		log.Fatal(err)
	}
	defer sqlite.Close()

	if err := db.CreateTable(sqlite); err != nil {
		log.Fatal(err)
	}

	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		if request.URL.Path == "/" {
			controllers.ShowIndex(writer, request)
		} else {
			controllers.Redirect(sqlite).ServeHTTP(writer, request)
		}
	})

	http.HandleFunc("/shorten", controllers.Shorten(sqlite))
	log.Fatal(http.ListenAndServe(":8050", nil))
}
