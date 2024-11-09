package main

import (
	"data_structures_algorithm/utils"
	"fmt"
)

func main() {
	data := []int{5, 4, 3, 2, 1}
	BubbleSortLowestToHigh(data)
	fmt.Println(data)

	BubbleSortHighestToLow(data)
	fmt.Println(data)

	BubbleSortImprovement(data)
	fmt.Println(data)
}

func BubbleSortLowestToHigh(arr []int) {
	defer utils.Timer("Bubble Sort Lowest To High")
	l := len(arr)

	for i := 0; i < (l - 1); i++ {
		for j := 0; j < (l - i - 1); j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
}

func BubbleSortHighestToLow(arr []int) {
	defer utils.Timer("Bubble Sort Highest To Low")
	l := len(arr)

	for i := 0; i < (l - 1); i++ {
		for j := 0; j < (l - i - 1); j++ {
			if arr[j] < arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
}

// in this function like Lowest to high, but adding brake expression
// you can see the explanation on https://www.w3schools.com/dsa/dsa_algo_bubblesort.php
func BubbleSortImprovement(arr []int) {
	defer utils.Timer("Bubble Sort Improvement")
	l := len(arr)

	for i := 0; i < (l - 1); i++ {
		Swapped := false
		for j := 0; j < (l - i - 1); j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
				Swapped = true
			}
		}

		if !Swapped {
			break
		}
	}

}
