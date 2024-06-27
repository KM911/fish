package format

import (
	"strings"
	"unicode"
	"unsafe"
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

// StringToBytes converts a string to a byte slice without memory allocation.
// NOTE: The returned byte slice MUST NOT be modified since it shares the same backing array
// with the given string.
func StringToBytes(s string) []byte {
	p := unsafe.StringData(s)
	b := unsafe.Slice(p, len(s))
	return b
}

// BytesToString converts bytes to a string without memory allocation.
// NOTE: The given bytes MUST NOT be modified since they share the same backing array
// with the returned string.
func BytesToString(b []byte) string {
	// Ignore if your IDE shows an error here; it's a false positive.
	p := unsafe.SliceData(b)
	return unsafe.String(p, len(b))
}
