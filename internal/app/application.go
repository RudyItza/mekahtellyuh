package app

import (
	"fmt"
	"net/http"

	"github.com/RudyItza/mekatellyuh/internal/data"
)

// Application struct holds application dependencies like StoryModel
type Application struct {
	Stories *data.StoryModel
}

// ServerError handles internal server errors (500)
func (app *Application) ServerError(w http.ResponseWriter, err error) {
	http.Error(w, fmt.Sprintf("Internal Server Error: %v", err), http.StatusInternalServerError)
}

// ClientError handles client errors (e.g., 404 Not Found)
func (app *Application) ClientError(w http.ResponseWriter, statusCode int) {
	http.Error(w, http.StatusText(statusCode), statusCode)
}
