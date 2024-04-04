package main

import (
	"fmt"
	"net/http"
	"strconv"
)

func (app *application) createMovieHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "creating a movie....")
}

func (app *application) showMovieHandler(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseInt(r.PathValue("id"), 10, 64)
	if id < 1 || err != nil {
		http.NotFound(w, r)
		return
	}
	fmt.Fprintf(w, "showing the details of a movie with id %d...\n", id)
}
