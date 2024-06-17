package main

import (
	"net/http"
)

func (app *application) ShowHome(w http.ResponseWriter, r *http.Request) {
	app.render(w, "home.html", nil)
}

func (app *application) ShowAbout(w http.ResponseWriter, r *http.Request) {
	app.render(w, "about.html", nil)
}
