package utils

import (
	"os"
	"syscall"
	"time"
)

func FileCreationTime(finfo os.FileInfo) time.Time {
	return timespecToTime(finfo.Sys().(*syscall.Stat_t).Ctimespec)
}

func timespecToTime(ts syscall.Timespec) time.Time {
	return time.Unix(int64(ts.Sec), int64(ts.Nsec))
}
