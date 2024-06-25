package key

import (
	"fmt"
	"strings"
	"testing"
)

func shouhuConvent(url string) string {
	// https://tv.sohu.com/v/dXMvMzM1OTQyNzA5LzE3MjM5ODAzOC5zaHRtbA==.html
	code := url[22 : len(url)-5]
	fmt.Println(code)
	// 通过base64解码
	decodeString := Base64Decode(code)
	// us/335942709/172398038.shtml
	// we only need 335942709
	fmt.Println(decodeString)
	split := strings.Split(decodeString, "/")
	if len(split) < 2 {
		// logs.Errorf("split error: %v", split)
		return ""
	}
	return split[1]
}

func TestBase64(t *testing.T) {
	// str := "dXMvMzM1OTQyNzA5LzE3MjM5ODAzOC5zaHRtbA=="
	fmt.Println(shouhuConvent("https://tv.sohu.com/v/dXMvMzM1OTQyNzA5LzE3MjM5ODAzOC5zaHRtbA==.html"))
	println(len("https://tv.sohu.com/v/"))
}
