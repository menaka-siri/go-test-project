package main

import (
	"fmt"

	"github.com/pluralsight/webservice/models"
)

func main() {
	fmt.Println("Hello from a module, Gophers!")

	u := models.User{
		ID:        2,
		FirstName: "Tricia",
		LastName:  "McMillan",
	}

	fmt.Println(u)

}
