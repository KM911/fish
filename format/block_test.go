package format

import "testing"

var (
	title     = "Programming"
	languages = []string{
		"go", "rust", "python", "java",
	}
)

func TestBlockMessage(t *testing.T) {
	BlockMessage(title, languages)
}
