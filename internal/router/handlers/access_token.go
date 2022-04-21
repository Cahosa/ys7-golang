package handlers

import (
	"Ys7/internal/api"
	"net/http"
)

func AccessToken(w http.ResponseWriter, r *http.Request) {
	token, err := api.Token()

	if err != nil {
		ErrorHandler(w, err.Error(), "500")
		return
	}

	w.Header().Set("Content-Type", "text/plain")
	w.Write([]byte(token))
}
