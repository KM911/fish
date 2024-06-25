package fs

import (
	"fmt"
	"testing"
)

func TestAbsPath(t *testing.T) {
	if AbsPath("D:/TOOL") != "D:/TOOL" {
		t.Error("AbsPath failed")
	}
	if AbsPath("1.txt") != CmdPath()+"/1.txt" {
		t.Error("AbsPath failed")
	}
}

func TestCmdPath(t *testing.T) {
	fmt.Println(CmdPath())
}

func TestExecutePath(t *testing.T) {
	fmt.Println(ExecutePath())
}
