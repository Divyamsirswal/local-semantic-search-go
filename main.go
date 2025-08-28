package main

import (
	"fmt"
	"log"
	"os"
	"semantic-search-api/database"
	"semantic-search-api/handlers"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/ollama/ollama/api"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}
	dbURL := os.Getenv("DATABASE_URL")

	dbpool, err := database.NewConnection(dbURL)
	if err != nil {
		log.Fatalf("Could not connect to database: %v", err)
	}
	defer dbpool.Close()

	ollamaClient, err := api.ClientFromEnvironment()
	if err != nil {
		log.Fatalf("Could not create Ollama client: %v", err)
	}

	database.SeedDatabase(dbpool, ollamaClient)
	searchHandler := handlers.MakeSearchHandler(dbpool, ollamaClient)

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.GET("/search", searchHandler)

	fmt.Println("Server is running on http://localhost:8081")

	e.Logger.Fatal(e.Start(":8081"))
}
