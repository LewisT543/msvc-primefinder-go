package main

import (
	"context"
	"fmt"
	"github.com/LewisT543/msvc-primefinder-go/internal"
)

func main() {
	app := internal.New()

	err := app.Start(context.TODO())
	if err != nil {
		fmt.Println("failed to start app: ", err)
	}
}
