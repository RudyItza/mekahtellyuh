package app

import (
	"net/http"
	"strconv"

	"github.com/RudyItza/mekatellyuh/internal/data"
)

func (app *Application) HomeHandler(w http.ResponseWriter, r *http.Request) {
	stories, err := app.Stories.GetAll()
	if err != nil {
		app.ServerError(w, err)
		return
	}
	app.Render(w, r, "home.tmpl", map[string]any{"Stories": stories})
}

func (app *Application) ViewStoryHandler(w http.ResponseWriter, r *http.Request) {
	// Fetch all stories from your storage (database, etc.)
	stories, err := app.Stories.GetAll() // Assuming you have a GetAll function in your Stories model
	if err != nil {
		app.ClientError(w, http.StatusInternalServerError)
		return
	}

	// Render the template and pass the stories data
	app.Render(w, r, "view_story.tmpl", map[string]any{"Stories": stories})
}

func (app *Application) SubmitStoryForm(w http.ResponseWriter, r *http.Request) {
	// Just render the form with no data initially
	app.Render(w, r, "submit_story.tmpl", nil)
}
func (app *Application) SubmitStoryHandler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		app.ServerError(w, err)
		return
	}

	title := r.FormValue("title")
	content := r.FormValue("content")
	username := r.FormValue("username")
	email := r.FormValue("email")

	// Validate inputs
	validationErrors := ValidateInputs(title, content, username, email)

	// If there are validation errors, render the form again with errors and previously entered values
	if len(validationErrors) > 0 {
		app.Render(w, r, "submit_story.tmpl", map[string]any{
			"Errors":   validationErrors,
			"Title":    title,
			"Content":  content,
			"Username": username,
			"Email":    email,
		})
		return
	}

	// Create the story and insert it
	story := &data.Story{
		Title:    title,
		Content:  content,
		Username: username,
		Email:    email,
	}
	err = app.Stories.Insert(story)
	if err != nil {
		app.ServerError(w, err)
		return
	}

	// Redirect to the homepage after a successful submission
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

// EditStoryForm renders the form to edit a story
func (app *Application) EditStoryForm(w http.ResponseWriter, r *http.Request) {
	// Extract story ID from query parameters
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id <= 0 {
		app.ClientError(w, http.StatusNotFound)
		return
	}

	// Retrieve the story from the database
	story, err := app.Stories.Get(id)
	if err != nil {
		app.ClientError(w, http.StatusNotFound)
		return
	}

	// Render the edit form with the story data
	app.Render(w, r, "edit_story.tmpl", map[string]any{
		"Story": story, // Pass the story data for pre-filling the form
	})
}

// EditStoryHandler handles the form submission and updates the story
func (app *Application) EditStoryHandler(w http.ResponseWriter, r *http.Request) {
	// Parse the form data
	err := r.ParseForm()
	if err != nil {
		app.ServerError(w, err)
		return
	}

	// Extract the story ID from the form data
	id, err := strconv.Atoi(r.FormValue("id"))
	if err != nil || id <= 0 {
		app.ClientError(w, http.StatusNotFound)
		return
	}

	// Create a story object from the form data
	story := &data.Story{
		ID:       id,
		Title:    r.FormValue("title"),
		Content:  r.FormValue("content"),
		Username: r.FormValue("username"),
		Email:    r.FormValue("email"),
	}

	// Validate the inputs
	errors := ValidateInputs(story.Title, story.Content, story.Username, story.Email)

	// If validation errors exist, re-render the form with errors
	if len(errors) > 0 {
		// Render the edit form with validation errors
		app.Render(w, r, "edit_story.tmpl", map[string]any{
			"Story":  story,  // Pass the current story back to the form
			"Errors": errors, // Pass the validation errors to display
		})
		return
	}

	// Update the story in the database
	err = app.Stories.Update(story)
	if err != nil {
		app.ServerError(w, err)
		return
	}

	// Redirect to the homepage or success page after update
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func (app *Application) DeleteStoryHandler(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id <= 0 {
		app.ClientError(w, http.StatusNotFound)
		return
	}
	err = app.Stories.Delete(id)
	if err != nil {
		app.ServerError(w, err)
		return
	}
	http.Redirect(w, r, "/", http.StatusSeeOther)
}
