package sort

func SelectSort(arr []int) {
	for i := 0; i < len(arr); i++ {
		min := i
		for j := i; j < len(arr); j++ {
			if arr[min] > arr[j] {
				min = j
			}
		}
		arr[min], arr[i] = arr[i], arr[min]
	}
}

func BubbleSort(arr []int) {
	for i := 0; i < len(arr); i++ {
		for j := 1; j < len(arr)-i; j++ {
			if arr[j-1] > arr[j] {
				arr[j-1], arr[j] = arr[j], arr[j-1]
			}
		}
	}
}

func InsertSort(arr []int) {
	for i := 1; i < len(arr); i++ {
		for j := i; j > 0 && arr[j] < arr[j-1]; j-- {
			arr[j], arr[j-1] = arr[j-1], arr[j]
		}
	}
}

// QuickSort sorts an array using the quicksort algorithm.
func QuickSort(arr []int) {
	if len(arr) < 2 {
		return
	}
	pivot := arr[0]
	left, right := 1, len(arr)-1
	for left < right {
		for left < right && arr[left] < pivot {
			left++
		}
		for left < right && arr[right] >= pivot {
			right--
		}
		arr[left], arr[right] = arr[right], arr[left]
	}
	if arr[left] >= pivot {
		left--
	}
	arr[0], arr[left] = arr[left], arr[0]
	QuickSort(arr[:left])
	QuickSort(arr[left+1:])
}
