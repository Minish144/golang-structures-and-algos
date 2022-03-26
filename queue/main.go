package main

import "fmt"

type Queue struct {
	values []int
	Size   int
}

func (q *Queue) Append(val int) {
	q.values = append(q.values, val)
	q.Size++
}

func (q *Queue) Pop() int {
	first := q.values[0]
	q.values = q.values[1:]
	return first
}

func (q *Queue) Traverse() {
	for index, val := range q.values {
		fmt.Printf("[%d]: %d\n", index, val)
	}
}

func main() {
	queue := Queue{make([]int, 0), 0}

	queue.Append(5)
	queue.Append(10)
	queue.Append(15)
	queue.Append(20)

	fmt.Println(queue.Pop())

	queue.Traverse()
}
