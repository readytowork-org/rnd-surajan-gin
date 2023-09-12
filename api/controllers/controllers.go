package controllers

import "go.uber.org/fx"

var Module = fx.Options(
	fx.Provide(NewTaskController),
	fx.Provide(NewUserController),
)
