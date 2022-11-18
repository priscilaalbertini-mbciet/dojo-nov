package utils

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

func StartServer(a *fiber.App) {
	fiberConnUrl, _ := ConnectionURLBuilder("fiber")

	if err := a.Listen(fiberConnUrl); err != nil {
		log.Printf("Server is not running! Error:%v", err)
	}
}
