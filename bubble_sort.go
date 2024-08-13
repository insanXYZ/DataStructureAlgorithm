package main

import "fmt"

func main() {
	l := 300

	var values []int
	for i := l; i > 0; i-- {
		values = append(values, i)
	}

	var count int

	// for i := 1; i < l; i++ {
	// 	for j := 0; j < l-i; j++ {
	// 		count++
	// 		if values[j] > values[j+1] {
	// 			values[j], values[j+1] = values[j+1], values[j]
	// 			count++
	// 		}
	// 	}
	// 	count++
	// }

	for i := range l - 1 {
		for j := range l - i - 1 {
			count++
			if values[j] > values[j+1] {
				values[j], values[j+1] = values[j+1], values[j]
				count++
			}
		}
		count++
	}

	fmt.Println(values)
	fmt.Println("Operations =", count)

}
