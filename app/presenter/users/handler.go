package users

import (
	"github.com/gofiber/fiber/v2"
	"net/http"
	"todolist/app/presenter/users/request"
	"todolist/app/presenter/users/response"
	"todolist/bussiness/users"
	"todolist/helper/response_tmp"
)

type Presenter struct {
	serviceUser users.Service
}

func NewHandler(userService users.Service) *Presenter {
	return &Presenter{
		serviceUser: userService,
	}
}

func (handler *Presenter) Register(fiberContext *fiber.Ctx) error {
	var req request.User
	if err := fiberContext.BodyParser(&req); err != nil {
		response := response_tmp.APIResponse("Failed Bind", http.StatusBadRequest, "error", err)
		return fiberContext.JSON(response)
	}
	domain := request.ToDomain(req)
	resp, err := handler.serviceUser.Register(domain)
	if err != nil {
		response := response_tmp.APIResponse("Failed Bind", http.StatusInternalServerError, "error", err)
		return fiberContext.JSON(response)
	}
	responseRes := response_tmp.APIResponse("Success Register User", http.StatusOK, "Success", response.FromDomain(*resp))
	return fiberContext.JSON(responseRes)
}
