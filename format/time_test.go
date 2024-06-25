package format

import (
	"testing"
	"time"
)

func TestTimeLock(t *testing.T) {
	defer TimerStart().End()
	time.Sleep(1 * time.Second)
}

func BenchmarkSub(b *testing.B) {
	start := time.Now()
	for i := 0; i < b.N; i++ {
		_ = time.Now().Sub(start)
	}
}

func BenchmarkSince(b *testing.B) {
	start := time.Now()
	for i := 0; i < b.N; i++ {
		_ = time.Since(start)
	}
}
