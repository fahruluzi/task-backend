package application

import (
	"context"
	"fetch/config"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"os"
	"os/signal"
	"syscall"
)

type App struct {
	TerminalHandler chan os.Signal
	Context         context.Context
	Config          *config.Configuration
	Fiber           *fiber.App
}

const (
	name = "fetch-app"
)

// NewApp is a function to create application instance
func NewApp(ctx context.Context) (*App, error) {
	cfgLoaded := config.LoadConfiguration(name)

	// create a channel for listening to OS signals and connecting OS interrupts to the channel
	terminalHandler := make(chan os.Signal, 1)
	signal.Notify(terminalHandler,
		os.Interrupt,
		syscall.SIGHUP,
		syscall.SIGTERM,
		syscall.SIGQUIT,
	)

	httpServer := fiber.New(fiber.Config{
		CaseSensitive: true,
		StrictRouting: true,
		ServerHeader:  "Microservice",
		AppName:       "Fetch App",
	})

	app := &App{
		Context:         ctx,
		Config:          cfgLoaded,
		TerminalHandler: terminalHandler,
		Fiber:           httpServer,
	}

	return app, nil
}

// Close is a function to gracefully close the application
func (app *App) Close() {
	fmt.Println("APP SUCCESSFULLY CLOSED")
}
