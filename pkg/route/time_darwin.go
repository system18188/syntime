// +build darwin

package route

import (
	"syscall"
	"time"
)

func SetLocalTime(t time.Time) error {
	tv := syscall.Timeval{
		Sec:  t.Unix(),
		Usec: 0,
	}
	return syscall.Settimeofday(&tv)
}
