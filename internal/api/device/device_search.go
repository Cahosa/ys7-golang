package device

import (
	"Ys7/config"
	"encoding/json"
	"strings"
)

func DeviceSearch(key string) ([]byte, string) {

	if key == "" {
		return []byte("key is empty"), "failed"
	}

	var (
		pageStart = 0
		pageSize  = 50
		total     = 1
	)
	cameraListInfo := config.DeviceList{}

	data := []interface{}{}

	for total != 0 {
		rsp, code := GetCameraList(pageStart, pageSize)

		if code != "200" {
			return nil, code
		}

		json.Unmarshal(rsp, &cameraListInfo)

		for _, item := range cameraListInfo.Data {
			name := item.DeviceName
      serial := item.DeviceSerial

			if strings.Contains(name, key) || strings.Contains(serial, key) {
				data = append(data, item)
			}
		}

		total = cameraListInfo.Page.Total
		pageStart++
	}

	if len(data) == 0 {
		return []byte("No camera found"), "failed"
	}

	jsonData, _ := json.Marshal(data)
	return jsonData, "success"
}
