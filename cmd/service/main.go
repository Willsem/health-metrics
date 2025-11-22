package main

import (
	"health-metrics/internal/infra/startup"
)

func main() {
	app := startup.NewApp()
	app.Run()
}
