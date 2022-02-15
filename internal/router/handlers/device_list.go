package handlers

import (
	"Ys7/config"
	"Ys7/internal/api/device"
	"net/http"
)

func DeviceList(w http.ResponseWriter, r *http.Request) {
	var (
		pageStart = r.FormValue("pageStart")
		pageSize  = r.FormValue("pageSize")
	)

	data, code := device.GetDeviceList(pageStart, pageSize)
	if code != "200" {
		ErrorHandler(w, config.ErrorCode[code], "400")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(data)
}
