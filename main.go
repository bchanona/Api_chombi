package main

import (
	"fmt"

	"github.com/bchanona/Api_chombi.git/src/helper/config"
)

func main() {
	db, err := config.ConnMySQL()

	if err != nil {
		fmt.Println("Database connection error: ", err)
	}

	println("Database connected", db)
}