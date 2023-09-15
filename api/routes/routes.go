package routes

import "go.uber.org/fx"

var Module = fx.Options(fx.Provide(NewRoutes), fx.Provide(NewTaskRoutes), fx.Provide(NewUserRoutes), fx.Provide(NewJwtAuthRoutes), fx.Provide(NewTestRoutes))

// Custom Type which is a slice "[]" of "Route" type.
type Routes []Route

type Route interface {
	Setup()
}

// Here, "TaskRoutes" is under the same package name "routes", so we cannot import it like: "routes.TaskRoutes".
// It is automatically available as "TaskRoutes".
func NewRoutes(taskRoutes TaskRoutes, userRoutes UserRoutes, jwtAuthRoutes JwtAuthRoutes, testRoutes TestRoutes) Routes {
	return Routes{taskRoutes, userRoutes, jwtAuthRoutes, testRoutes}
}

func (r Routes) Setup() {
	for _, route := range r {
		route.Setup()
	}
}
