package fc

import "testing"

func TestFilter(t *testing.T) {
	array := []int{1, 2, 3, 4, 5, 6}
	result := Filter(array, func(value int) bool {
		return value > 3
	})
	if len(result) != 3 {
		t.Error("Filter failed")
	}
	if result[0] != 4 {
		t.Error("Filter failed")
	}
	if result[1] != 5 {
		t.Error("Filter failed")
	}
	if result[2] != 6 {
		t.Error("Filter failed")
	}

}

func BenchmarkFilter(b *testing.B) {
	array := []int{1, 2, 3, 4, 5, 6}
	isodd := func(value int) bool {
		return value%2 == 1
	}

	for i := 0; i < b.N; i++ {
		Filter(array, isodd)
	}

}
func BenchmarkFilterFor(b *testing.B) {
	array := []int{1, 2, 3, 4, 5, 6}
	isodd := func(value int) bool {
		return value%2 == 1
	}
	for i := 0; i < b.N; i++ {
		var result []int
		for _, value := range array {
			if isodd(value) {
				result = append(result, value)
			}
		}
	}

}
