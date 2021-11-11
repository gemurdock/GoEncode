package main

import (
	"log"
	"os"

	"github.com/gemurdock/goencode/internal/appcli"
)

func main() {
	app := appcli.StartCLI()

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
