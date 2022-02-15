package handlers

import (
	"Ys7/internal/api/device"
	"net/http"
)

func ChangeName(w http.ResponseWriter, r *http.Request) {
	var (
		deviceSerial = r.FormValue("deviceSerial")
		deviceName   = r.FormValue("deviceName")
	)

	w.Header().Set("Content-Type", "application/json")
	w.Write(device.ChangeDeviceName(deviceSerial, deviceName))
}
