package main

import (
	"context"
	"go-research/internal/app"
	"go-research/internal/pkg/config"
	"go-research/internal/pkg/util"
	"net/http"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

// Typing go run . env/yml mux/echo (to run your setup desired)
func main() {
	cfg_option, pkg_option := config.GetArgs(os.Args)
	util.InitializeLogger()
	cfg := config.ProvideConfig(cfg_option)
	StartApp(pkg_option, cfg)
}

func StartApp(pkg_option string, cfg *config.Config) {
	switch pkg_option {
	case "mux":
		StartMuxApp(cfg)
	case "echo":
		StartEchoApp(cfg)
	case "fiber":
		StartFiberApp(cfg)
	default:
		StartMuxApp(cfg)
	}
}

func StartMuxApp(cfg *config.Config) {
	var r *mux.Router
	var err error
	r = mux.NewRouter()
	r.Use(util.MuxLoggingMiddleware)
	r, err = app.NewMuxRoute(r, context.Background(), cfg)
	if err != nil {
		panic(err)
	}
	headersOk := handlers.AllowedHeaders([]string{"Origin", "X-Requested-With", "Content-Type", "Accept", "Authorization", "Signature", "Cache-Control"})
	methodsOk := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"})
	srv := &http.Server{
		Handler: handlers.CORS(headersOk, methodsOk)(r),
		Addr:    "127.0.0.1:3000",
	}
	util.Logger.Infof("server started, listening on port 3000")
	util.Logger.Fatal(srv.ListenAndServe())
}

func StartEchoApp(cfg *config.Config) {
	var e *echo.Echo
	var err error
	e = echo.New()
	e.Use(util.EchoLoggingMiddleware)
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowMethods: []string{echo.GET, echo.HEAD, echo.POST, echo.PUT, echo.PATCH, echo.DELETE, echo.OPTIONS},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderXRequestedWith, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAuthorization, "Signature", "Cache-Control"},
	}))
	e, err = app.NewEchoRoute(e, cfg)
	if err != nil {
		panic(err)
	}
	util.Logger.Fatal(e.Start("127.0.0.1:3000"))
}

func StartFiberApp(cfg *config.Config) {
	var f *fiber.App
	var err error
	f = fiber.New()
	f.Use(fiber.Handler(util.FiberLoggingMiddleware))
	f.Use(cors.New(cors.Config{
		AllowHeaders: "Origin,X-Requested-With,Content-Type,Accept,Authorization,Signature,Cache-Control",
		AllowMethods: "GET,HEAD,POST,PUT,PATCH,DELETE,OPTIONS",
	}))
	f, err = app.NewFiberRoute(f, cfg)
	if err != nil {
		panic(err)
	}
	util.Logger.Fatal(f.Listen("127.0.0.1:3000"))
}
