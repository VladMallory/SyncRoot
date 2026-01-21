package main

import (
	"log"
	"syncroot/internal/app"
)

func main() {
	if err := app.StartApp(); err != nil {
		log.Fatal(err)
	}
}
