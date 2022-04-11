package main

import (
	"github.com/infinytum/go-mojito"
	"github.com/infinytum/go-mojito-example/handlers"
)

func init() {
	// Override the default zerolog logger with the
	// builtin logger implementation
	// We'd suggest not doing this though as the logs will look awful, you will see what we mean.
	mojito.Register(func() mojito.Logger {
		return mojito.NewBuiltinLogger()
	}, true)
}

var (
	Address = "0.0.0.0:8123"
	Logger  mojito.Logger
	Router  mojito.Router
)

func main() {
	Logger = mojito.DefaultLogger()
	Router = mojito.DefaultRouter()

	Logger.Info("Registering application routes")

	Router.WithMiddleware(mojito.RequestLogger)

	Router.GET("/assets/*path", mojito.AssetsHandler("/assets"))
	Router.GET("/", handlers.Home)

	Logger.Infof("Server has started on %s", Address)
	Logger.Error(Router.ListenAndServe(Address))
}
