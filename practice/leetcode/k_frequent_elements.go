package practice

import "container/heap"

// https://leetcode.com/problems/top-k-frequent-elements/

func topKFrequent(nums []int, k int) []int {
	dict := make(map[int]int)

	for _, v := range nums {
		if _, ok := dict[v]; ok {
			dict[v] = dict[v] + 1
		} else {
			dict[v] = 1
		}
	}

	pq := make(PriorityQueue, len(dict))

	i := 0
	for value, priority := range dict {
		pq[i] = &Item{
			value:    value,
			priority: priority,
			index:    i,
		}
		i++
	}
	heap.Init(&pq)

	ans := make([]int, 0)
	//j := 0

	for pq.Len() > 0 && k > 0 {
		item := heap.Pop(&pq).(*Item)
		ans = append(ans, item.value)
		// ans[j] = item.value
		// j++
		k--
	}
	return ans
}

type Item struct {
	value    int // The value of the item; arbitrary.
	priority int // The priority of the item in the queue.
	// The index is needed by update and is maintained by the heap.Interface methods.
	index int // The index of the item in the heap.
}

type PriorityQueue []*Item

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	// We want Pop to give us the highest, not lowest, priority so we use greater than here.
	return pq[i].priority > pq[j].priority
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	item := x.(*Item)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil  // avoid memory leak
	item.index = -1 // for safety
	*pq = old[0 : n-1]
	return item
}
