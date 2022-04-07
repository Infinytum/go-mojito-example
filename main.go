package main

import (
	_ "github.com/infinytum/go-scalar/autoload"
	"github.com/infinytum/go-scalar/service/logger"

	"github.com/infinytum/go-scalar"
	"github.com/infinytum/go-scalar-example/handlers"
)

func init() {
	// Override the default zerolog logger with the
	// builtin logger implementation
	// We'd suggest not doing this though as the logs will look awful
	scalar.Register(func() scalar.Logger {
		return logger.NewBuiltin()
	}, true)
}

var (
	Address = "0.0.0.0:8123"
	Logger  scalar.Logger
	Router  scalar.Router
)

func main() {
	Logger = scalar.DefaultLogger()
	Router = scalar.DefaultRouter()

	Logger.Info("Registering application routes")

	Router.WithMiddleware(scalar.RequestLogger)

	Router.GET("/assets/*path", scalar.AssetsHandler)
	Router.GET("/", handlers.Home)

	Logger.Infof("Server has started on %s", Address)
	Logger.Error(Router.ListenAndServe(Address))
}
