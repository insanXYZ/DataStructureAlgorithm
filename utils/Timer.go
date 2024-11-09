package utils

import (
	"fmt"
	"time"
)

func Timer(name string) {
	start := time.Now()
	fmt.Printf("%s took %v\n", name, time.Since(start))
}
