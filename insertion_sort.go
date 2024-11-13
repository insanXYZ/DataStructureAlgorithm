package main

import "fmt"

func main() {
	data := []int{22, 44, 11, 66, 10, 99, 1}
	InsertionSort(data)
}

func InsertionSort(arr []int) {
	for i := 0; i < len(arr); i++ {
		for j := i; j > 0 && arr[j-1] > arr[j]; j-- {
			arr[j], arr[j-1] = arr[j-1], arr[j]
		}
		fmt.Println("index ", i, " array ", arr)
	}
}
