package main

import "fmt"

func main() {
	// fibonacciLoop(10)
	// fibonacciRecursi(0, 1, 10, 0)
	// fmt.Println(fibonacciFormula(10))
}

func fibonacciLoop(r int) {

	a, b := 0, 1

	for i := 0; i < r; i++ {
		s := a + b

		fmt.Println(a)
		a = b
		b = s
	}

}

func fibonacciRecursi(a, b, r, c int) {
	if c < r {
		fmt.Println(a)
		s := a + b
		a = b
		b = s
		c++

		fibonacciRecursi(a, b, r, c)

	} else {
		return
	}
}

func fibonacciFormula(n int) int {
	if n <= 1 {
		return n
	}
	return fibonacciFormula(n-1) + fibonacciFormula(n-2)

}
