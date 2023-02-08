package app

import (
	"context"
	"go-research/internal/pkg/config"

	"github.com/gofiber/fiber/v2"
	"github.com/gorilla/mux"
	"github.com/labstack/echo"
)

const (
	GET    = "GET"
	POST   = "POST"
	PUT    = "PUT"
	DELETE = "DELETE"
)

func NewMuxRoute(r *mux.Router, ctx context.Context, conf *config.Config) (*mux.Router, error) {
	app, err := NewMuxApp(conf)
	if err != nil {
		return r, err
	}
	userRouter := r.PathPrefix("/users").Subrouter()
	userRouter.HandleFunc("/{name}", app.User.GetUsersByName).Methods(GET)
	return r, nil
}

func NewEchoRoute(e *echo.Echo, conf *config.Config) (*echo.Echo, error) {
	app, err := NewEchoApp(conf)
	if err != nil {
		return e, err
	}
	userRouter := e.Group("/users")
	userRouter.GET("/:name", app.User.GetUsersByName)
	return e, nil
}

func NewFiberRoute(f *fiber.App, conf *config.Config) (*fiber.App, error) {
	app, err := NewFiberApp(conf)
	if err != nil {
		return f, err
	}
	userRouter := f.Group("/users")
	userRouter.Get("/:name", app.User.GetUsersByName)
	return f, nil
}
