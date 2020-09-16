package heap

/**
堆(Heap)是计算机科学中一类特殊的数据结构的统称。堆通常是一个可以被看做一棵完全二叉树的数组对象。
1、堆中某个节点的值总是不大于或不小于其父节点的值；
2、堆总是一棵完全二叉树。

将根节点最大的堆叫做最大堆或大根堆，根节点最小的堆叫做最小堆或小根堆。常见的堆有二叉堆、斐波那契堆等

堆是非线性数据结构，相当于一维数组，有两个直接后继。
堆的定义如下：n个元素的序列{k1,k2,ki,…,kn}当且仅当满足下关系时，称之为堆。
(ki <= k2i,ki <= k2i+1)或者(ki >= k2i,ki >= k2i+1), (i = 1,2,3,4...n/2)

最大堆（相应的可以定义最小堆）：堆中某个节点的值总是不大于其父节点的值

二叉树实现堆
*/

type Heap struct {
	Data []int
}

func (h *Heap) Size() int {
	return len(h.Data)
}

// 堆中最大元素
func (h *Heap) MaxHeap() int {
	if len(h.Data) == 0 {
		return 0
	}
	return h.Data[0]
}

func (h *Heap) Parent(index int) int {
	if len(h.Data) == 0 {
		return 0
	}
	n := index / 2
	return h.Data[n]
}

func (h *Heap) LeftChild(index int) int {
	if h.isEmpty() {
		return 0
	}
	n := index * 2
	return h.Data[n]
}

func (h *Heap) RightChild(index int) int {
	if len(h.Data) == 0 {
		return 0
	}
	n := index*2 + 1
	return h.Data[n]
}

func (h *Heap) isEmpty() bool {
	return len(h.Data) == 0
}

// 往堆中添加元素
func (h *Heap) Add(d int) {
	h.Data = append(h.Data, d)
	h.siftUp(len(h.Data))
}

// 上浮 检查父节点 是否替换位置
func (h *Heap) siftUp(k int) {
	for {
		if k > 0 && h.Parent(k) < h.Data[k] {
			h.Data[k/2], h.Data[k] = h.Data[k], h.Data[k/2]
			k = k / 2
		} else {
			break
		}
	}
}

// 取出堆中最大元素 （场景：实现堆的降序输出）
func (h *Heap) ExtractMax() int {
	// 最末端元素移至最大节点,删除末端元素
	if h.Size() == 0 {
		panic("heap is null")
	}
	max := h.MaxHeap()
	h.Data[0] = h.Data[h.Size()]
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
			if j+1 < h.Size() && h.Data[j+1] > h.Data[j] {
				j = h.RightChild(k)
			}

			// 比较父节点与子节点最大值比较，不符合堆，发生交换；反之，中断下沉
			if h.Data[k] > h.Data[j] {
				break
			}
			h.Data[k], h.Data[j] = h.Data[j], h.Data[k]
			k = j
		}
	}
}

// 取出最大元素，并且添加一个元素
func (h *Heap) Replace(d int) int {
	max := h.Data[0]
	h.Data[0] = d
	h.siftDown(0)
	return max
}

// 讲一个给定的二叉树生成heap
func (h *Heap) Heapify(data []int) {

}
