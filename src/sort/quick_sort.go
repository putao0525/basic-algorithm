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
		arr[left] = arr[right] //右边的元素放到了左边， 右边的位置就可以放入左边的元素了
		for right > left && arr[left] < pivot {
			left++
		}
		arr[right] = arr[left] //左边的元素放入右边，左边的元素就可以放入右边的元素了
	}
	//此时，左右相等，存在一个重复的元素
	arr[left] = pivot //用备份的元素覆盖重复的元素
	return left
}
