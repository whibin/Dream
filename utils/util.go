package utils

import (
	"Dream/conf"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

func LocalPathToUrl(path string, t int) string {
	if t == 1 {
		return conf.Config.Net.PrefixUrl + "/draw/" + path
	}
	if t == 2 {
		return conf.Config.Net.PrefixUrl + "/sound/" + path
	}
	return ""
}

func GetFirstDateOfMonth(d time.Time) time.Time {
	d = d.AddDate(0, 0, -d.Day()+1)
	return getZeroTime(d)
}

func GetLastDateOfMonth(d time.Time) time.Time {
	return GetFirstDateOfMonth(d).AddDate(0, 1, 0)
}

func getZeroTime(d time.Time) time.Time {
	return time.Date(d.Year(), d.Month(), d.Day(), 0, 0, 0, 0, d.Location())
}

type openIdStruct struct {
	Openid     string `json:"openid"`
	SessionKey string `json:"session_key"`
	Errcode    string `json:"errcode"`
}

func GetOpenId(appId, code, secret string) string {
	s := "https://api.weixin.qq.com/sns/jscode2session?appid=%s&secret=%s&js_code=%s&grant_type=authorization_code"
	url := fmt.Sprintf(s, appId, secret, code)
	resp, _ := http.Get(url)
	var bytes []byte
	resp.Body.Read(bytes)
	var openStruct openIdStruct
	json.Unmarshal(bytes, &openStruct)
	if openStruct.Errcode != "" {
		return openStruct.Openid
	}
	return ""
}
