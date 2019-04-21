package route

import (
	"net/http"
	"syntime/pkg/config"
	"time"

	"github.com/sirupsen/logrus"
)

func Run() {

	defer func() {
		if err := recover(); err != nil {
			logrus.Error(err)
		}
	}()

	client := http.Client{}
	client.Timeout = config.Timeout
	resp, err := client.Head(config.Url)
	if err != nil {
		logrus.Error(err.Error())
		return
	}

	if resp.StatusCode != http.StatusOK {
		logrus.Error("http 响应代码：", resp.StatusCode)
		return
	}

	loc, err := time.LoadLocation(config.Location)
	if err != nil {
		logrus.Error(err.Error())
		return
	}

	t, err := time.ParseInLocation(config.Layout, resp.Header.Get(config.HeadersKey), loc)
	if err != nil {
		logrus.Error(err.Error())
		return
	}

	t.In(loc)

	if err := SetLocalTime(t); err != nil {
		logrus.Error("设置时间出错", err.Error())
		return
	}

	logrus.Info("时区：", config.Location, " 时间：", t.Format(`2006/01/02 15:04:05`), " Unix：", t.Unix(), " 地址：", config.Url)
}
