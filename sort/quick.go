package sort

func QuickSort(arr []int) {
	// 只有一个元素或者没有元素, 直接返回
	if len(arr) < 2 {
		return
	}
	// 选取第一个元素作为pivot
	pivot := arr[0]

	left, right := 1, len(arr)-1

	for left < right {
		// 一次循环里, 会进行两次交换
		for left < right && arr[left] < pivot {
			left++
		}
		for left < right && arr[right] >= pivot {
			right--
		}
		arr[left], arr[right] = arr[right], arr[left]
	}
	// 这一步是为了将我们的pivot放到正确的位置上
	if arr[left] >= pivot {
		left--
	}
	arr[0], arr[left] = arr[left], arr[0]
	// 递归版本的快排
	QuickSort(arr[:left])
	QuickSort(arr[left+1:])
}
