package search

func BinarySearch(arr []int, target int) int {
	return binarySearch(arr, 0, len(arr)-1, target)
}

func binarySearch(arr []int, start, end int, target int) int {
	if start > end {
		return -1
	}
	mid := (start + end) / 2
	if arr[mid] == target {
		return mid
	} else if arr[mid] < target {
		return binarySearch(arr, mid+1, end, target)
	} else {
		return binarySearch(arr, start, mid-1, target)
	}
}
