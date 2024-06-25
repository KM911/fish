package fs

import (
	"fmt"
	"testing"
)

func TestDir(t *testing.T) {
	fmt.Println(Dir(CmdPath()))
}

func TestIsEmptyDir(t *testing.T) {
	// do some assert
	if IsEmptyDir("D:/TOOL") {
		t.Error("D:/TOOL is not empty")
	}
}

func TestSortFoldersByDepth(t *testing.T) {
	folders := []string{
		"a/b", "a/c", "a/a/a", "a",
	}
	if SortByDepth(folders)[0] != "a/a/a" {
		t.Error("SortFoldersByDepth failed")
	}

}

func TestDeepDir(t *testing.T) {
	fmt.Println(DeepDir("D:/TOOL"))
}

func TestIsDir(t *testing.T) {
	if !IsDir("D:/TOOL") {
		t.Error("D:/TOOL is a dir")
	}
}
