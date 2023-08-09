package bootstrap

import (
	"context"
	"rnd-surajan-gin/api/controllers"
	"rnd-surajan-gin/api/routes"
	"rnd-surajan-gin/api/services"
	"rnd-surajan-gin/infrastucture"

	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
)

// All the necessary arguments for each constructor functions is provided using this code.
/* For Eg: For "NewTaskController" constructor func, it needs "services.TaskService",
which will be provided once we include "services.Module" in "fx.Options" below 👇. */
var Module = fx.Options(controllers.Module, services.Module, routes.Module, fx.Invoke(bootstrap))

func bootstrap(lifecycle fx.Lifecycle, taskRoutes routes.TaskRoutes) {
	lifecycle.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			// If we don't put all these inside a go routine, go will throw err: "context deadline exceeded", after server stays open for sometime.
			go func() {
				// Gin Server
				r := gin.Default()
				// Routes
				taskRoutes.TasksRouteSetup(r)
				// Listen and serve on "localhost:8080"
				r.Run(infrastucture.GetBaseUrl())
			}()

			return nil
		},
		OnStop: func(ctx context.Context) error {
			return nil
		},
	})
}
