package fs

import (
	"os"
	"strings"
)

// os.Rename could not cross the driver

func CheckCrossDriverWindows(src, dst string) bool {
	// this is for windows

	// for linux /tmp /home will cross the driver
	srcDriver := strings.Split(src, ":")[0]
	dstDriver := strings.Split(dst, ":")[0]
	return srcDriver != dstDriver
}
func CheckCrossDriverLinux(src, dst string) bool {
	srcDriver := strings.Split(src, "/")[0]
	dstDriver := strings.Split(dst, "/")[0]
	return srcDriver != dstDriver
}

func MoveFile(src, dst string) (err error) {
	// do not user os.Rename
	// because it could not cross the driver

	// 1. check cross the driver
	if !CheckCrossDriverLinux(src, dst) {
		err = os.Rename(src, dst)
		if err != nil {
			return
		}
		return
	}
	// 2. copy file
	data, err := os.ReadFile(src)
	if err != nil {
		return
	}
	file, err := os.Create(dst)
	if err != nil {
		return
	}

	file.Write(data)
	file.Close()
	return nil
}
