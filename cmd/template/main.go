package main

import (
	"context"

	"Template/internal/app"
)

func main() {

	ctx := context.Background()
	err := app.Run(ctx)
	if err != nil {
		panic(err)
	}
}
