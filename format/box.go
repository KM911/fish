package format

import (
	"strings"

	"github.com/gookit/color"
)

// ╭─────────┬──────────╮
// │ _title  │ _content │
// ╰─────────┴──────────╯

func boxMessage(_title string, _content string, l1 int, l2 int) {
	formatStringBuilder.WriteString("╭")
	formatStringBuilder.WriteString(strings.Repeat("─", l1+2))
	formatStringBuilder.WriteString("┬")
	formatStringBuilder.WriteString(strings.Repeat("─", l2+2))
	formatStringBuilder.WriteString("╮\n│ ")
	formatStringBuilder.WriteString(_title)
	formatStringBuilder.WriteString(" │ ")
	formatStringBuilder.WriteString(_content)
	formatStringBuilder.WriteString(" │\n╰")
	formatStringBuilder.WriteString(strings.Repeat("─", l1+2))
	formatStringBuilder.WriteString("┴")
	formatStringBuilder.WriteString(strings.Repeat("─", l2+2))
	formatStringBuilder.WriteString("╯\n")
	println(formatStringBuilder.String())
	formatStringBuilder.Reset()
}

func BoxError(_title string, _content string) {

	boxMessage(color.Error.Render(_title), _content, len(_title), len(_content))
}

func BoxWarning(_title string, _content string) {
	boxMessage(color.Warn.Render(_title), _content, len(_title), len(_content))
}

func BoxInfo(_title string, _content string) {
	boxMessage(color.Info.Render(_title), _content, len(_title), len(_content))
}

func BoxSuccess(_title string, _content string) {
	boxMessage(color.Success.Render(_title), _content, len(_title), len(_content))
}

func BoxMessage(_title string, _content string) {
	boxMessage(color.Note.Render(_title), _content, len(_title), len(_content))
}
