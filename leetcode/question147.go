package main

import "fmt"

/**
https://leetcode-cn.com/problems/top-k-frequent-elements/solution/qian-k-ge-gao-pin-yuan-su-by-leetcode-solution/


n个元素中获取m个频次最高的元素值，算法复杂度高于nlog(n)

方案：
1、 最小堆 nlog(m)
- 对元素值出现次数按map计数；
- m个元素入最小堆；
- 超过m的元素，如果比堆顶元素还要大，推出堆顶元素，插入当前元素；
*/

func main() {
	k := 2
	nums := []int{5, 3, 1, 1, 1, 3, 73, 1}
	numDict := getDict(nums)
	h := NewHeap()
	i := 0
	for num, freq := range numDict {
		if i < k {
			h.Add(freq, num)
		} else {
			if h.MinHeap() < freq {
				h.Replace(freq, num)
			}
		}
		i++
	}

	ret := make([]int, 0, k)
	for _, d := range h.Data {
		ret = append(ret, d.Num)
	}
	fmt.Println(ret)
}

func getDict(nums []int) map[int]int {
	ret := make(map[int]int)
	for _, num := range nums {
		ret[num]++
	}
	return ret
}

func NewHeap() *Heap {
	return &Heap{}
}

type Heap struct {
	Data []struct {
		Num  int
		Freq int
	}
}

func (h *Heap) Size() int {
	return len(h.Data)
}

func (h *Heap) MinHeap() int {
	if len(h.Data) == 0 {
		return 0
	}
	return h.Data[0].Freq
}

func (h *Heap) Parent(index int) int {
	if len(h.Data) == 0 {
		return 0
	}
	n := index / 2
	return h.Data[n].Freq
}

func (h *Heap) LeftChild(index int) int {
	if h.isEmpty() {
		return 0
	}
	n := (index+1)*2 - 1
	return n
}

func (h *Heap) RightChild(index int) int {
	if len(h.Data) == 0 {
		return 0
	}
	n := (index + 1) * 2
	return n
}

func (h *Heap) isEmpty() bool {
	return len(h.Data) == 0
}

// 往堆中添加元素
func (h *Heap) Add(d int, num int) {
	h.Data = append(h.Data, struct {
		Num  int
		Freq int
	}{Num: num, Freq: d})
	h.siftUp(len(h.Data) - 1)
}

// 上浮 检查父节点 是否替换位置
func (h *Heap) siftUp(k int) {
	for {
		if k > 0 && h.Parent(k) > h.Data[k].Freq {
			h.Data[k/2], h.Data[k] = h.Data[k], h.Data[k/2]
			k = k / 2
		} else {
			break
		}
	}
}

// 取出堆中最小元素
func (h *Heap) ExtractMin() int {
	// 最末端元素移至最大节点,删除末端元素
	if h.Size() == 0 {
		panic("heap is null")
	}

	max := h.MinHeap()
	h.Data[0] = h.Data[h.Size()-1]
	h.Data = h.Data[:h.Size()-1]
	// 下沉 节点值与左右节点最大值比较，不符合堆替换；下称 一直到满足堆条件；
	h.siftDown(0)
	return max
}

func (h *Heap) siftDown(k int) {
	for {
		// 节点有child
		if h.LeftChild(k) < h.Size() {

			// 找出child中最大节点j
			j := h.LeftChild(k)
			if j+1 < h.Size() && h.Data[j+1].Freq < h.Data[j].Freq {
				j = h.RightChild(k)
			}

			// 比较父节点与子节点最大值比较，不符合堆，发生交换；反之，中断下沉
			if h.Data[k].Freq < h.Data[j].Freq {
				break
			}
			h.Data[k], h.Data[j] = h.Data[j], h.Data[k]
			k = j
		} else {
			break
		}
	}
}

// 取出最小元素，并且添加一个元素
func (h *Heap) Replace(d int, num int) int {
	min := h.Data[0].Freq
	h.Data[0] = struct {
		Num  int
		Freq int
	}{Num: num, Freq: d}
	h.siftDown(0)
	return min
}
