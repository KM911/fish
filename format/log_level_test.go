package format

import "testing"

func TestExample(t *testing.T) {
	infoRender("debug")
	warningRender("warning")
	fatalRender("fatal")
	Success("success")
}
