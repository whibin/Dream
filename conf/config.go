package conf

import (
	"github.com/BurntSushi/toml"
	"io/ioutil"
	"log"
)

type Configuration struct {
	DB  DataBase `toml:"database"`
	Net Http     `toml:"http"`
	O   Other    `toml:"other"`
}

type DataBase struct {
	IP       string `toml:"ip"`
	Port     string `toml:"port"`
	Username string `toml:"username"`
	Password string `toml:"password"`
	Database string `toml:"database"`
	Conf     string `toml:"conf"`
}

type Http struct {
	Port      string `toml:"port"`
	PrefixUrl string `toml:"prefix-url"`
}

type Other struct {
	LocalPathPrefix string `toml:"local-path-prefix"`
}

var Config Configuration

func init() {
	var (
		fcontent []byte
		err      error
	)
	Config = *new(Configuration)
	// 在go语言中./表示当前工作目录，例如当前调用读取操作的文件为Project/xxx/config.go，则./表示Project/
	if fcontent, err = ioutil.ReadFile("./conf/conf.toml"); err != nil {
		log.Fatal("open error ", err)
		return
	}
	if err = toml.Unmarshal(fcontent, &Config); err != nil {
		log.Fatal("toml.Unmarshal error ", err)
		return
	}
	return
}
