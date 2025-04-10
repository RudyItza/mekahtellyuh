package app

import "net/http"

func (app *Application) Routes() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/", app.HomeHandler)
	mux.HandleFunc("/story/view", app.ViewStoryHandler)
	mux.HandleFunc("/story/submit", app.SubmitStoryForm)    // GET request for the form
	mux.HandleFunc("/story/create", app.SubmitStoryHandler) // POST request to submit a new story
	mux.HandleFunc("/story/edit", app.EditStoryForm)        // GET request for the edit form
	mux.HandleFunc("/story/update", app.EditStoryHandler)   // POST request to update the story
	mux.HandleFunc("/story/delete", app.DeleteStoryHandler)
	return app.RecoverPanic(app.LogRequest(mux)) // Apply logging and panic recovery middleware
}
