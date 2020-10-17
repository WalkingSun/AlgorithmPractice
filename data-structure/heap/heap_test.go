package heap

import (
	"fmt"
	"testing"
)

func TestHeap_Heapify(t *testing.T) {
	data := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	heap := NewHeap()
	r := heap.Heapify(data)
	fmt.Println(r)
}
