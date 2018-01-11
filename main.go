package main

import (
	"log"

	"github.com/aquilinopintoa/sport_book/actions"
)

func main() {
	app := actions.App()
	if err := app.Serve(); err != nil {
		log.Fatal(err)
	}
}
