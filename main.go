package main

import (
	"apiexample/src/config"
	"apiexample/src/routes"
	"log"

	"github.com/spf13/viper"
)

func main() {
	config.Setup()
	routes := routes.Init()

	log.Printf("Running on port %s", viper.GetString("PORT"))
	routes.Run()
}
