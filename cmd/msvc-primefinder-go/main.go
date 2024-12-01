package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"

	"github.com/LewisT543/msvc-primefinder-go/internal"
)

func main() {
	app := internal.New()

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	err := app.Start(ctx)
	if err != nil {
		fmt.Println("failed to start app: ", err)
	}

}
