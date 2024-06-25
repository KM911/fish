package fc

import (
	"fmt"
	"strings"
	"testing"
)

func TestMap(t *testing.T) {
	stringSlice := []string{"A", "B", "C"}
	fmt.Println(Map(stringSlice, strings.ToLower))
}

func BenchmarkMap(b *testing.B) {
	b.ReportAllocs()
	stringSlice := []string{"A", "B", "C"}
	for i := 0; i < b.N; i++ {
		Map(stringSlice, strings.ToLower)
	}
}

// 结论有了呀 就是我们的
func BenchmarkMapAppend(b *testing.B) {
	b.ReportAllocs()
	stringSlice := []string{"A", "B", "C"}
	for i := 0; i < b.N; i++ {
		MapAppend(stringSlice, strings.ToLower)
	}
}

func BenchmarkMapSameType(b *testing.B) {
	b.ReportAllocs()
	stringSlice := []string{"A", "B", "C"}
	for i := 0; i < b.N; i++ {
		_ = MapSameType(stringSlice, strings.ToLower)
	}
}

func BenchmarkMapFor(b *testing.B) {
	b.ReportAllocs()
	stringSlice := []string{"A", "B", "C"}
	for i := 0; i < b.N; i++ {
		newSlice := make([]string, len(stringSlice))
		for index, value := range stringSlice {
			newSlice[index] = strings.ToLower(value)
		}
	}
}
