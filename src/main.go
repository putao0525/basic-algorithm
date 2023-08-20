package main

import (
	"basic-algorithm/src/sort"
	"fmt"
)

func testSort() {
	arr := []int{5, 8, 6, 3, 9, 2, 1, 7}
	//sort.BubbleSort(arr)
	//arr = []int{2, 3, 4, 5, 6, 7, 8, 1}
	sort.CockTailSort(arr)
	fmt.Printf("%v\n", arr)
}

func main() {
	testSort()
}
