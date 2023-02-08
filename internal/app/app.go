package app

import (
	"go-research/internal/pkg/config"
	"go-research/internal/pkg/database"
	userhandler "go-research/internal/user/adapter/handler"
	useradapter "go-research/internal/user/adapter/repository"
	userport "go-research/internal/user/port"
	"go-research/internal/user/service"
)

type ApplicationContext struct {
	User userport.MuxUserHandler
}

type EchoApplicationContext struct {
	User userport.EchoUserHandler
}

type FiberApplicationContext struct {
	User userport.FiberUserHandler
}

func NewMuxApp(conf *config.Config) (*ApplicationContext, error) {
	db, err := database.ProvideDBConfig(conf.Sql)
	if err != nil {
		panic(err)
	}
	userRepository := useradapter.NewUserRepository(db)
	userService := service.NewUserService(db, userRepository)
	userHandler := userhandler.NewUserHandler(userService)
	return &ApplicationContext{
		User: userHandler,
	}, nil
}

func NewEchoApp(conf *config.Config) (*EchoApplicationContext, error) {
	db, err := database.ProvideDBConfig(conf.Sql)
	if err != nil {
		panic(err)
	}
	userRepository := useradapter.NewUserRepository(db)
	userService := service.NewUserService(db, userRepository)
	userHandler := userhandler.NewEchoUserHandler(userService)
	return &EchoApplicationContext{
		User: userHandler,
	}, nil
}

func NewFiberApp(conf *config.Config) (*FiberApplicationContext, error) {
	db, err := database.ProvideDBConfig(conf.Sql)
	if err != nil {
		panic(err)
	}
	userRepository := useradapter.NewUserRepository(db)
	userService := service.NewUserService(db, userRepository)
	userHandler := userhandler.NewFiberUserHandler(userService)
	return &FiberApplicationContext{
		User: userHandler,
	}, nil
}
