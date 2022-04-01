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

type CameraList struct {
	Msg  string `json:"msg"`
	Code string `json:"code"`
	Data []struct {
		ID           string `json:"id"`
		DeviceSerial string `json:"deviceSerial"`
		ChannelNo    int    `json:"channelNo"`
		ChannelName  string `json:"channelName"`
		Status       int    `json:"status"`
		IsShared     string `json:"isShared"`
		PicURL       string `json:"picUrl"`
		IsEncrypt    int    `json:"isEncrypt"`
		VideoLevel   int    `json:"videoLevel"`
		Permission   int    `json:"permission"`
		IsAdd        int    `json:"isAdd"`
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
