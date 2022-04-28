package main

import (
	"easyfitanalysis/app"
	"log"
	"time"
)

func main() {
	loc, err := time.LoadLocation("Europe/Berlin")
	if err != nil {
		log.Panicln("Error setting the correct timezone")
	}

	time.Local = loc

	a := app.App{}
	a.Init()

	a.Run(":8010")
}
