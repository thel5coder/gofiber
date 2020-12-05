package routers

import (
	"github.com/gofiber/fiber/v2"
	"gofiber/server/handlers"
)

type UserRoutes struct {
	RouterGroup fiber.Router
	Handler     handlers.Handler
}

//register user routes
func (route UserRoutes) RegisterRoute() {
	handler := handlers.UserHandler{Handler: route.Handler}

	userRoutes := route.RouterGroup.Group("/user")
	userRoutes.Get("", handler.Browse)
	userRoutes.Get("/:id", handler.Read)
	userRoutes.Put("/:id", handler.Edit)
	userRoutes.Post("", handler.Add)
	userRoutes.Delete("/:id", handler.Delete)
}
