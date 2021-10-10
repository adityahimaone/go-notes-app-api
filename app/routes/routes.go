package routes

import (
	"github.com/gofiber/fiber/v2"
	"todolist/app/presenter/users"
)

type HandlerList struct {
	UserHandler users.Presenter
}

func (handler *HandlerList) Routes(fiberContext *fiber.App) {
	api := fiberContext.Group("/api/v1")
	//user Endpoint
	api.Post("/users/register", handler.UserHandler.Register)
}
