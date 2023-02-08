package handler

import (
	"go-research/internal/pkg/util"
	"go-research/internal/user/port"
	"go-research/internal/user/service"
	"net/http"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func NewFiberUserHandler(userService service.UserService) port.FiberUserHandler {
	return &FiberUserHandler{
		service: userService,
	}
}

type FiberUserHandler struct {
	service service.UserService
}

func (u *FiberUserHandler) GetUsersByName(ctx *fiber.Ctx) error {
	pageIdx, err := strconv.Atoi(ctx.Query("pageIdx"))
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(util.Reponse{
			Err: util.ErrorWrongTypePageIdx.Error(),
		})
	}
	pageSize, err := strconv.Atoi(ctx.Query("pageSize"))
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(util.Reponse{
			Err: util.ErrorWrongTypePageSize.Error(),
		})
	}
	name := ctx.Query("name")
	res, err := u.service.GetUsersByName(ctx.Context(), name, pageSize, pageIdx)
	if err != nil {
		util.Logger.Errorf("internal server error: %v", err)
		return ctx.Status(http.StatusInternalServerError).JSON(util.Reponse{
			Data: res,
			Err:  err.Error(),
		})
	} else {
		return ctx.Status(http.StatusOK).JSON(util.Reponse{
			Data: res,
		})
	}
}
