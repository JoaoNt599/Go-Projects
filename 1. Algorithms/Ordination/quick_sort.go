package main

import (
	"fmt"
)

func partition(array []int, low, high int) int {
	pivot := array[high]
	i := low - 1

	for j := low; j < high; j++ {
		if array[j] <= pivot {
			i++
			array[i], array[j] = array[j], array[i]
		}
	}

	array[i+1], array[high] = array[high], array[i+1]
	return i + 1
}

func quickSort(array []int, low, high int) {
	if low < high {
		p := partition(array, low, high)
		quickSort(array, low, p-1)
		quickSort(array, p+1, high)
	}
}

func main() {
	data := []int{3, 6, 24, 46, 7, 2}
	fmt.Println("Unsorted array:", data)
	quickSort(data, 0, len(data)-1)
	fmt.Println("Sorted array:", data)
}
