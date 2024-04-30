package main

import (
	"net/http"
)

func (app *application) CreateUserTable(w http.ResponseWriter, r *http.Request) {
	 app.DB.CreateUserTable()
	// if err != nil {
	// 	app.errorJSON(w, err)
	// 	return
	// }

}

