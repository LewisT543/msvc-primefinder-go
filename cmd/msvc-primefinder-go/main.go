package main

import (
	"context"
	"fmt"
	"github.com/LewisT543/msvc-primefinder-go/application"
	"github.com/LewisT543/msvc-primefinder-go/application/config"
	"os"
	"os/signal"
)

func main() {
	c, err := config.LoadConfig()
	if err != nil {
		fmt.Println("failed to load config: ", err)
	}

	app := application.New(c)

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	err = app.Start(ctx)
	if err != nil {
		fmt.Println("failed to start app: ", err)
	}

}
