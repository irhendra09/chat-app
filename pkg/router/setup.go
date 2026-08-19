package router

import "github.com/gofiber/fiber/v3"

func InstallRouter(app *fiber.App) {
	setup(app, NewApiRouter())
}

func setup(app *fiber.App, router ...*ApiRouter) {
	for _, r := range router {
		r.InstallRouter(app)
	}
}
