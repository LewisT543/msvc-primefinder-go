package main

import (
	"context"
	"fmt"
	"github.com/LewisT543/msvc-primefinder-go/application/config"
	"os"
	"os/signal"

	"github.com/LewisT543/msvc-primefinder-go/application"
)

func main() {
	app := application.New(config.LoadConfig())

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	err := app.Start(ctx)
	if err != nil {
		fmt.Println("failed to start app: ", err)
	}

}
