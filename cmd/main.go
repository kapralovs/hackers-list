package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/kapralovs/hackers-list/internal/server"
)

func main() {
	port := ":8010"
	router := fiber.New()
	server := server.New(port, router)
	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}
