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

func (h *Heap) Left(index int) int {
	if len(h.Data) == 0 {
		return 0
	}
	n := index * 2
	return h.Data[n]
}

func (h *Heap) Right(index int) int {
	if len(h.Data) == 0 {
		return 0
	}
	n := index*2 + 1
	return h.Data[n]
}
