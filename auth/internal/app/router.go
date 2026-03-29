package app

import "github.com/gofiber/fiber/v2"

func (a *App) registerRoutes() {
	a.Fiber.Get("/health", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"status": "ok"})
	})
}
