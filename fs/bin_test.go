package fs

import (
	"fmt"
	"testing"
)

func TestWhich(t *testing.T) {
	fmt.Printf("Which(\"hotpot\"): %v\n", Which("hotpot"))
}
func TestWhich2(t *testing.T) {
	fmt.Printf("Which(\"hotpot\"): %v\n", Which("hotpot__"))
}

func TestWhich1(t *testing.T) {
	fmt.Printf("Which(\"hotpot\"): %v\n", Which(""))
}
