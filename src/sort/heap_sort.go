package sort

// HeapSort 堆排序一般是一个完全二叉树
// 父求子：父亲的索引是i , 左孩子是2i+1，右孩子是2i+2 ,
// 子求父：(i-1)/2
func HeapSort(arr []int) {
	//1 构建堆，这里用到最大堆
	buildHeap(arr)
	//2 取出最大元素，放到数组的尾部，就是从小到大的顺序了。
	// 依次取出堆顶元素，放到数组末尾
	n := len(arr)
	for n > 0 {
		// 将堆顶元素和数组末尾元素交换位置
		arr[0], arr[n-1] = arr[n-1], arr[0]
		n--
		//去掉尾部的元素，然后，进行下沉操作
		elementDown(0, arr[:n]) //这里使用的是切片，没有开辟新空间
	}
}

// 从一个数组直接构建一个堆
func buildHeap(arr []int) {
	//拿到最后一个叶子节点的父节点，也就是非叶子节点，然后遍历所有非叶子节点
	for i := parentIndex(len(arr) - 1); i >= 0; i-- {
		elementDown(i, arr)
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

func elementDown(parentIndex int, arr []int) {
	n := len(arr)
	for leftChildIndex(parentIndex) < n { //判断是否有左孩子，没有左孩子，肯定没有有孩子
		i := leftChildIndex(parentIndex)
		maxIndex := i //默认左孩子最大
		//判断是否有右孩子
		if i+1 < n && arr[i+1] > arr[i] {
			maxIndex = i + 1
		}
		//不需要下沉
		if arr[parentIndex] > arr[maxIndex] {
			return
		} else {
			//下沉：选择最大的孩子进行交换
			arr[parentIndex], arr[maxIndex] = arr[maxIndex], arr[parentIndex]
			parentIndex = maxIndex
		}
	}
}
