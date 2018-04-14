package main

import "fmt"

func recursiveBinarySearch(array []int, target, lowIndex, highIndex int) int {
	if highIndex < lowIndex {
		return -1
	}
	mid := int((lowIndex + highIndex) / 2)
	if array[mid] > target {
		return recursiveBinarySearch(array, target, lowIndex, mid)
	} else if array[mid] < target {
		return recursiveBinarySearch(array, target, mid+1, highIndex)
	} else {
		return mid
	}
}

func iterativeBinarySearch(array []int, target, startIndex, endIndex int) int {
	for startIndex < endIndex {
		mid := int((startIndex + endIndex) / 2)
		if array[mid] > target {
			endIndex = mid
		} else if array[mid] < target {
			startIndex = mid
		} else {
			return mid
		}
	}
	return -1
}

func main() {
	slice := []int{1, 2, 3, 4, 5, 6, 6, 10, 15, 15, 17, 18, 19, 30}
	find := recursiveBinarySearch(slice, 10, 0, len(slice)-1)
	fmt.Printf("Find the element by recursive binary search at index %d \n", find)
	find = iterativeBinarySearch(slice, 10, 0, len(slice)-1)
	fmt.Printf("Find the element by iterative binary search at index %d \n", find)
}
