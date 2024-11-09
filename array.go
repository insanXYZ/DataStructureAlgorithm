package main

import "fmt"

func main() {
	data := []int{5, 4, 3, 2, 1}
	fmt.Println(FindTheLowestValue(data))
}

func FindTheLowestValue(arr []int) int {
	res := arr[0]

	for _, v := range arr {
		if v < res {
			res = v
		}
	}

	return res
}
