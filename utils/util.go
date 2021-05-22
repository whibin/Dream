package utils

import (
	"Dream/conf"
	. "Dream/database"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/garyburd/redigo/redis"
)

func LocalPathToUrl(path string, t int) string {
	if t == 1 {
		return conf.Config.Other.PrefixUrl + "/draw/" + path
	}
	if t == 2 {
		return conf.Config.Other.PrefixUrl + "/sound/" + path
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

func GetOpenId(appId, code, secret string) string {
	openid, _ := redis.String(RedisDB.Do("GET", code))
	if openid != "" {
		return openid
	}
	s := "https://api.weixin.qq.com/sns/jscode2session?appid=%s&secret=%s&js_code=%s&grant_type=authorization_code"
	url := fmt.Sprintf(s, appId, secret, code)
	fmt.Println(url)
	resp, _ := http.Get(url)
	var wxMap map[string]string
	json.NewDecoder(resp.Body).Decode(&wxMap)
	defer resp.Body.Close()
	newOpenId := wxMap["openid"]
	RedisDB.Do("SET", code, newOpenId)
	RedisDB.Do("EXPIRE", code, 5*60)
	return newOpenId
}
