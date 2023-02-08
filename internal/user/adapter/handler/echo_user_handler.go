package handler

import (
	"go-research/internal/pkg/util"
	"go-research/internal/user/port"
	"go-research/internal/user/service"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

func NewEchoUserHandler(userService service.UserService) port.EchoUserHandler {
	return &EchoUserHandler{
		service: userService,
	}
}

type EchoUserHandler struct {
	service service.UserService
}

func (u *EchoUserHandler) GetUsersByName(ctx echo.Context) error {
	pageIdx, err := strconv.Atoi(ctx.QueryParam("pageIdx"))
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, util.Reponse{
			Err: util.ErrorWrongTypePageSize.Error(),
		})
	}
	pageSize, err := strconv.Atoi(ctx.QueryParam("pageSize"))
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, util.Reponse{
			Err: util.ErrorWrongTypePageSize.Error(),
		})
	}
	name := ctx.Param("name")
	res, err := u.service.GetUsersByName(ctx.Request().Context(), name, pageSize, pageIdx)
	if err != nil {
		util.Logger.Errorf("internal server error: %v", err)
		return ctx.JSON(http.StatusInternalServerError, util.Reponse{
			Data: res,
			Err:  err.Error(),
		})
	} else {
		return ctx.JSON(http.StatusOK, util.Reponse{
			Data: res,
		})
	}
}
