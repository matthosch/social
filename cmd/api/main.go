package main

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/matthosch/social/internal/env"
)

func main() {
	// Initialize environment variables
	err := godotenv.Load()
	if err != nil {
		log.Print("Error loading .env file")
	}

	cfg := config{
		addr: env.GetString("ADDR", ":8080"),
	}

	app := &application{
		config: cfg,
	}

	mux := app.mount()
	log.Fatal(app.run(mux))

}
