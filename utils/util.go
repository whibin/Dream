package utils

import (
	"Dream/conf"
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
