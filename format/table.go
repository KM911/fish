package format

import (
	"fmt"
	"strings"
)

// top     ╭ ─ ┬ ╮
// middle  ├ ─ ┼ ┤
// content │   │ │
// bottom  ╰ ─ ┴ ╯

func TableMessage(_title []string, _content [][]string) {
	tableLength := make([]int, len(_title))
	for i, v := range _title {
		tableLength[i] = StringLen(v) + 1
	}
	contentLength := make([]int, len(_title))
	for col := range contentLength {
		contentLength[col] = 0
	}
	for row := range _content {
		for col := range _content[row] {
			if StringLen(_content[row][col])+1 > contentLength[col] {
				contentLength[col] = StringLen(_content[row][col]) + 1
			}
		}
	}

	for i := range _title {
		if contentLength[i] > tableLength[i] {
			tableLength[i] = contentLength[i]
		}
	}
	DrawTableLine(tableLength, "╭", "─", "┬", "╮")
	DrawTableContent(_title, tableLength)
	DrawTableLine(tableLength, "├", "─", "┼", "┤")
	for _, v := range _content {
		DrawTableContent(v, tableLength)
	}
	DrawTableLine(tableLength, "╰", "─", "┴", "╯")
	fmt.Println(formatStringBuilder.String())
	formatStringBuilder.Reset()
}

func DrawTableLine(length []int, left string, seq string, split string, right string) {
	formatStringBuilder.WriteString(left)
	// formatStringBuilder.WriteString(strings.Repeat(seq, length))
	for i, v := range length {
		formatStringBuilder.WriteString(strings.Repeat(seq, v+1))
		if i != len(length)-1 {
			formatStringBuilder.WriteString(split)
		}
	}
	formatStringBuilder.WriteString(right)
	formatStringBuilder.WriteString("\n")
}

func DrawTableContent(content []string, length []int) {
	formatStringBuilder.WriteString("│ ")
	for i, v := range content {
		formatStringBuilder.WriteString(v)
		formatStringBuilder.WriteString(strings.Repeat(" ", length[i]-StringLen(v)))
		formatStringBuilder.WriteString("│ ")

	}
	formatStringBuilder.WriteString("\n")
}
