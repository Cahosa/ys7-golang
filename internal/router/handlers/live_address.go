package handlers

import (
	"Ys7/config"
	"Ys7/internal/api/device"
	"net/http"
)

func LiveAddress(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	address, code := device.GetLiveAddress(r.Form)

	if code != "200" {
		ErrorHandler(w, config.ErrorCode[code], "400")
		return
	}

	w.Write([]byte(address))
}
