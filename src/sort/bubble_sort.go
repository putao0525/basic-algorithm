package sort

/**
冒泡排序
*/

func BubbleSort(arr []int) {
	n := len(arr)
	//为啥是n-1?
	//比如按照从小到大排序，每次冒泡最大的，其他的都冒完了， 数据中最后一个元素不需要冒泡了。所以少一轮循环
	for i := 0; i < n-1; i++ {
		//为啥是n-i-1?
		//8个元素需要比较7次，7个元素需要比较6次，在减去已经冒泡好的数据
		for j := 0; j < n-i-1; j++ {
			if arr[j] >= arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
}
