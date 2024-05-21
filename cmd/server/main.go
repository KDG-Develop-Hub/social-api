package main

import (
	"github.com/kdg-develop-hub/api/config"
	. "github.com/kdg-develop-hub/api/internal/app"
)

func main() {
	c := config.MustLoad()
	app := NewApp(c)
	app.Run()
}
