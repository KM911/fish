package system

import "os/exec"

// could not solve this problem
// For example, /usr/bin/code will open /opt/vscode/code
// A just open B, and they have different name
// Warning : ""
func IsAlive(src string) bool {
	pgrep := exec.Cmd{
		Path: "/usr/bin/pgrep",
		Args: []string{"", src},
	}
	pgrep.Run()
	return pgrep.ProcessState.Success()
}
func RunBackground(args []string) {
	app := exec.Cmd{
		Path: args[0],
		Args: args,
	}
	app.Start()
}
