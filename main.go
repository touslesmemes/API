package main

import (
	"log"

	"git.aprentout.com/touslesmemes/api2/actions"
)

func main() {
	app := actions.App()
	if err := app.Serve(); err != nil {
		log.Fatal(err)
	}
}
