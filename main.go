package main

import (
	"easyfitanalysis/app"
	"flag"
	"log"
	"time"
)

func main() {

	portFlag := flag.String("Port", ":8080", "Specifies the port to use")

	loc, err := time.LoadLocation("Europe/Berlin")
	if err != nil {
		log.Panicln("Error setting the correct timezone")
	}

	time.Local = loc

	a := app.App{}
	a.Init()

	log.Printf("Starting server on port: %s", *portFlag)
	a.Run(*portFlag)
}
