package pkg

import (
	"encoding/json"
	//"time"

	"github.com/gofiber/fiber/v3"
)

func InitFiber() (*fiber.App, error) {

	appFiber := fiber.New(fiber.Config{
		JSONEncoder: json.Marshal,
		JSONDecoder: json.Unmarshal,
		CaseSensitive: true,
		StrictRouting: false,
		ServerHeader: "MyFiberApp",
		//ReadTimeout:  10 * time.Second, // Tiempo máximo para leer una solicitud
		//WriteTimeout: 10 * time.Second, // Tiempo máximo para escribir una respuesta
	})

	return appFiber, nil
}