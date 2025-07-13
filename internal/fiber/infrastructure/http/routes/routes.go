package routes

import (
	"inventory/pkg"
	"inventory/pkg/middleware"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/cors"
	"github.com/gofiber/fiber/v3/middleware/static"
)

func SetRoutes(app *fiber.App, db *pkg.Database) error {
	app.Use(middleware.PaginationMiddleware)
	app.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{"Origin, Content-Type, Accept, Authorization"},
		AllowMethods: []string{
			fiber.MethodGet,
			fiber.MethodPost,
			fiber.MethodHead,
			fiber.MethodPut,
			fiber.MethodDelete,
			fiber.MethodPatch,
			fiber.MethodOptions,
		},
	}))

	SetUserRoutes(app, db)
	Apiv1 := app.Group("api/v1")
	SetClientRoutes(Apiv1, db)
	SetProviderRoutes(Apiv1, db)
	SetProductsRoutes(Apiv1, db)
	SetCategoryRoutes(Apiv1, db)
	SetOrderRoutes(Apiv1, db)
	SetupTransactionRoutes(Apiv1, db)
	app.Get("/media*", static.New("./media"))
	// Apiv1.Use(middleware.VerifyTokenMiddleware)

	return nil
}
