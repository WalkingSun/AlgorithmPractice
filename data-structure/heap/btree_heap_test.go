package heap

import (
	"reflect"
	"testing"
)

func TestHeap_Heapify(t *testing.T) {
	tests := []struct {
		data []int
		want []int
	}{
		0: {
			data: []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
			want: []int{9, 8, 7, 4, 5, 6, 3, 2, 1},
		},
	}

	heap := NewHeap()
	for _, tt := range tests {
		r := heap.Heapify(tt.data)
		if !reflect.DeepEqual(r, tt.want) {
			t.Error("verify fail")
		}
	}
}

func TestHeap_Add(t *testing.T) {
	tests := []struct {
		data []int
		want []int
	}{
		0: {
			data: []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
			want: []int{9, 8, 7, 6, 4, 3, 2, 5, 1},
		},
	}

	heap := NewHeap()
	for _, tt := range tests {
		for _, d := range tt.data {
			heap.Add(d)
		}
		if !reflect.DeepEqual(heap.Data, tt.want) {
			t.Error("verify fail,data: ", heap.Data)
		}
	}
}

func TestHeap_Replace(t *testing.T) {

	tests := []struct {
		data    []int
		replace int
		want    []int
	}{
		0: {
			data:    []int{8, 7, 6, 4, 3, 2, 5, 1},
			replace: 9,
			want:    []int{9, 7, 6, 4, 3, 2, 5, 1},
		},
	}

	heap := NewHeap()
	for _, tt := range tests {
		heap.Data = tt.data
		heap.Replace(tt.replace)
		if !reflect.DeepEqual(heap.Data, tt.want) {
			t.Error("verify fail,data: ", heap.Data)
		}
	}
}
