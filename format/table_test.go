package format

import (
	"testing"
)

var (
	titles   = []string{"pid", "name", "status", "mem", "cpu"}
	contents = [][]string{
		{"1", "/opt/visual-studio-code/code", "running", "500M", "1%"},
		{"2", "Xwayland", "Sleeping", "153M", "2%"},
		{"3", "chrome", "running", "2.5GB", "30%"},
	}
)

func TestTable(t *testing.T) {
	TableMessage(titles, contents)
}
