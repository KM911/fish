package format

import "testing"

func TestRecover(t *testing.T) {
	SingleLogger("log.log")
	defer Recover()
	panic("test")
}
