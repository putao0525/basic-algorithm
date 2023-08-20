package sort

import "fmt"

// HeapSort 堆排序一般是一个完全二叉树， 所以，有规律.
// 父求子：父亲的索引是i , 左孩子是2i，右孩子是2i+1 ,
// 子求父：(i-1)/2
func HeapSort(arr []int) {
	//1 构建堆，这里用到最大堆
	buildHeap(arr)
	fmt.Printf("构建堆：%v\n", arr)
	//2 取出最大元素，放到数组的尾部，就是从小到大的顺序了。
	// 依次取出堆顶元素，放到数组末尾
	n := len(arr)
	for n > 0 {
		// 将堆顶元素和数组末尾元素交换位置
		arr[0], arr[n-1] = arr[n-1], arr[0]
		n--
		siftDown(0, arr[:n])
	}
}

func buildHeap(arr []int) {
	//拿到最后一个非叶子节点，然后遍历所有非叶子节点
	for i := parentIndex(len(arr) - 1); i >= 0; i-- {
		siftDown(i, arr)
	}
}

func parentIndex(i int) int {
	return (i - 1) / 2
}
func leftChildIndex(i int) int {
	return i*2 + 1
}
func rightChildIndex(i int) int {
	return i*2 + 2
}

func siftDown(parentIndex int, arr []int) {
	n := len(arr)
	for leftChildIndex(parentIndex) < n {
		i := leftChildIndex(parentIndex)
		maxIndex := i
		if i+1 < n && arr[i+1] > arr[i] {
			maxIndex = i + 1
		}
		if arr[parentIndex] > arr[maxIndex] {
			return
		} else {
			arr[parentIndex], arr[maxIndex] = arr[maxIndex], arr[parentIndex]
			parentIndex = maxIndex
		}
	}
}
