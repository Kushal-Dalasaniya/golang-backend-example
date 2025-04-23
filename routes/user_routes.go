package routes

import (
	"github.com/Kushal-Dalasaniya/golang-backend/database"
	"github.com/Kushal-Dalasaniya/golang-backend/handlers"
	"github.com/Kushal-Dalasaniya/golang-backend/repositories"
	"github.com/Kushal-Dalasaniya/golang-backend/services"
	"github.com/gorilla/mux"
)

/* UserModule registers all user-related routes */
type UserModule struct {
	Handler handlers.UserHandler
}

/* NewUserModule initializes UserModule */
func NewUserModule() *UserModule {
	userRepo := repositories.NewUserRepository(database.DB)
	userService := services.NewUserService(userRepo)
	userHandler := handlers.NewUserHandler(userService)

	return &UserModule{Handler: userHandler}
}

func (m *UserModule) RegisterRoutes(router *mux.Router) {
	subRouter := router.PathPrefix("/api/users").Subrouter()
	subRouter.HandleFunc("", m.Handler.CreateUser).Methods("POST")
	subRouter.HandleFunc("", m.Handler.GetUsers).Methods("GET")
	subRouter.HandleFunc("/{id:[0-9]+}", m.Handler.GetUserByID).Methods("GET")
	subRouter.HandleFunc("/{id:[0-9]+}", m.Handler.UpdateUser).Methods("PUT")
	subRouter.HandleFunc("/{id:[0-9]+}", m.Handler.DeleteUser).Methods("DELETE")
}
