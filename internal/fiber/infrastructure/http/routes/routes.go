package routes

import (
	"inventory/pkg"
	"inventory/pkg/middleware"

	"github.com/gofiber/fiber/v3"
)

func SetRoutes(app *fiber.App, db *pkg.Database) {
	SetUserRoutes(app, db)
	Apiv1 := app.Group("api/v1")
	SetClientRoutes(Apiv1, db)
	SetProviderRoutes(Apiv1, db)
	Apiv1.Use(middleware.VerifyTokenMiddleware)

}
