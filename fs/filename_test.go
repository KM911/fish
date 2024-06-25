package fs

import (
	"fmt"
	"testing"
)

func TestFileName(t *testing.T) {
	if FileName("D:/TOOL/1.txt") != "1" {
		fmt.Println(FileName("D:/TOOL/1.txt"))
		t.Error("FileName failed")
	}
}

func BenchmarkFileName(b *testing.B) {
	for i := 0; i < b.N; i++ {
		FileName("D:/TOOL/1.txt")
	}
}

func TestFullFileName(t *testing.T) {
	if FullFileName("D:/TOOL/1.txt") != "D:/TOOL/1" {
		t.Error("FullFileName failed")
	}
}

func TestIsExit(t *testing.T) {
	if !IsExit("D:/TOOL") {
		t.Error("D:/TOOL is exit")
	}
	if IsExit("D:/TOOL/1.txt") {
		t.Error("D:/TOOL/1.txt is exit")
	}
}

func TestChangeFileType(t *testing.T) {
	if ChangeFileType("D:/TOOL/1.txt", "mp4") != "D:/TOOL/1.mp4" {
		t.Error("ChangeFileType failed")
	}
	if ChangeFileType("1.txt", "mp4") != "1.mp4" {
		t.Error("ChangeFileType failed")
	}
	if ChangeFileType("1", "mp4") != "1.mp4" {
		t.Error("ChangeFileType failed")
	}
}

func TestIsFileType(t *testing.T) {
	if !IsFileType("D:/TOOL/1.txt", "txt") {
		t.Error("IsFileType failed")
	}
	if IsFileType("D:/TOOL/1.txt", "mp4") {
		t.Error("IsFileType failed")
	}
}
