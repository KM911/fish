package format

import (
	"strings"
	"unicode"
)

var (
	StringBuilder = strings.Builder{}
)

func MaxLength(s []string) int {
	currentMax := 0
	for _, v := range s {
		if len(v) > currentMax {
			currentMax = len(v)
		}
	}
	return currentMax
}

// StringLen return the width of string
// abcd 4
// 你好 4
func StringLen(s string) int {
	length := 0
	for _, c := range s {
		// unicode
		if unicode.Is(unicode.Han, c) {
			length += 2
		} else {
			length++
		}
	}

	return length
}
