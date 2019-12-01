# 时间同步

## 使用说明 
以管理员权限运行
sudo syntime -help

```
Usage of syntime:
  -e int
        定时更新间隔，单位分 (default 60)
  -key string
        Response Headers 时间键名 (default "Date")
  -layout string
        时间格式 (default "Mon, 02 Jan 2006 15:04:05 MST")
  -level string
        日记等级 可选值:panic,fatal,error,warn,warning,info,debug (default "debug")
  -out duration
        远程连接超时时间，单位毫秒 (default 10s)
  -url string
        请求地址 (default "http://www.baidu.com/")
  -utc string
        时区 (default "MST")

```