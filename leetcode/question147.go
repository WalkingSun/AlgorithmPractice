package main

/**
https://leetcode-cn.com/problems/top-k-frequent-elements/solution/qian-k-ge-gao-pin-yuan-su-by-leetcode-solution/


n个元素中获取m个频次最高的元素值，算法复杂度高于nlog(n)

方案：
1、 最小堆 nlog(m)
- 对元素值出现次数按map计数；
- m个元素入最小堆；
- 超过m的元素，如果比堆顶元素还要大，推出堆顶元素，插入当前元素；
*/
