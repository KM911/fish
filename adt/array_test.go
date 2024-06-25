package adt

import "testing"

func TestInArray(t *testing.T) {
	case1 := []string{"a", "b", "c"}
	if !InArray("a", case1) {
		t.Error("InArray failed")
	}
	case2 := []int{1, 2, 3}
	if !InArray(1, case2) {
		t.Error("InArray failed")
	}
	case3 := []float64{1.1, 2.2, 3.3}
	if !InArray(1.1, case3) {
		t.Error("InArray failed")
	}

}

func TestIsSameArray(t *testing.T) {
	case0 := []string{"a", "b", "c"}
	case1 := []string{"a", "b", "c"}
	case2 := []string{"a", "b"}
	case3 := []string{"a", "b", "d"}
	if !IsSameArray(case0, case1) {
		t.Error("IsSameArray failed")
	}
	if IsSameArray(case0, case2) {
		t.Error("IsSameArray failed")
	}
	if IsSameArray(case0, case3) {
		t.Error("IsSameArray failed")
	}
}