package main

import (
	"log"
	"vktest/cmd/app"
)

func main() {
	err := app.Run()
	if err != nil {
		log.Fatal(err)
	}

}
