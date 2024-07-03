package format

import (
	"fmt"
	"runtime"
	"strconv"
	"testing"
	"time"
)

func TestLog(t *testing.T) {
	LogFormat("title", "content")
}

var (
	title_   = "title"
	content_ = "content hello world"
)

func BenchmarkBuilderLog(b *testing.B) {
	// output.Write([]byte(fmt.Sprintf("%s %s:%d %s\n%s\n", time.Now().Format("2006-01-02 15:04:05"), file, line, title, content)))
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_, file, line, _ := runtime.Caller(1)
		s := fmt.Sprintf("%s %s:%d %s\n%s\n", time.Now().Format("2006-01-02 15:04:05"), file, line, title, content_)
		_ = s
	}
}

func BenchmarkFormatLog(b *testing.B) {
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		_, file, line, _ := runtime.Caller(1)
		StringBuilder.WriteString(time.Now().Format("2006-01-02 15:04:05"))
		StringBuilder.WriteByte(' ')
		StringBuilder.WriteString(file)
		StringBuilder.WriteByte(':')
		StringBuilder.WriteString(strconv.Itoa(line))
		StringBuilder.WriteString(title_)
		StringBuilder.WriteByte('\n')
		StringBuilder.WriteString(content_)
		s := []byte(StringBuilder.String())
		_ = s
	}
}
