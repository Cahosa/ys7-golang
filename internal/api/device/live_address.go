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

func GetLiveAddress(params url.Values) (string, string) {
	liveUrl := config.Live

	token, err := api.Token()
	if err != nil {
		log.Println(err)
	}

	params.Add("accessToken", token)

	resp, err := http.PostForm(liveUrl, params)
	if err != nil {
		log.Println(err)
	}

	body, _ := ioutil.ReadAll(resp.Body)
	liveJson := config.LiveAddress{}
	json.Unmarshal(body, &liveJson)

	return liveJson.Data.URL, liveJson.Code
}
