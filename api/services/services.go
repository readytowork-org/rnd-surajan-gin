package services

import "go.uber.org/fx"

var Module = fx.Options(
	fx.Provide(NewTaskService),
	fx.Provide(NewUserService),
	fx.Provide(NewJwtAuthService),
)
