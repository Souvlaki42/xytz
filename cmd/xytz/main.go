package main

import (
	"log"
	"xytz/internal/app"
)

func main() {
	if err := app.Run(); err != nil {
		log.Fatal("unable to run the app")
	}
}
