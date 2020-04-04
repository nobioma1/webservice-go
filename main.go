package main

import (
	"fmt"

	"github.com/nobioma1/webservice-go/models"
)

func main() {
	u := models.User{
		ID:        2,
		FirstName: "Noble",
		LastName:  "Obioma",
	}

	fmt.Println(u)
}
