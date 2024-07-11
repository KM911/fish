package format

import "github.com/gookit/color"

const (
	log_level = 0
	// debug
	// only display in debug mode
	// error -- warning
	// panic -- fatal
)

var (
	Info    = infoRender
	Warning = warningRender
	Fatal   = fatalRender
	Success = successRender
)

func nothing(str string) {}

// close all info
func Silent() {
	Info = nothing
	Warning = nothing
	Fatal = nothing
	Success = nothing
}

func Common() {
	Info = info
	Warning = warning
	Fatal = fatal
	Success = success
}

func info(v string) {
	FileLogFormat("Debug", v)
}

func InfoMsg(title, content string) {
	TerminalLogFormat(color.BgHiBlue.Render(title), content)
}

func NoteMsg(title, content string) {
	TerminalLogFormat(color.Note.Render(title), content)
}

func WarningMsg(title, content string) {
	TerminalLogFormat(color.BgYellow.Render(title), content)
}

func FatalMsg(title, content string) {
	TerminalLogFormat(color.BgRed.Render(title), content)
}

func SuccessMsg(title, content string) {
	TerminalLogFormat(color.BgGreen.Render(title), content)
}

func infoRender(v string) {
	FileLogFormat(color.BgHiBlue.Render("Info"), v)

}

func success(v string) {
	FileLogFormat("Success", v)
}
func successRender(v string) {
	FileLogFormat(color.BgGreen.Render("Success"), v)
}

func warning(v string) {
	FileLogFormat("Warning", v)
}

func warningRender(v string) {
	FileLogFormat(color.BgYellow.Render("Warning"), v)
}

func fatal(v string) {
	FileLogFormat("Panic", v)
}

func fatalRender(v string) {
	FileLogFormat(color.BgRed.Render("Error"), v)

}
