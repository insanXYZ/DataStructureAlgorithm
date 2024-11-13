package main

import "fmt"

func main() {
	arr := []int{5, 4, 3, 2, 1}
	SelectionSort(arr)
	fmt.Println(arr)
}

func SelectionSort(arr []int) {
	for i := 0; i < len(arr); i++ {
		min_index := i
		for j := 0 + i; j < len(arr); j++ {
			if arr[j] < arr[min_index] {
				min_index = j
			}
		}
		arr[i], arr[min_index] = arr[min_index], arr[i]
	}
}
