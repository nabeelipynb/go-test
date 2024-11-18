package main

import (
	"embed"
	"log"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/filesystem"
)

// Embed the entire dist folder inside the frontend directory
//go:embed frontend/dist/*
var staticFiles embed.FS

func main() {
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

	// Print out all files in the embedded "frontend/dist" directory for debugging
	fs, _ := staticFiles.ReadDir("frontend/dist")
	for _, file := range fs {
		log.Println(file.Name())
	}

	log.Fatal(app.Listen("127.0.0.1:3000"))
}
