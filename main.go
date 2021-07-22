package main

import (
	"alterra_store/configs"
	"alterra_store/routes"
)

func main() {
	configs.InitDB()
	e := routes.New()
	e.Start(":8000")
}
