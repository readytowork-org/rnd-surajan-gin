package routes

import "go.uber.org/fx"

var Module = fx.Options(fx.Provide(NewRoutes), fx.Provide(NewTaskRoutes), fx.Provide(NewTestRoutes))

// Custom Type which is a slice "[]" of "Route" type.
type Routes []Route

type Route interface {
	Setup()
}

// Here, "TaskRoutes" is under the same package name "routes", so we cannot import it like: "routes.TaskRoutes".
// It is automatically available as "TaskRoutes".
func NewRoutes(taskRoutes TaskRoutes, testRoutes TestRoutes) Routes {
	return Routes{taskRoutes, testRoutes}
}

func (r Routes) Setup() {
	for _, route := range r {
		route.Setup()
	}
}
