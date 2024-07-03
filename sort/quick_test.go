package sort

import (
	"testing"

	"github.com/KM911/fish/adt"
)

func TestQuickSort(t *testing.T) {

	QuickSort(case_16_best)
	if !adt.IsSameArray(case_16, case_16_best) {
		t.Error("QuickSort failed")
	}
	QuickSort(case_16_middle)
	if !adt.IsSameArray(case_16, case_16_middle) {
		t.Error("QuickSort failed")
	}
	QuickSort(case_16_wrost)
	if !adt.IsSameArray(case_16, case_16_wrost) {
		t.Error("QuickSort failed")
	}
}
