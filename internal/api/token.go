package api

import (
	"Ys7/config"
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"time"

	"github.com/patrickmn/go-cache"
)

var CacheData = cache.New(0, 0)

func getToken() {
	var (
		appkeyUrl = "https://open.ys7.com/api/lapp/token/get"
		appkey    = ""
		secret    = ""
	)

	params := url.Values{"appKey": {appkey}, "appSecret": {secret}}
	resp, err := http.PostForm(appkeyUrl, params)
	if resp != nil {
		defer resp.Body.Close()
	}

	if err != nil {
		log.Println("萤石云 API 请求错误", err)
		return
	}

	tokenJson := config.TokenResp{}
	body, _ := ioutil.ReadAll(resp.Body)
	json.Unmarshal([]byte(body), &tokenJson)

	if tokenJson.Code != "200" {
		log.Println(tokenJson.Msg, tokenJson.Code)
		return
	}

	CacheData.Set("token", tokenJson.Data.AccessToken, time.Duration(tokenJson.Data.ExpireTime))
}

func Token() (string, error) {

	token, found := CacheData.Get("token")
	if found {
		log.Println("从缓存获取 token")
		return token.(string), nil
	}

	// 缓存中没有 token，获取 token
	getToken()
	token, found = CacheData.Get("token")
	if found {
		log.Println("更新 token 并从缓存获取 token")
		return token.(string), nil
	}

	return "", errors.New("获取 token 失败")
}
