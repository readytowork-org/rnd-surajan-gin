package main

import (
	"rnd-surajan-gin/bootstrap"
	"rnd-surajan-gin/environment"

	"go.uber.org/fx"
)

func init() {
	// Initialize Env
	environment.EnvInit()
}

func main() {
	fx.New(bootstrap.Module).Run()
}
