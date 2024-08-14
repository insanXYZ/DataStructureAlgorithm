package main

import (
	"fmt"
)

func main() {
	l := 5

	var values []int
	for i := l; i > 0; i-- {
		values = append(values, i)
	}

	fmt.Println("before =", values)
	var count int

	for i := 0; i < len(values)-1; i++ {
		minI := i
		for j := i + 1; j < len(values); j++ {
			if values[j] < values[minI] {
				minI = j
			}
		}
		values[i], values[minI] = values[minI], values[i]
	}

	fmt.Println("after =", values)
	fmt.Println(count)

}
