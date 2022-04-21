package handlers

import (
	"Ys7/config"
	"Ys7/internal/api/device"
	"encoding/json"
	"net/http"
)

func DeviceCapture(w http.ResponseWriter, r *http.Request) {
	deviceSerial := r.FormValue("deviceSerial")

	data, code := device.GetDeviceCapture(deviceSerial)

	w.Header().Set("Content-Type", "application/json")
	if code != "200" {
		errJson := config.Error{
			Msg:  config.ErrorCode[code],
			Code: code,
		}
		message, _ := json.Marshal(errJson)

		w.Write(message)
		return
	}

	w.Write(data)
}
