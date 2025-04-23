package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Kushal-Dalasaniya/golang-backend/container"
	"github.com/Kushal-Dalasaniya/golang-backend/database"
	"github.com/gorilla/mux"
)

func main() {
	/* Connect Database */
	database.ConnectDB()

	/* Initialize Router */
	router := mux.NewRouter()

	/* Initialize Container (Auto-loads all modules) */
	_ = container.NewContainer(router)

	/* Start Server */
	port := "8080"
	fmt.Println("🚀 Server running on port:", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
