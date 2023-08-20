package sort

func QuickSort(arr []int, startIndex, endIndex int) {
	if startIndex >= endIndex {
		return
	}
	//pivotIndex := partition(arr, startIndex, endIndex)
	pivotIndex := partitionV2(arr, startIndex, endIndex)
	QuickSort(arr, startIndex, pivotIndex-1)
	QuickSort(arr, pivotIndex+1, endIndex)
}

// 随机的选择一个元算， 比这个元素的大的，放在右边，比这个元素小的，放在左边
// 返回这个元素的位置
// 重复这个操作， 最终有序
// 双边快速排序
func partition(arr []int, startIndex, endIndex int) int {
	pivot := arr[startIndex] //备份第一元素的地址 ，
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

// 单边快速排序
func partitionV2(arr []int, startIndex, endIndex int) int {
	pivot := arr[startIndex]
	mark := startIndex
	for i := startIndex + 1; i <= endIndex; i++ {
		if arr[i] < pivot {
			mark++
			arr[i], arr[mark] = arr[mark], arr[i]
		}
	}
	arr[startIndex] = arr[mark]
	arr[mark] = pivot
	return mark
}
