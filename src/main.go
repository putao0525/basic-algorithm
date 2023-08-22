package main

import (
	"basic-algorithm/src/sort"
	"fmt"
)

func testSort() {
	arr := []int{5, 8, 6, 3, 9, 2, 1, 7, 4}
	//sort.BubbleSort(arr)
	//arr = []int{2, 3, 4, 5, 6, 7, 8, 1}
	//sort.CockTailSort(arr)
	//sort.QuickSort(arr, 0, len(arr)-1)
	//sort.HeapSort(arr)
	//fmt.Printf("%v\n", search.BinarySearch(arr, 3))
	//sort.SelectSort(arr)
	fmt.Printf("%v\n", sort.MergeSort(arr))
	//fmt.Printf("%v\n", arr)
}

func main() {
	testSort()
}
