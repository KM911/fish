package format

import (
	"fmt"
	"time"
)

type timer struct {
	start time.Time
}

func TimerStart() *timer {
	t := timer{}
	t.start = time.Now()
	return &t
}

func (_t *timer) End() {
	InfoMessage("Past ", fmt.Sprint(time.Since(_t.start)))
}

func UnixTime() int64 {
	return time.Now().Unix()
}

func Timestamp() string {
	time.AfterFunc(1*time.Second, func() {})
	return time.Now().Format("2006-01-02 15:04:05")
}
