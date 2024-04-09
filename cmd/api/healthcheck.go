package main

import (
	"net/http"
)

func (app *application) healthcheck(w http.ResponseWriter, r *http.Request) {
	jresp := envelop{
		"status": "available",
		"system_info": map[string]string{
			"environment": app.config.env,
			"version":     version,
		},
	}
	err := app.writeJSON(w, http.StatusOK, jresp, nil)
	if err != nil {
		app.logger.Println(err)
		app.serverorErrorResponse(w, r, err)
	}
}
