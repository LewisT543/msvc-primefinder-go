package main

import (
	"context"
	"fmt"
	"github.com/LewisT543/msvc-primefinder-go/setup"
	"os"
	"os/signal"
)
 
func main() {
	conf, err := setup.LoadConfig()
	if err != nil {
		fmt.Println("failed to load config: ", err)
	}

	app := setup.New(conf)

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	err = app.Start(ctx)
	if err != nil {
		fmt.Println("failed to start app: ", err)
	}

}
