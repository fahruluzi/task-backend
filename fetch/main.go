package main

import (
	"context"
	"fetch/application"
	"fetch/application/controller"
	"fetch/middleware"
	"fmt"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"log"
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	// create a context
	ctx, cancel := context.WithCancel(context.Background())
	app, err := application.NewApp(ctx)
	if err != nil {
		log.Print("Failed to initialize app. Error: ", err)
		panic(err)
	}

	defer app.Close()

	// run a go routine in which the cancel() function is called once an OS interrupt is received
	go func() {
		data := <-app.TerminalHandler

		fmt.Printf("\nsystem call: %+v\n", data)
		_ = app.Fiber.Shutdown()
		cancel()
	}()

	app.Fiber.Use(logger.New())

	// Serve Swagger
	app.Fiber.Static("/swagger", "./")

	jwtMiddleware := middleware.NewJWTMiddleware(app.Config.JWTSecret)
	app.Fiber.Use(jwtMiddleware.Validate)

	validateController := controller.NewValidateController()
	app.Fiber.Get("/validate", validateController.Validate)

	fetchController := controller.NewFetchController()
	app.Fiber.Get("/fetch", fetchController.Fetch)

	aggregationController := controller.NewAggregationController()
	app.Fiber.Get("/aggregation", jwtMiddleware.IsAdmin, aggregationController.Aggregation)

	log.Fatal(app.Fiber.Listen(app.Config.Port))

	<-ctx.Done()
}
