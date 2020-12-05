package bootstrap

import (
	"github.com/gofiber/fiber/v2"
	"gofiber/server/bootstrap/routers"
	"gofiber/server/handlers"
	"net/http"
)

func (boot Bootstrap) RegisterRouters() {
	handler := handlers.Handler{
		FiberApp:   boot.App,
		UcContract: &boot.UcContract,
		Jwe:        boot.JweCred,
		Db:         boot.DB,
		Validator:  boot.Validator,
		Translator: boot.Translator,
	}

	//test purpose
	boot.App.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.Status(http.StatusOK).JSON("work")
	})

	apiV1 := boot.App.Group("/api/v1")

	//user routes
	userRoutes := routers.UserRoutes{
		RouterGroup: apiV1,
		Handler:     handler,
	}
	userRoutes.RegisterRoute()
}
