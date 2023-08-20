package main

import (
	"basic-algorithm/src/sort"
	"fmt"
)

func testSort() {
	arr := []int{5, 8, 6, 3, 9, 2, 1, 7}
	sort.BubbleSort(arr)
	fmt.Printf("%v\n", arr)
}

func main() {
	testSort()
}
