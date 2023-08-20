package sort

func QuickSort(arr []int, startIndex, endIndex int) {
	if startIndex >= endIndex {
		return
	}
	pivotIndex := partition(arr, startIndex, endIndex)
	QuickSort(arr, startIndex, pivotIndex-1)
	QuickSort(arr, pivotIndex+1, endIndex)
}

func partition(arr []int, startIndex, endIndex int) int {
	pivot := arr[startIndex] //备份第一元素的地址
	left := startIndex
	right := endIndex
	for left != right { //相当于while循环
		for right > left && arr[right] >= pivot {
			right--
		}
		arr[left] = arr[right]
		for right > left && arr[left] < pivot {
			left++
		}
		arr[right] = arr[left]
	}
	arr[left] = pivot
	return left
}
