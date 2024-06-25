package adt

import "testing"

func TestRandomStringName(t *testing.T) {
	println(RandomStringName(10))
}

func BenchmarkRandomNumberName(b *testing.B) {
	for i := 0; i < b.N; i++ {
		RandomStringName(10)
	}
}

func BenchmarkName(b *testing.B) {
	//
	for i := 0; i < b.N; i++ {

	}
}
