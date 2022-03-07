package main

import (
	"context"
	"fetch/application"
	"fmt"
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

	// Serve Swagger
	app.Fiber.Static("/swagger", "./")

	log.Fatal(app.Fiber.Listen(app.Config.Port))

	<-ctx.Done()
}
