package main

import (
	"flag"
	"github.com/sirupsen/logrus"
	"runtime"
	"syntime/pkg/config"
	"syntime/pkg/route"
	"syntime/pkg/scheduler"

	"time"
)


func init(){
	flag.IntVar(&config.Every,"e",60,"定时更新间隔，单位分")
	flag.StringVar(&config.Url,"url","http://www.baidu.com/","请求地址")
	flag.StringVar(&config.HeadersKey,"key","Date","Response Headers 时间键名")
	flag.StringVar(&config.Location,"utc","MST","时区")
	flag.StringVar(&config.Layout,"layout",time.RFC1123,"时间格式")
	flag.DurationVar(&config.Timeout,"out",10 * time.Second,"远程连接超时时间，单位毫秒")
	flag.StringVar(&config.Level,"level","debug","日记等级 可选值:panic,fatal,error,warn,warning,info,debug")
}

func main()  {
	flag.Parse()

	logrus.SetFormatter(&logrus.TextFormatter{})
	level, err := logrus.ParseLevel(config.Level)
	if err != nil {
		logrus.Fatal(err.Error())
	}
	logrus.SetLevel(level)

	scheduler.Every(config.Every).Minutes().Run(route.Run)

	logrus.Info("启动")
	// 保持程序不退出。
	runtime.Goexit()
}
