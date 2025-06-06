package container

import (
	"github.com/Kushal-Dalasaniya/golang-backend/routes"
	"github.com/gorilla/mux"
)

type Container struct {
	Modules []Module // List of modules (User, Auth, etc.)
}

/* RegisterModule adds a new module dynamically */
func (c *Container) RegisterModule(m Module) {
	c.Modules = append(c.Modules, m)
}

/* NewContainer initializes dependencies dynamically */
func NewContainer(router *mux.Router) *Container {

	/* Initialize Container */
	c := &Container{}

	/* Register all modules (Users, Auth, etc.) */
	addModules(c)

	/* Register routes for each module */
	for _, module := range c.Modules {
		module.RegisterRoutes(router)
	}

	return c
}

/* Register all modules (Users, Auth, etc.) */
func addModules(c *Container) {
	c.RegisterModule(routes.NewUserModule())
}
