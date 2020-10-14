package utils

import (
	"os"
	"syscall"
	"time"
)

func FileCreationTime(finfo os.FileInfo) time.Time {
	return time.Unix(0, finfo.Sys().(*syscall.Win32FileAttributeData).CreationTime.Nanoseconds())
}
