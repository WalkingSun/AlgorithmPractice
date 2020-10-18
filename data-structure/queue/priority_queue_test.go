package queue

import (
	"sort"
	"testing"
)

func TestPriorityQueue(t *testing.T) {
	tests := []struct {
		data []int
		want int
	}{
		0: {
			data: []int{1, 9, 3, 6, 7},
			want: 9,
		},
		1: {
			data: []int{3, 9, 3, 16, 5},
			want: 16,
		},
	}

	priorityQueue := NewPriorityQueue()
	for _, tt := range tests {
		for _, d := range tt.data {
			priorityQueue.Push(d)
		}
		if tt.want != priorityQueue.Pop() {
			t.Error("verify fail")
		}
	}

	sort.Sort()
}
