package utils

import (
	"Dream/conf"
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
