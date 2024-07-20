package main

import (
	"fmt"
	"inLineCommandApp/app"
	"log"
	"os"
)

func main() {
	fmt.Println("In Line Command App")

	application := app.Build()

	if err := application.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
