package handlers

import (
	"Ys7/config"
	"Ys7/internal/api/device"
	"net/http"
)

func DeviceInfo(w http.ResponseWriter, r *http.Request) {
	deviceSerial := r.FormValue("deviceSerial")
	data, code := device.GetDeviceInfo(deviceSerial)
	if code != "200" {
		ErrorHandler(w, config.ErrorCode[code], "400")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(data)
}
