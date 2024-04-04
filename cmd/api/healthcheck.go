package main

import (
	"fmt"
	"net/http"
)

func (app *application) healthcheck(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "The server is running on %s mode\n", app.config.env)
	fmt.Fprintf(w, "Version: %s\n", version)
}
