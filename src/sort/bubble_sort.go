package sort

// BubbleSort 冒泡排序
func BubbleSort(arr []int) {
	n := len(arr)
	//为啥是n-1?
	//比如按照从小到大排序，每次冒泡最大的，其他的都冒完了， 数据中最后一个元素不需要冒泡了。所以少一轮循环
	for i := 0; i < n-1; i++ {
		//为啥是n-i-1?
		//8个元素需要比较7次，7个元素需要比较6次，在减去已经冒泡好的数据
		isSorted := true //如果某一轮已经完全有序了，下一轮就不需要执行了
		for j := 0; j < n-i-1; j++ {
			if arr[j] >= arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
				isSorted = false
			}
		}
		if isSorted {
			return
		}
	}
}

// CockTailSort 鸡尾酒排序
func CockTailSort(arr []int) {
	n := len(arr)
	//为啥外层 n/2 内层一次 ，先从左往右冒泡，然后，从右往左冒泡
	for i := 0; i < n/2; i++ {
		isSorted := true //如果某一轮已经完全有序了，下一轮就不需要执行了
		for j := i; j < n-i-1; j++ {
			if arr[j] >= arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
				isSorted = false
			}
		}
		if isSorted {
			return
		}
		isSorted = true
		for j := n - i - 1; j > i; j-- {
			if arr[j-1] >= arr[j] {
				arr[j-1], arr[j] = arr[j], arr[j-1]
				isSorted = false
			}
		}
		if isSorted {
			return
		}
	}
}
