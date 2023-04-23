package main

import (
	"C8/app"
	"C8/config"

	"github.com/joho/godotenv"
)

func init() {
	godotenv.Load()

	err := config.InitPostgres()
	if err != nil {
		panic(err)
	}
}

func main() {
	app.StartApplication()
}
