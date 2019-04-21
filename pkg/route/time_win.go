// +build windows

package route

import (
	"bufio"
	"fmt"
	"io"
	"os/exec"
	"time"

	"github.com/sirupsen/logrus"
)

func SetLocalTime(t time.Time) error {
	execCommand("cmd", []string{"/c", t.Format(`date 2006/01/02 & time 15:04:05`)})
	return nil
}

func execCommand(commandName string, params []string) bool {
	cmd := exec.Command(commandName, params...)
	//显示运行的命令
	fmt.Println(cmd.Args)
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		logrus.Error(err.Error())
		return false
	}
	cmd.Start()
	defer cmd.Wait()
	reader := bufio.NewReader(stdout)
	//实时循环读取输出流中的一行内容
	for {
		line, err2 := reader.ReadString('\n')
		if err2 != nil || io.EOF == err2 {
			break
		}
		logrus.Info(line)
	}
	return true
}
