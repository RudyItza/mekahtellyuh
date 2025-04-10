package main

import (
	"log"
	"net/http"

	"github.com/RudyItza/mekatellyuh/internal/app"
	"github.com/RudyItza/mekatellyuh/internal/data"
	"github.com/RudyItza/mekatellyuh/internal/db"
)

func main() {
	// Open the DB connection using the db.InitDB function
	dbConn, err := db.InitDB() // db is now being used correctly
	if err != nil {
		log.Fatalf("Failed to open database: %v", err)
	}
	defer func() {
		if err := dbConn.Close(); err != nil {
			log.Printf("Error closing database: %v", err)
		}
	}()

	// Initialize the application with the database connection
	application := &app.Application{
		Stories: &data.StoryModel{DB: dbConn}, // Pass the DB connection to the StoryModel
	}

	// Start the HTTP server on port 4000
	log.Println("Server starting on :4000...")
	err = http.ListenAndServe(":4000", application.Routes()) // Use the Routes method from the app
	if err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
