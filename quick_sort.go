package main

import (
	"data_structures_algorithm/utils"
	"fmt"
)

func main() {
	defer utils.Timer("quick sort ")
	data := []int{64, 34, 25, 12, 22, 11, 90, 5}
	QuickSort(data, 0, len(data)-1)
	fmt.Println(data)
}

func QuickSort(arr []int, low, high int) {
	if low < high {
		pivot_index := Partition(arr, low, high)
		fmt.Println("left")
		QuickSort(arr, low, pivot_index-1)
		fmt.Println("right")
		QuickSort(arr, pivot_index+1, high)
	}
}

func Partition(arr []int, low, high int) int {
	pivot := arr[high]
	i := low - 1

	for j := low; j < high; j++ {
		if arr[j] <= pivot {
			i += 1
			arr[i], arr[j] = arr[j], arr[i]
			fmt.Println(arr)
		}
	}

	arr[i+1], arr[high] = arr[high], arr[i+1]
	fmt.Println(arr)
	return i + 1
}
