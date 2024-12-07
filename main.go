package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
)

func main() {
	conf, err := LoadConfig()
	if err != nil {
		fmt.Println("failed to load config: ", err)
	}

	app := New(conf)

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	err = app.Start(ctx)
	if err != nil {
		fmt.Println("failed to start app: ", err)
	}

}
