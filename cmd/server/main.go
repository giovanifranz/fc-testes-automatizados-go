package main

import (
	"database/sql"
	"net/http"

	"github.com/giovanifranz/testes-go/internal/infra/controller"
)

func main() {
	db, err := sql.Open("sqlite3", "db.sqlite3")

	if err != nil {
		panic(err)
	}

	defer db.Close()

	controllers := controller.NewBaseHandler(db)
	http.HandleFunc("/clients", controllers.CreateClientHandler)
	http.ListenAndServe(":3000", nil)
}
