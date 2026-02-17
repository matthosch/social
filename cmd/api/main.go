package main

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/matthosch/social/internal/env"
	"github.com/matthosch/social/internal/store"
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

	store := store.NewStorage(nil)

	app := &application{
		config: cfg,
		store:  store,
	}

	mux := app.mount()
	log.Fatal(app.run(mux))

}
