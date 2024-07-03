package sort

import (
	"fmt"
	"testing"

	"github.com/KM911/fish/adt"
)

var (
	case_16        = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16}
	case_16_best   = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 16, 15}
	case_16_middle = []int{8, 7, 6, 5, 4, 3, 2, 1, 16, 15, 14, 13, 12, 11, 10, 9}
	case_16_wrost  = []int{16, 15, 14, 13, 12, 11, 10, 9, 8, 7, 6, 5, 4, 3, 2, 1}
)

func TestBubbleSort(t *testing.T) {

	BubbleSort(case_16_best)
	if !adt.IsSameArray(case_16, case_16_best) {

		t.Error("BubbleSort failed")
	}
	BubbleSort(case_16_middle)
	if !adt.IsSameArray(case_16, case_16_middle) {
		t.Error("BubbleSort failed")
	}
	BubbleSort(case_16_wrost)
	if !adt.IsSameArray(case_16, case_16_wrost) {
		t.Error("BubbleSort failed")
	}
	fmt.Println(case_16_best, case_16_middle, case_16_wrost)

}
