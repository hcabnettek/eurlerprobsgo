package main

import (
	"net/http"

	"github.com/hcabnettek/eulerprobs/controllers"
)

func main() {
	controllers.RegisterControllers()
	http.ListenAndServe(":3000", nil)
}
