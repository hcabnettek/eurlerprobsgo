package controllers

import (
	"net/http"
)

func RegisterControllers() {
	uc := newUseController()

	http.Handle("/users", *uc)
	http.Handle("/users/", *uc)
}
