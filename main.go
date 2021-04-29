package main

import (
	"log"

	"github.com/strang1ato/tibivi/cmd"
)

func main() {
	if err := cmd.Run(); err != nil {
		log.Panicln(err)
	}
}
