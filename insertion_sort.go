package main

import "fmt"

func insertionSort(arr []int) {
	for i := 1; i < len(arr); i++ {
		key := arr[i]
		j := i - 1
		for j >= 0 && arr[j] > key {
			arr[j+1] = arr[j]
			j--
		}
		arr[j+1] = key
	}
}

func main() {
	data := []int{34, 8, 64, 51, 32, 21}
	fmt.Println("Sebelum diurutkan:", data)

	insertionSort(data)
	fmt.Println("Setelah diurutkan:", data)
}
