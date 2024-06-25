package format

import (
	"log"
	"os"
	"runtime/debug"
)

// 其实不对呀

type LogOption struct {
	disable    bool
	level      int
	outputDest string
}

func SingleLogger(src string) {
	logFile, err := os.OpenFile(src, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0766)
	if err != nil {
		panic(err)
	}
	log.SetOutput(logFile)
}

func Recover() {
	if err := recover(); err != nil {
		log.Println(err)
		log.Println(string(debug.Stack()))
	}
}
