/*
"Generated by OpenAI's chatGPT" - w/ modifications
*/
package main

import (
	"log"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	animalMiddleware "github.com/TimetoPretend54/go-chatgpt-copilot/animal-api/middleware"
	"github.com/TimetoPretend54/go-chatgpt-copilot/animal-api/routes"
)

func main() {
	// Create a new Echo instance
	e := echo.New()

	// Middleware
	e.Use(animalMiddleware.Logger())
	// e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Load animal routes
	routes.LoadAnimalRoutes(e)

	// Start server
	log.Fatal(e.Start(":8080"))
}
