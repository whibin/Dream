package log

import (
	"os"

	"github.com/sirupsen/logrus"
)

var Log *logrus.Logger

func init() {
	Log = logrus.New()
	// 设置输出样式，自带的只有两种样式logrus.JSONFormatter{}和logrus.TextFormatter{}
	Log.SetFormatter(&logrus.TextFormatter{})
	file, _ := os.OpenFile("./log/error.log", os.O_CREATE|os.O_APPEND, 0666)
	Log.SetOutput(file)
	Log.SetLevel(logrus.InfoLevel)
}
