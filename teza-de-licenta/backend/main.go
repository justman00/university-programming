package main

import (
	"log"

	"github.com/justman00/teza-de-licenta/cmd"
)

func main() {
	if err := cmd.Exec(); err != nil {
		log.Fatal(err)
	}
}
