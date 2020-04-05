package main

import (
	"net/http"

	"github.com/nobioma1/webservice-go/controllers"
)

func main() {
	controllers.RegisterControllers()
	http.ListenAndServe(":3000", nil)
}
