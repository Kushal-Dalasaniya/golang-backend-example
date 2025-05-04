package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/Kushal-Dalasaniya/golang-backend/container"
	"github.com/Kushal-Dalasaniya/golang-backend/database"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Println("No .env file found")
	}

	/* Connect Database */
	database.ConnectDB()

	/* Initialize Router */
	router := mux.NewRouter()

	router.Use(middleware)

	/* Initialize Container (Auto-loads all modules) */
	_ = container.NewContainer(router)

	/* Start Server */
	port := os.Getenv("PORT")
	fmt.Println("🚀 Server running on port:", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}

func middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		println(r.Header.Get("Authorization"))
		w.Header().Set("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}