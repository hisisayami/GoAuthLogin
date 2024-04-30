package main

import (
	"fmt"
	"goauthlogin/internal/repository"
	"goauthlogin/internal/repository/dbrepo"
	"log"
	"net/http"
)

const port = 9090

type application struct {
	DB repository.DatabaseRepo
}

func main() {

	//set application config
	var app application

	// Initialize database connection
	db, err := connectToDB()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	app.DB = &dbrepo.SqlDBRepo{DB: db}
	defer app.DB.Connection().Close()

	fmt.Println("Starting 9090...")

	//start a web server
	err = http.ListenAndServe(fmt.Sprintf(":%d", port), app.routes())
	if err != nil {
		log.Fatal(err)
	}

}
