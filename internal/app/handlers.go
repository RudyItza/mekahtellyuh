package app

import (
	"github.com/RudyItza/mekatellyuh/internal/data"
	"net/http"
	"strconv"
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
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id <= 0 { // Ensure the ID is a valid positive integer
		app.ClientError(w, http.StatusNotFound)
		return
	}
	story, err := app.Stories.Get(id)
	if err != nil {
		app.ClientError(w, http.StatusNotFound)
		return
	}
	app.Render(w, r, "view_story.tmpl", map[string]any{"Story": story})
}

func (app *Application) SubmitStoryForm(w http.ResponseWriter, r *http.Request) {
	app.Render(w, r, "submit_story.tmpl", nil)
}

func (app *Application) SubmitStoryHandler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm() // Check if parsing the form succeeds
	if err != nil {
		app.ServerError(w, err)
		return
	}

	story := &data.Story{
		Title:    r.FormValue("title"),
		Content:  r.FormValue("content"),
		Username: r.FormValue("username"),
		Email:    r.FormValue("email"),
	}
	err = app.Stories.Insert(story)
	if err != nil {
		app.ServerError(w, err)
		return
	}
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func (app *Application) EditStoryForm(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id <= 0 {
		app.ClientError(w, http.StatusNotFound)
		return
	}
	story, err := app.Stories.Get(id)
	if err != nil {
		app.ClientError(w, http.StatusNotFound)
		return
	}
	app.Render(w, r, "edit_story.tmpl", map[string]any{"Story": story})
}

func (app *Application) EditStoryHandler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		app.ServerError(w, err)
		return
	}

	id, err := strconv.Atoi(r.FormValue("id"))
	if err != nil || id <= 0 {
		app.ClientError(w, http.StatusNotFound)
		return
	}
	story := &data.Story{
		ID:       id,
		Title:    r.FormValue("title"),
		Content:  r.FormValue("content"),
		Username: r.FormValue("username"),
		Email:    r.FormValue("email"),
	}
	err = app.Stories.Update(story)
	if err != nil {
		app.ServerError(w, err)
		return
	}
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
