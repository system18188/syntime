package config

import "time"

var Url string            // 请求地址
var HeadersKey string     // 请求头键名
var Timeout time.Duration // 远程连接超时间
var Location string       // 时区
var Layout string         // 时间格式
var Every int             // 定时运行间隔
var Level string          // 日记等级
