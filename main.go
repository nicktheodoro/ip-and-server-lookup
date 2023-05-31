package main

import (
	"ip-and-server-lookup/app"
	"log"
	"os"
)

func main() {
	application := app.GenerateApp()

	if error := application.Run(os.Args); error != nil {
		log.Fatal(error)
	}
}
