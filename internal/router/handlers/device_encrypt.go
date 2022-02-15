package handlers

import (
	"Ys7/internal/api/device"
	"net/http"
)

func Encrypt(w http.ResponseWriter, r *http.Request) {
	deviceSerial := r.FormValue("deviceSerial")

	w.Header().Set("Content-Type", "application/json")
	w.Write(device.Encrypt(deviceSerial))
}

func Decrypt(w http.ResponseWriter, r *http.Request) {
	var (
		deviceSerial = r.FormValue("deviceSerial")
		validateCode = r.FormValue("validateCode")
	)

	w.Header().Set("Content-Type", "application/json")
	w.Write(device.Decrypt(deviceSerial, validateCode))
}
