package main

import "apiexample/src/routes"

func main() {
	routes := routes.Init()
	routes.Run()
}
