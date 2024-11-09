package main

import (
	"time"
)

func main() {
	// wg := &sync.WaitGroup{}
	start := time.Now().Unix()

	for i := 0; i < 1000000; i++ {
		// go func() {
		// 	wg.Add(1)
		// 	defer wg.Done()
		// fmt.Println(i)
		// }()
	}

	// wg.Wait()
	// final := time.Now().Unix() - start

}
