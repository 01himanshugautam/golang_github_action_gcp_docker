package main

import (
	"go_test/fiber"
	"log"
)

func main() {
	app := fiber.InitFiber()
	err := app.Listen("0.0.0.0:3000")
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
