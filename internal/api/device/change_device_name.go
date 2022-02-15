package device

import (
	"Ys7/config"
	"Ys7/internal/api"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

func ChangeDeviceName(deviceSerial, deviceName string) []byte {
	changeUrl := config.ChangeDeviceName

	token, err := api.Token()
	if err != nil {
		log.Println(err)
	}

	params := url.Values{
		"accessToken":  {token},
		"deviceSerial": {deviceSerial},
		"deviceName":   {deviceName},
	}

	resp, err := http.PostForm(changeUrl, params)
	if err != nil {
		log.Println(err)
	}

	body, _ := ioutil.ReadAll(resp.Body)

	return body
}
