package system

import (
	"fmt"
	"testing"
)

func TestExecuteCommand(t *testing.T) {
	ExecuteCommand("ls")
	ExecuteCommand("pwd")
}

func TestExecuteCommandResult(t *testing.T) {
	fmt.Printf("ExecuteCommandResult(\"ls\"): %v\n", ExecuteCommandResult("pwd"))
	fmt.Printf("ExecuteCommandSilentResult(\"ls\"): %v\n", ExecuteCommandSilentResult("ls"))
}
