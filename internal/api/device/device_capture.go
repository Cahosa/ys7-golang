package device

import (
	"Ys7/config"
	"Ys7/internal/api"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

func GetDeviceCapture(deviceSerial string) ([]byte, string) {
	captureUrl := config.Capture

	token, err := api.Token()
	if err != nil {
		log.Println(err)
	}

	parmas := url.Values{"accessToken": {token}, "deviceSerial": {deviceSerial}}

	resp, err := http.PostForm(captureUrl, parmas)

	if err != nil {
		log.Println(err)
	}

	body, _ := ioutil.ReadAll(resp.Body)
	captureData := config.DeviceCapture{}
	json.Unmarshal(body, &captureData)
	data, _ := json.Marshal(captureData.Data)

	return data, captureData.Code
}
