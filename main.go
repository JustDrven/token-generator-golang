package main

import (
	"petrnemecek.dev/token-generator/controller"
	"petrnemecek.dev/token-generator/models"
)

func main() {

	server := models.Server{
		Addr: ":9090",
	}

	controller.Init()

	server.Listen()

}