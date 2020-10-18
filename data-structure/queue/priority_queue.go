package queue

import (
	"../heap"
)

var (
	newHeap = heap.NewHeap()
)

func NewPriorityQueue() Queue {
	return &priorityQueue{}
}

type priorityQueue struct {
}

func (p *priorityQueue) GetSize() int {
	return newHeap.Size()
}

func (p *priorityQueue) Push(d int) {
	newHeap.Add(d)
}

func (p *priorityQueue) Pop() int {
	return newHeap.ExtractMax()
}

func (p *priorityQueue) IsEmpty() bool {
	if p.GetSize() != 0 {
		return false
	}
	return true
}
