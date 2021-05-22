package conf

import (
	"io/ioutil"
	"log"

	"github.com/BurntSushi/toml"
)

type Configuration struct {
	DB    MySQL `toml:"mysql"`
	Net   Http  `toml:"http"`
	Other Other `toml:"other"`
	Redis Redis `toml:"redis"`
}

type MySQL struct {
	IP        string `toml:"ip"`
	Port      string `toml:"port"`
	Username  string `toml:"username"`
	Password  string `toml:"password"`
	Database  string `toml:"database"`
	Variables string `toml:"variables"`
}

type Redis struct {
	IP       string `toml:"ip"`
	Port     string `toml:"port"`
	Password string `toml:"password"`
}

type Http struct {
	Port string `toml:"web-port"`
}

type Other struct {
	PrefixUrl       string `toml:"prefix-url"`
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
}
