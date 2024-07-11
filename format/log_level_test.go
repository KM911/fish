package format

import "testing"

func TestExample(t *testing.T) {
	infoRender("debug")
	warningRender("warning")
	fatalRender("fatal")
	Success("success")
}

func TestMsg(t *testing.T) {
	NoteMsg("note", "message")
	InfoMsg("info", "message")
	WarningMsg("warning", "message")
	FatalMsg("fatal", "message")
	SuccessMsg("success", "message")
}
