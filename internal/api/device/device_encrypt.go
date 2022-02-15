package device

import (
	"Ys7/config"
	"Ys7/internal/api"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

func Encrypt(deviceSerial string) []byte {
	encryptUrl := config.Encrypt

	token, err := api.Token()
	if err != nil {
		log.Println(err)
	}

	params := url.Values{
		"accessToken": {token}, "deviceSerial": {deviceSerial},
	}

	resp, err := http.PostForm(encryptUrl, params)
	if err != nil {
		log.Println(err)
	}

	body, _ := ioutil.ReadAll(resp.Body)

	return body
}

func Decrypt(deviceSerial string, validateCode string) []byte {
	decryptUrl := config.Decrypt

	token, err := api.Token()
	if err != nil {
		log.Println(err)
	}

	params := url.Values{
		"accessToken": {token}, "deviceSerial": {deviceSerial}, "validateCode": {validateCode},
	}

	resp, err := http.PostForm(decryptUrl, params)
	if err != nil {
		log.Println(err)
	}

	body, _ := ioutil.ReadAll(resp.Body)

	return body
}
