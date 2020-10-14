package utils

import (
	"os"
	"syscall"
	"time"
)

func FileCreationTime(finfo os.FileInfo) time.Time {
	statT := finfo.Sys().(*syscall.Stat_t)
	return timespecToTime(statT.Ctim)
}

func timespecToTime(ts syscall.Timespec) time.Time {
	return time.Unix(int64(ts.Sec), int64(ts.Nsec))
}
