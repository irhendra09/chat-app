package bootstrap

import (
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/logger"
	"github.com/gofiber/fiber/v3/middleware/recover"
	"github.com/irhendra09/chat-app/pkg/database"
	"github.com/irhendra09/chat-app/pkg/env"
	"github.com/irhendra09/chat-app/pkg/router"
)

func NewAplication() *fiber.App {
	env.SetupEnvFile()
	database.SetupPostgres()

	app := fiber.New(fiber.Config{})
	app.Use(recover.New())
	app.Use(logger.New())

	router.InstallRouter(app)
	return app
}
