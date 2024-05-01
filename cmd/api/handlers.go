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

func (app *application) Hello(w http.ResponseWriter, r *http.Request) {
	// Attempt to create the user table
	//err := app.DB.CreateUserTable()
	// if err != nil {
	// 		// If there is an error, return an error message in JSON format
	// 		// app.errorJSON(w, err)
	// 		return
	// }

	// If everything is okay, send "Hello" as the response
	w.WriteHeader(http.StatusOK)  // Optional: explicitly set HTTP status to 200 OK
	w.Header().Set("Content-Type", "text/plain")  // Set content type as plain text
	w.Write([]byte("Hello"))  // Write "Hello" to the response body
}

