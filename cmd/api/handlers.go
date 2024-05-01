package main

import (
	"fmt"
	"goauthlogin/internal/models"
	"net/http"
)

func (app *application) CreateUserTable(w http.ResponseWriter, r *http.Request) {
	fmt.Println("handler db created")

	app.DB.CreateUserTable()
	fmt.Println("db created")
	// if err != nil {
	// 	app.errorJSON(w, err)
	// 	return
	// }

}

func (app *application) authenticate(w http.ResponseWriter, r *http.Request) {
	//read json payload
	var requestPayload struct {
		UserName    string `json:"username"`
		Password string `json:"password"`
	}

	err := app.readJSON(w, r, &requestPayload)
	if err != nil {
		app.errorJSON(w, err, http.StatusBadRequest)
		return
	}

	//validate user against db
	user, err := app.DB.GetUserName(requestPayload.UserName)
	if err != nil {
		response := models.LoginResponse{
            User:    nil,
            Message: "Authentication unsuccessful: failed to get user",
            Success: false,
        }
        app.writeJSON(w, http.StatusBadRequest, response)
        return
	}

	//check password
	valid, err := user.PasswordMatches(requestPayload.Password)
	if err != nil || !valid {
		response := models.LoginResponse{
            User:    nil,
            Message: "Authentication unsuccessful: invalid credentials",
            Success: false,
        }
        app.writeJSON(w, http.StatusBadRequest, response)
        return
	}

	// If everything is fine, send back the success response
    response := models.LoginResponse{
        User:    user,
        Message: "Authentication successful",
        Success: true,
    }

	app.writeJSON(w, http.StatusOK, response)
}

func (app *application) createUser(w http.ResponseWriter, r *http.Request) {
	// Read JSON payload containing user data
	var user models.User
	err := app.readJSON(w, r, &user)
	if err != nil {
		app.errorJSON(w, err, http.StatusBadRequest)
		return
	}

	// Create the user in the database
	err = app.DB.CreateUser(&user)
	if err != nil {
		app.errorJSON(w, err, http.StatusInternalServerError)
		return
	}

	// Respond with a success message or user data
	app.writeJSON(w, http.StatusCreated, map[string]string{"message": "User created successfully"})
}
