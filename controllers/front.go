package controllers

import "net/http"

// RegisterControllers function to handle requests to the server
func RegisterControllers() {
	uc := newUserController()

	http.Handle("/users", *uc)
	http.Handle("/users/", *uc)
}
