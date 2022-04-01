package handlers

import (
	"Ys7/config"
	"Ys7/internal/api/device"
	"net/http"
	"strconv"
)

func DeviceList(w http.ResponseWriter, r *http.Request) {
	var (
		pageStart, _ = strconv.Atoi(r.FormValue("pageStart"))
		pageSize, _  = strconv.Atoi(r.FormValue("pageSize"))
	)

	if pageSize == 0 {
		pageSize = 10
	}

	data, code := device.GetCameraList(pageStart, pageSize)
	if code != "200" {
		ErrorHandler(w, config.ErrorCode[code], "400")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(data)
}
