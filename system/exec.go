package system

import (
	"os"
	"os/exec"
	"runtime"
)

var (
	CreateCommand func(command string) *exec.Cmd
)

func init() {
	if runtime.GOOS == "windows" {
		CreateCommand = createCmd
	} else {
		CreateCommand = createBash
	}
}

func redirectStd(cmd *exec.Cmd) {
	cmd.Stdout = os.Stdout
	cmd.Stdin = os.Stdin
	cmd.Stderr = os.Stderr
}

func createCmd(command string) (cmd *exec.Cmd) {

	cmd = exec.Command("cmd", "/C", command)

	return
}

func createBash(command string) *exec.Cmd {
	return exec.Command("bash", "-c", command)
}

func ExecuteCommand(command string) int {
	cmdExecutor := CreateCommand(command)
	redirectStd(cmdExecutor)
	cmdExecutor.Run()
	return cmdExecutor.ProcessState.ExitCode()
}

func ExecuteCommandSilent(command string) int {
	cmdExecutor := CreateCommand(command)
	cmdExecutor.Run()
	return cmdExecutor.ProcessState.ExitCode()
}

func ExecuteCommandResult(command string) string {
	cmdExecutor := CreateCommand(command)
	output, _ := cmdExecutor.Output()
	os.Stdout.Write(output)
	return string(output)

}

func ExecuteCommandSilentResult(command string) string {
	cmdExecutor := CreateCommand(command)
	output, _ := cmdExecutor.Output()
	return string(output)
}
