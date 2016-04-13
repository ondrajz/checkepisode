package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func NewAPI(app *App) *mux.Router {
	var r = mux.NewRouter()

	r.Path("/api/search").Methods("GET").
		HandlerFunc(app.searchHandler)

	return r
}

func (a *App) searchHandler(w http.ResponseWriter, r *http.Request) {
	var q = r.FormValue("q")
	if len(q) == 0 {
		http.Error(w, "search query empty", http.StatusBadRequest)
		return
	}

	log.Printf("searching '%s'", q)
	shows, err := SearchShows(q)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	b, err := json.MarshalIndent(shows, "", " ")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Fprint(w, string(b))
}
