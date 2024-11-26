package routes

import (
	"inventory/pkg"

	"github.com/gofiber/fiber/v3"
)

func SetRoutes(app *fiber.App, db *pkg.Database) {
	SetUserRoutes(app,db)
	//Apiv1 := app.Group("api/v1")

}