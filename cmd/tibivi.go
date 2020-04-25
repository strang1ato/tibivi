package main

import (
	"log"

	"github.com/oltarzewskik/tibivi/tibivi"
)

func main() {
	if err := tibivi.Run(); err != nil {
		log.Panicln(err)
	}
}
