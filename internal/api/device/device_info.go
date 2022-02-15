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

func GetDeviceInfo(deviceSerial string) ([]byte, string) {
	infoUrl := config.Info

	token, err := api.Token()
	if err != nil {
		log.Println(err)
	}

	params := url.Values{
		"accessToken": {token}, "deviceSerial": {deviceSerial},
	}

	resp, err := http.PostForm(infoUrl, params)
	if err != nil {
		log.Println(err)
	}

	body, _ := ioutil.ReadAll(resp.Body)
	infoJson := config.DeviceInfo{}
	json.Unmarshal(body, &infoJson)
	data, _ := json.Marshal(infoJson.Data)

	return data, infoJson.Code
}
