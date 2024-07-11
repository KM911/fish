package fs

import (
	"os"
	"os/exec"
)

/*
should use exec.LookPath
*/
func Which(src string) string {
	r, w, _ := os.Pipe()
	which := exec.Cmd{
		Path:   "/usr/bin/which",
		Args:   []string{"", src},
		Stdout: w,
		Stderr: w,
	}
	which.Run()
	buffer := make([]byte, 128)
	n, _ := r.Read(buffer)
	if which.ProcessState.ExitCode() != 0 {
		return ""
	}
	return string(buffer[:n])
}
