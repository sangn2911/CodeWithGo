package port

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/labstack/echo"
)

type MuxUserHandler interface {
	GetUsersByName(http.ResponseWriter, *http.Request)
}

type EchoUserHandler interface {
	GetUsersByName(echo.Context) error
}

type FiberUserHandler interface {
	GetUsersByName(*fiber.Ctx) error
}
