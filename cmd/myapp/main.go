package main

import (
	"log"
	"syncObsidian/internal/app"
)

func main() {
	if err := app.StartApp(); err != nil {
		log.Fatal(err)
	}
}
