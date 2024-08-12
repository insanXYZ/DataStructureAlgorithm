package main

import "fmt"

type Queue struct {
	Arrays []int
}

func (q *Queue) enqueue(i int) {
	q.Arrays = append(q.Arrays, i)
}

func (q *Queue) dequeue() int {
	if !q.isEmpty() {
		c := q.Arrays[0]
		q.Arrays = q.Arrays[1:]
		return c
	}
	return 0
}

func (q *Queue) isEmpty() bool {
	return len(q.Arrays) == 0
}

func main() {
	q := Queue{}
	q.enqueue(1)
	q.enqueue(2)
	fmt.Println(q.dequeue())
}
