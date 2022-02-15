package config

type Error struct {
	Msg  string `json:"message"`
	Code string `json:"code"`
}

type TokenResp struct {
	Msg  string `json:"msg"`
	Code string `json:"code"`
	Data struct {
		AccessToken string `json:"accessToken"`
		ExpireTime  int64  `json:"expireTime"`
	} `json:"data"`
}

type DeviceList struct {
	Msg  string `json:"msg"`
	Code string `json:"code"`
	Data []struct {
		DeviceSerial  string `json:"deviceSerial"`
		DeviceName    string `json:"deviceName"`
		DeviceType    string `json:"deviceType"`
		Status        int    `json:"status"`
		DeviceVersion string `json:"deviceVersion"`
		AddTime       int64  `json:"addTime"`
	} `json:"data"`
	Page struct {
		Total int `json:"total"`
		Size  int `json:"size"`
		Page  int `json:"page"`
	} `json:"page"`
}

type DeviceInfo struct {
	Msg  string `json:"msg"`
	Code string `json:"code"`
	Data struct {
		DeviceSerial string `json:"deviceSerial"`
		DeviceName   string `json:"deviceName"`
		Model        string `json:"model"`
		Status       int    `json:"status"`
		IsEncrypt    int    `json:"isEncrypt"`
	} `json:"data"`
}

type LiveAddress struct {
	Msg  string `json:"msg"`
	Code string `json:"code"`
	Data struct {
		ID         string `json:"id"`
		URL        string `json:"url"`
		ExpireTime string `json:"expireTime"`
	} `json:"data"`
}
