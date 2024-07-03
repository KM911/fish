package format

import "github.com/gookit/color"

const (
	log_level = 0
	// debug
	// only display in debug mode
	// error -- warning
	// panic -- fatal
)

// info warning fatal

var (
	Info    = infoRender
	Warning = warningRender
	Fatal   = fatalRender
	Success = successRender
)

func nothing() {}

func Release() {
	Info = info
	Warning = warning
	Fatal = fatal
	Success = success
}

func info(v string) {
	LogFormat("Debug", v)
}

func infoRender(v string) {
	// LogFormat(color.BgHiBlue.Render("Debug"), color.Note.Render(v))
	LogFormat(color.BgHiBlue.Render("Debug"), v)

}

func success(v string) {
	LogFormat("Success", v)
}
func successRender(v string) {
	// LogFormat(color.BgGreen.Render("Success"), color.Success.Render(v))
	LogFormat(color.BgGreen.Render("Success"), v)
}

func warning(v string) {
	LogFormat("Warning", v)
}

func warningRender(v string) {
	// LogFormat(color.BgYellow.Render("Warning"), color.Warn.Render(v))
	LogFormat(color.BgYellow.Render("Warning"), v)

}

func fatal(v string) {
	LogFormat("Panic", v)
}

func fatalRender(v string) {
	// LogFormat(color.BgRed.Render("Error"), color.Error.Render(v))
	LogFormat(color.BgRed.Render("Error"), v)

}
