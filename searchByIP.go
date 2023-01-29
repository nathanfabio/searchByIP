package main

import (
	"github.com/nathanfabio/searchByIP.git/app"
	"os"
	"log"
)

func main() {
	application := app.Generate()
	if err := application.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}