package main

import (
	"log"

	"github.com/oltarzewskik/tibivi/cmd"
)

func main() {
	if err := cmd.Run(); err != nil {
		log.Panicln(err)
	}
}
