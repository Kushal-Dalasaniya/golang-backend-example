package container

import "github.com/gorilla/mux"

// Module interface ensures every module can register itself
type Module interface {
	RegisterRoutes(router *mux.Router);
}
