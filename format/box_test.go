package format

import "testing"

func TestBoxMessage(t *testing.T) {
	BoxMessage("message", "message")
	BoxError("error", "error")
	BoxWarning("warning", "warning")
	BoxInfo("info", "info")
	BoxSuccess("success", "success")
}
