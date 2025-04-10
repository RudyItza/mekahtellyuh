package main

import (
	"log"
	"net/http"

	"github.com/RudyItza/mekatellyuh/internal/app"
	"github.com/RudyItza/mekatellyuh/internal/data"
	"github.com/RudyItza/mekatellyuh/internal/db"
)

func main() {
	dbConn, err := db.OpenDB()
	if err != nil {
		log.Fatal(err)
	}
	defer dbConn.Close()

	application := &app.Application{
		Stories: &data.StoryModel{DB: dbConn},
	}

	log.Println("Server starting on :4000...")
	err = http.ListenAndServe(":4000", application.Routes())
	log.Fatal(err)
}
