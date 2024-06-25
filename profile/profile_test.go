package profile

import (
	"fmt"
	"testing"
	"time"
)

func TestProfile(t *testing.T) {
	defer Profile(0).Stop()
	for i := 0; i < 1000; i++ {
		fmt.Println(i)
		time.Sleep(time.Millisecond * 10)
	}
}
