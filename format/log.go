package format

import (
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"runtime/debug"
	"strings"
	"time"
)

var (
	StringBuilder = strings.Builder{}
	output        = os.Stdout
)

// time file:line Error title
// Error content
func init() {
}

func SetOutputTemp() {
	SetOutputFile(filepath.Join(os.TempDir(), "format_"+time.Now().GoString()))
}

func SetOutputFile(src string) {
	logFile, err := os.OpenFile(src, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0766)
	if err != nil {
		panic(err)
	}
	output = logFile
}

func SetOutput(w io.Writer) {
	log.SetOutput(w)
}

func Recover() {
	if err := recover(); err != nil {
		log.Println(err)
		log.Println(string(debug.Stack()))
	}
}

/*
time file:line Error title
Error content
*/
func LogFormat(title, content string) {
	_, file, line, _ := runtime.Caller(1)

	// StringBuilder.WriteString(time.Now().Format("2006-01-02 15:04:05"))
	// StringBuilder.WriteByte(' ')
	// StringBuilder.WriteString(file)
	// StringBuilder.WriteByte(':')
	// StringBuilder.WriteString(strconv.Itoa(line))
	// StringBuilder.WriteString(title)
	// StringBuilder.WriteByte('\n')
	// StringBuilder.WriteString(content)

	// output.Write([]byte(fmt.Sprintf("%s %s:%d %s\n%s\n", time.Now().Format("2006-01-02 15:04:05"), file, line, title, content)))

	fmt.Fprintf(output, "%s %s:%d %s\n%s\n", title, file, line, time.Now().Format("2006-01-02 15:04:05"), content)
}
