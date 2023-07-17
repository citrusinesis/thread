package pkg

import (
	"github.com/citrusinesis/thread/pkg/domain"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"log"
)

func NewApp(port string) {
	app := fiber.New()
	app.Use(logger.New())
	domain.InjectRouter().RegisterRoute(app)

	err := app.Listen(port)
	if err != nil {
		log.Fatalf("Listen Error: %v", err)
	}
}
