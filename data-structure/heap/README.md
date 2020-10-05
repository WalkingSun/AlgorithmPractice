# heap
堆(Heap)是计算机科学中一类特殊的数据结构的统称。堆通常是一个可以被看做一棵完全二叉树的数组对象。特点：
- 堆中某个节点的值总是不大于或不小于其父节点的值；
- 堆总是一棵完全二叉树。

将根节点最大的堆叫做最大堆或大根堆，根节点最小的堆叫做最小堆或小根堆。常见的堆有二叉堆、斐波那契堆等

堆是非线性数据结构，相当于一维数组，有两个直接后继。

堆的定义如下：n个元素的序列{k1,k2,ki,…,kn}当且仅当满足下关系时，称之为堆。
(ki <= k2i,ki <= k2i+1)或者(ki >= k2i,ki >= k2i+1), (i = 1,2,3,4...n/2)

若将和此次序列对应的一维数组（即以一维数组作此序列的存储结构）看成是一个完全二叉树，则堆的含义表明，完全二叉树中所有非终端结点的值均不大于（或不小于）其左、右孩子结点的值。由此，若序列{k1,k2,…,kn}是堆，则堆顶元素（或完全二叉树的根）必为序列中n个元素的最小值（或最大值）。

## 数组实现堆

![image](https://raw.githubusercontent.com/WalkingSun/WindBlog/gh-pages/images/blog/image-20200727.png)

### 堆中添加元素

新增元素52，与父亲节点比较，不符合堆条件，替换位置；父亲节点与其父节点不符合继续替换，一直到符合条件位置。(sift up)

![image-20200906211231547](https://raw.githubusercontent.com/WalkingSun/WindBlog/gh-pages/images/ws2/image-20200906211231547.png)

### 从heap中取出元素

### relpace
取出最大元素后，放入一个新元素

实现1：先extralMax，再add，两次O（log n）操作；

实现2：直接替换堆顶元素，执行sift down，一次O(logN)操作；

### Heapify
给定一个数组，转化成堆。

实现1：将n个元素逐个插入到一个空堆中 O(nlogN)

实现2：Heapify 将任意数组整理成堆的形状 O(n)

- 当前数组看作为二叉树；
- 向前遍历，每个节点执行sift down（下沉）；
- 跳过没有子节点的节点；

![image-20201005220055079](https://raw.githubusercontent.com/WalkingSun/WindBlog/gh-pages/images/ws2/image-20201005220055079.png)

