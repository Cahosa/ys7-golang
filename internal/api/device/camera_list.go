package device

import (
	"Ys7/config"
	"Ys7/internal/api"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strconv"
)

func GetCameraList(pageStart, pageSize int) ([]byte, string) {
	listUrl := config.List

	token, err := api.Token()
	if err != nil {
		log.Println(err)
	}

	params := url.Values{
		"accessToken": {token}, "pageStart": {strconv.Itoa(pageStart)}, "pageSize": {strconv.Itoa(pageSize)},
	}

	resp, err := http.PostForm(listUrl, params)
	if err != nil {
		log.Println(err)
	}

	body, _ := ioutil.ReadAll(resp.Body)
	listJson := config.CameraList{}
	json.Unmarshal(body, &listJson)
	data, _ := json.Marshal(listJson)

	return data, listJson.Code
}
