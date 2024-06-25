package adt

import (
	"math/rand"
	"strings"
)

var StringTable = []string{
	"0", "1", "2", "3", "4", "5", "6", "7", "8", "9",
	"a", "b", "c", "d", "e", "f", "g", "h", "i",
	"j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z",
	"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z"}

/*
base64是有 10 + 26 +26 = 62个字符组成的
*/
func RandomStringName(length int) (StringName string) {
	StringBuilder := strings.Builder{}
	for i := 0; i < length; i++ {
		StringBuilder.WriteString(StringTable[rand.Intn(62)])
	}
	StringName = StringBuilder.String()
	return StringName
}
