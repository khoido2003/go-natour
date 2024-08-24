package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/khoido2003/go-natour/controller"
	"github.com/khoido2003/go-natour/database"
	router "github.com/khoido2003/go-natour/routes"
	_ "github.com/lib/pq"
)

func main() {

	// Load environment variables from.env file
	godotenv.Load()

	///////////////////////////////////////

	// Connect to PostgreSQL database

	dbURL := os.Getenv("DB_URL")

	if dbURL == "" {
		log.Fatal("DB_URL is not found")
	}

	conn, err := sql.Open("postgres", dbURL)

	if err != nil {
		log.Fatal("Can't connect to db: ", err)
	}

	queries := database.New(conn)

	apiConfig := controller.APIConfig{
		DB: queries,
	}

	//////////////////////////////////////////

	// Initialize chi router

	router := router.RouterController(apiConfig)

	//////////////////////////////////////////

	portStr := os.Getenv("PORT")

	if portStr == "" {
		log.Fatal("PORT is not found")
	}

	srv := &http.Server{
		Handler: router,
		Addr:    ":" + portStr,
	}

	log.Printf("Server starting on port %v", portStr)

	err = srv.ListenAndServe()

	if err != nil {
		log.Fatal(err)
	}
}
