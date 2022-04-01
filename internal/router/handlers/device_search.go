package handlers

import (
	"Ys7/config"
	"Ys7/internal/api/device"
	"net/http"
)

func DeviceSearch(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	key := r.Form.Get("key")
	data, msg := device.DeviceSearch(key)

	if msg == "failed" {
		ErrorHandler(w, string(data), "400")
		return
	} else if msg != "200" && msg != "success" {
		ErrorHandler(w, config.ErrorCode[msg], "400")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(data)
}
