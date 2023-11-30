package main

import (
	"Ozon/internal/app"
	"context"
)

func main() {
	a := app.New(context.Background())
	app.Run(a)
}
