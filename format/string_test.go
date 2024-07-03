package format

import (
	"bytes"
	"strings"
	"testing"
)

// String --> Byte []byte(string)
// Byte --> String ByteToString([]byte)
// BenchmarkStringToBytesStandard-12    	1000000000	         0.2497 ns/op	       0 B/op	       0 allocs/op
// BenchmarkBytesToStringStandard-12    	 6222952	       190.5 ns/op	    1792 B/op	       1 allocs/op
// BenchmarkStringToBytes-12            	1000000000	         0.4987 ns/op	       0 B/op	       0 allocs/op
// BenchmarkBytesToString-12            	1000000000	         0.4830 ns/op	       0 B/op	       0 allocs/op
const size = 310

func TestStringToBytes(t *testing.T) {
	s := "hello"
	b := StringToBytes(s)
	if !bytes.Equal(b, []byte(s)) {
		t.Errorf("expected %s, got %s", s, string(b))
	}

	c := append(b, " world"...)
	if !bytes.Equal(c, []byte("hello world")) {
		t.Errorf("expected %s, got %s", "hello world", string(b))
	}

	// s must be unchanged.
	if !bytes.Equal(b, []byte(s)) {
		t.Errorf("expected %s, got %s", s, string(b))
	}
}

func BenchmarkStringToBytesStandard(b *testing.B) {
	s := strings.Repeat("hello", size)
	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_ = []byte(s)
	}
}

func BenchmarkBytesToStringStandard(b *testing.B) {
	bb := []byte(strings.Repeat("hello", size))
	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_ = string(bb)
	}
}

func BenchmarkStringToBytes(b *testing.B) {
	s := strings.Repeat("hello", size)
	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_ = StringToBytes(s)
	}
}

func TestBytesToString(t *testing.T) {
	b := []byte{'h', 'e', 'l', 'l', 'o'}
	s := BytesToString(b)
	if !bytes.Equal(b, []byte(s)) {
		t.Errorf("expected %s, got %s", string(b), s)
	}

	// change b, s must be unchanged.
	// b will allocate a new backing array, and s will still point to the old one.
	b = append(b, " world"...)
	if !bytes.Equal(b, []byte("hello world")) {
		t.Errorf("expected %s, got %s", "hello world", string(b))
	}
}

func BenchmarkBytesToString(b *testing.B) {
	bb := []byte(strings.Repeat("hello", size))
	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_ = BytesToString(bb)
	}
}
