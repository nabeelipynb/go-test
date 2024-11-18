package main

import (
	"embed"
	"fmt"
	"log"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/filesystem"
)

// Embed the entire dist folder inside the frontend directory
//go:embed frontend/dist/*
var staticFiles embed.FS

// Handler function for Vercel
func Handler(w http.ResponseWriter, r *http.Request) {
	// Create a new Fiber app instance for each request
	app := fiber.New()

	// Serve the index.html file for the root URL
	app.Get("/", func(c *fiber.Ctx) error {
		// Ensure that the index.html file is served from the embedded filesystem
		return c.SendFile("frontend/dist/index.html")
	})

	// Serve the rest of the static files in the dist folder
	app.Use("/", filesystem.New(filesystem.Config{
		Root:       http.FS(staticFiles), // Wrap the embedded FS with http.FS
		PathPrefix: "frontend/dist",      // Ensure the correct path prefix is set for embedded files
		Browse:     false,                // Disable directory browsing
	}))

	// Start the Fiber app
	if err := app.Listener(http.ResponseWriter(w), r); err != nil {
		log.Fatal(err)
	}
}

func main() {
	// In a typical server environment, you would use app.Listen here.
	// For serverless environments like Vercel, we use the Handler function to be invoked by Vercel.
	fmt.Println("Serverless function is ready")

	// If you test it locally, you can call Handler directly:
	// Handler(nil, nil) // Only for testing locally
}
