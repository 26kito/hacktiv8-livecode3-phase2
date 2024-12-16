package main

import (
	"hacktiv8-lc3-p2/config"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	config.InitDB()
	db := config.DB

	config.Routes(db)
}
