package bootstrap

import (
	"time"

	"github.com/Kamva/shark"
	"github.com/kataras/iris"
	"github.com/kataras/iris/middleware/logger"
	"github.com/kataras/iris/websocket"
)

// Configurator is just a function which accepts the framework instance.
// It is used for configuring application instance
type Configurator func(*Application)

// Application is responsible to manage the state of the application.
// It contains and handles all the necessary parts to create a fast web server.
type Application struct {
	*iris.Application
	AppSpawnDate time.Time
}

// SetupWebSockets will setup web socket on given endpoint
func (a *Application) SetupWebSockets(endpoint string, onConnection websocket.ConnectionFunc) {
	webSocket := websocket.New(websocket.Config{})
	webSocket.OnConnection(onConnection)

	a.Get(endpoint, webSocket.Handler())
	a.Any("/iris-webSocket.js", func(context iris.Context) {
		_, _ = context.Write(websocket.ClientSource)
	})
}

// SetupErrorHandlers will handle response for any kind of errors
// It uses `message` and `code` value that has set in context
func (a *Application) SetupErrorHandlers() {
	a.OnAnyErrorCode(func(context iris.Context) {
		_, _ = context.JSON(iris.Map{
			"message": context.Values().GetString("message"),
			"code":    context.Values().GetString("code"),
		})

		return
	})
}

// Configure runs all given configurators in a pipeline
func (a *Application) Configure(Configurators ...Configurator) {
	for _, configurator := range Configurators {
		configurator(a)
	}
}

// Bootstrap will bootstrap the base application instance
func (a *Application) Bootstrap() *Application {
	a.SetupErrorHandlers()

	a.Use(shark.ErrorRenderer())
	a.Use(shark.ErrorReporter())
	a.Use(logger.New())

	return a
}

// Listen will run the application on given address
func (a *Application) Listen(address string, configurators ...iris.Configurator) {
	_ = a.Run(iris.Addr(address), configurators...)
}

// New will return a new instance of Application
func New() *Application {
	return &Application{
		AppSpawnDate: time.Now(),
		Application:  iris.New(),
	}
}
