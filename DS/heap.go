package DS

import (
	"fmt"
)

type HeapNode struct {
	Arr, Key []int64
	Capacity int64 //size of heap.
	Count    int64 //number of elements.
	HeapType int64
}

func (node *HeapNode) Parent(i int64) int64 {

	pos := (i - 1) / 2
	if (i-1) >= 0 && pos < int64(len(node.Arr)) {
		return pos
	}
	return -1
}

func (node *HeapNode) Child(i int64) (int64, int64) {
	firstChild := 2*i + 1
	secondChild := 2*i + 2

	if firstChild < node.Count && secondChild < node.Count {
		return firstChild, secondChild
	}
	if firstChild < node.Count && secondChild >= node.Count {
		return firstChild, -1
	}
	if firstChild >= node.Count && secondChild < node.Count {
		return -1, secondChild
	}

	return -1, -1

}

func (node *HeapNode) GetMaxOrMin() int64 {
	if node.Count > 0 {
		return node.Arr[0]
	}
	return -1
}

func (node *HeapNode) PercolateDown(i int64) {

	if i < 0 {
		return
	}

	var (
		l, r, max int64
	)

	l, r = node.Child(i)
	if l != -1 && node.Arr[l] > node.Arr[i] {
		max = l
	} else {
		max = i
	}
	if r != -1 && node.Arr[r] > node.Arr[max] {
		max = r
	}

	if max != i {
		temp := node.Arr[max]
		node.Arr[max] = node.Arr[i]
		node.Arr[i] = temp
	}

	if max == i {
		return
	}

	node.PercolateDown(max)

}

func (node *HeapNode) PercolateDownMin(i int64) {

	if i < 0 {
		return
	}

	var (
		l, r, min int64
	)

	l, r = node.Child(i)
	if l != -1 && node.Arr[l] < node.Arr[i] {
		min = l
	} else {
		min = i
	}
	if r != -1 && node.Arr[r] < node.Arr[min] {
		min = r
	}

	if min != i {
		temp := node.Arr[min]
		node.Arr[min] = node.Arr[i]
		node.Arr[i] = temp
	}

	if min == i {
		return
	}

	node.PercolateDown(min)

}

func (node *HeapNode) PercolateDownKey(i int64) {

	if i < 0 {
		return
	}

	var (
		l, r, max int64
	)

	l, r = node.Child(i)
	if l != -1 && node.Key[l] > node.Key[i] {
		max = l
	} else {
		max = i
	}
	if r != -1 && node.Key[r] > node.Key[max] {
		max = r
	}

	if max != i {
		temp := node.Key[max]
		node.Key[max] = node.Key[i]
		node.Key[i] = temp

		temp = node.Arr[max]
		node.Arr[max] = node.Arr[i]
		node.Arr[i] = temp
	}

	if max == i {
		return
	}

	node.PercolateDown(max)

}

func (node *HeapNode) Delete() int64 {
	if node == nil {
		return -1
	}
	if node.Count == 0 {
		return -1
	}
	removedElement := node.Arr[0]
	node.Arr[0] = node.Arr[node.Count-1]
	node.Arr[node.Count-1] = 0
	node.Count -= 1
	node.PercolateDown(0)
	return removedElement
}

func (node *HeapNode) DeleteMin() int64 {
	if node == nil {
		return -1
	}
	if node.Count == 0 {
		return -1
	}
	removedElement := node.Arr[0]
	node.Arr[0] = node.Arr[node.Count-1]
	node.Arr[node.Count-1] = 0
	node.Count -= 1
	node.PercolateDownMin(0)
	return removedElement
}

func (node *HeapNode) DeleteKey() (int64, int64) {
	if node == nil {
		return 0, -1
	}
	if node.Count == 0 {
		return 0, -1
	}
	removedElement := node.Arr[0]
	removedKey := node.Key[0]
	node.Arr[0] = node.Arr[node.Count-1]
	node.Key[0] = node.Key[node.Count-1]
	node.Arr[node.Count-1] = 0
	node.Key[node.Count-1] = 0
	node.Count -= 1

	node.PercolateDownKey(0)

	return removedKey, removedElement
}

func (node *HeapNode) Insert(data int64) {
	if node.Count == node.Capacity {
		node.ReSize(2 * node.Capacity)
	}
	node.Arr[node.Count] = data
	node.Count++
	node.PercolateUp(node.Count - 1)
}

func (node *HeapNode) InsertMin(data int64) {
	if node.Count == node.Capacity {
		node.ReSize(2 * node.Capacity)
	}
	node.Arr[node.Count] = data
	node.Count++
	node.PercolateUpMin(node.Count - 1)
}

//key value insert
func (node *HeapNode) InsertKey(key, data int64) bool {
	if node.Count == node.Capacity {
		return false
	}
	node.Key[node.Count] = key
	node.Arr[node.Count] = data
	node.Count++
	node.PercolateUpKey(node.Count - 1)

	return true
}

func (node *HeapNode) ReSize(num int64) {
	var newArr []int64
	newArr = make([]int64, num)

	for idx := range node.Arr {
		newArr[idx] = node.Arr[idx]
	}
	node.Capacity = num
	node.Arr = newArr
}

func (node *HeapNode) ReSizeKey(num int64) {
	var newArr []int64
	newArr = make([]int64, num)
	newKey := make([]int64, num)

	for idx := range node.Key {
		newArr[idx] = node.Arr[idx]
		newKey[idx] = node.Key[idx]
	}
	node.Capacity = num
	node.Arr = newArr
	node.Key = newKey
}

func (node *HeapNode) PercolateUp(i int64) {
	var (
		p int64
	)
	p = node.Parent(i)
	for p >= 0 {
		if node.Arr[p] < node.Arr[i] {
			temp := node.Arr[p]
			node.Arr[p] = node.Arr[i]
			node.Arr[i] = temp
		} else {
			break
		}
		i = p
		p = node.Parent(p)

	}
}

func (node *HeapNode) PercolateUpMin(i int64) {
	var (
		p int64
	)
	p = node.Parent(i)
	for p >= 0 {
		if node.Arr[p] > node.Arr[i] {
			temp := node.Arr[p]
			node.Arr[p] = node.Arr[i]
			node.Arr[i] = temp
		} else {
			break
		}
		i = p
		p = node.Parent(p)

	}
}

func (node *HeapNode) PercolateUpKey(i int64) {
	var (
		p int64
	)
	p = node.Parent(i)
	for p >= 0 {
		if node.Key[p] < node.Key[i] {
			temp := node.Key[p]
			node.Key[p] = node.Key[i]
			node.Key[i] = temp

			temp = node.Arr[p]
			node.Arr[p] = node.Arr[i]
			node.Arr[i] = temp
		} else {
			break
		}
		i = p
		p = node.Parent(p)

	}
}

func (node *HeapNode) DestroyHeap() *HeapNode {
	return nil
}

// time complexity O(nlogn)
// Leaf node always satisfies the heap property but non leaf node does not satisfies the heap property.
func (node *HeapNode) BuildHeap(arr []int64, num int64) {
	if node == nil {
		return
	}
	var idx int64
	if node.Capacity < num {
		node.ReSize(num)
	}
	for idx := range arr {
		node.Arr[idx] = arr[idx]
	}
	node.Count = num

	for idx = (num - 1) / 2; idx >= 0; idx-- {
		node.PercolateDown(idx - 1)
	}
}

func (node *HeapNode) BuildHeapMin(arr []int64, num int64) {
	if node == nil {
		return
	}
	var idx int64
	if node.Capacity < num {
		node.ReSize(num)
	}
	for idx := range arr {
		node.Arr[idx] = arr[idx]
	}
	node.Count = num

	for idx = (num - 1) / 2; idx >= 0; idx-- {
		node.PercolateDown(idx - 1)
	}
}

//Heap Sort ---------Time complexity : O(nlogn)
func HeapSort(arr []int64, num int64) []int64 {
	var (
		result []int64
		idx    int64
	)
	heap := InitHeap(num)
	heap.BuildHeap(arr, num)

	for idx = num - 1; idx >= 0; idx-- {
		result = append(result, heap.Arr[0])
		heap.Arr[0] = heap.Arr[idx]
		heap.Count -= 1
		heap.PercolateDown(0)
	}

	return result
}

//max nodes in heap will be 2^(h+1) - 1
//min nodes will be 2^h
func (node *HeapNode) FindMaxElementInHeap() int64 {
	if node == nil {
		return 0
	}
	var idx, max int64
	for idx = (node.Count + 1) / 2; idx < node.Count; idx++ {
		if max < node.Arr[idx] {
			max = node.Arr[idx]
		}
	}
	return max
}

//time- O(logn)
func (node *HeapNode) DeleteNode(idx int64) int64 {
	if node == nil {
		return -1
	}
	if idx > node.Count {
		return -1
	}
	key := node.Arr[idx]
	node.Arr[idx] = node.Arr[node.Count-1]
	node.Arr[node.Count-1] = 0
	node.Count -= 1
	node.PercolateDown(idx)
	return key
}

// time - O(kLogn)
func (node *HeapNode) KthSmallestElement(k int) int64 {
	if node == nil {
		return -1
	}
	var temp int64
	count := int(node.Count)
	for idx := 0; idx <= count-k; idx++ {
		temp = node.Delete()
	}

	return temp
}

/*
->Using Quicksort
->Using orig min heap and aux min heap. with k loops we will remove elements from origianl heap and add it
  to the aux heap. and heapify and than remove the max element from aux heap and add the left and right child
  of the max element.
*/
// time - O(k + (n-k)Logk)  .... Under dev.
func KthSmallestElementV2(arr []int64, k int) int64 {
	hp := InitHeap(int64(len(arr)))
	hp.BuildHeap(arr[:k+1], int64(k))

	var (
		idx  int
		temp int64
	)
	for idx = k; idx < len(arr); idx++ {
		fmt.Println(hp.GetMaxOrMin(), arr[idx])
		if hp.GetMaxOrMin() <= arr[idx] {
			temp = hp.GetMaxOrMin()
		} else {
			hp.Arr[0] = arr[idx]
			hp.BuildHeap(hp.Arr, int64(k))
		}
	}

	return temp
}

func (heap *HeapNode) MergeKList(list [][]int64, i, j, num int64) []int64 {
	if heap == nil {
		return nil
	}
	if num > heap.Capacity {
		heap.ReSize(num)
	}

	var idx, minVal int64
	var index []int64
	res := make([]int64, 0)
	minVal = -1
	for idx = 0; idx < i; idx++ {
		heap.InsertKey(list[idx][i-1], idx)
		index = append(index, i-2)
	}

	for idx = 0; idx < i*j; idx++ {
		key, data := heap.DeleteKey()
		res = append(res, key)
		if index[data] >= 0 {
			heap.InsertKey(list[data][index[data]], data)
		} else {
			heap.InsertKey(minVal, 0)
		}
		index[data]--
	}
	return res
}

func FindMedian(arr []int64, num int64) int64 {
	if arr == nil {
		return 0
	}
	var median int64
	leftMax := InitHeap(num)
	rightMin := InitHeap(num)
	for idx := range arr {
		median = findMedian(arr[idx], median, leftMax, rightMin)
	}

	return median
}

func findMedian(element int64, median int64, leftMax, rightMin *HeapNode) int64 {
	if leftMax == nil || rightMin == nil {
		return 0
	}
	var res int64
	sign := signum(leftMax, rightMin)
	switch sign {
	//case left tree
	case 1:
		if element < median {
			rightMin.InsertMin(leftMax.Delete())
			leftMax.Insert(element)
		} else {
			rightMin.InsertMin(element)
		}
		res = (rightMin.GetMaxOrMin() + leftMax.GetMaxOrMin()) / 2
	case 0:
		if element < median {
			leftMax.Insert(element)
			res = leftMax.GetMaxOrMin()
		} else {
			rightMin.InsertMin(element)
			res = rightMin.GetMaxOrMin()
		}
		//case right tree
	case -1:
		if element < median {
			leftMax.Insert(element)
		} else {
			leftMax.Insert(rightMin.DeleteMin())
			rightMin.InsertMin(element)
		}
		res = (rightMin.GetMaxOrMin() + leftMax.GetMaxOrMin()) / 2
	}
	return res
}

func signum(leftMax, rightMin *HeapNode) int64 {
	return leftMax.Count - rightMin.Count
}

func InitHeap(capacity int64) *HeapNode {
	arr := make([]int64, capacity)
	key := make([]int64, capacity)

	return &HeapNode{
		Capacity: capacity,
		Arr:      arr,
		Key:      key,
		HeapType: 1, // 1 := max heap type
	}
}

/* Points to remember for heaps.
-> ShiftUp is expensive operation than shift down because of half the heap is present at bottom.
-> Insert an element take O(Logn) and inserting n elements take O(nLogn)
-> We can build heap with shiftDown operation in O(n)
-> Delete operation takes up to O(nLogn)
-> We can merge two heaps with size m,n and the complexity will be O(Logm * Logn) if m=n, O(log^2n) using fibonacci heap
-> Three types of heap, Binary heap, binomial heap and fibonacci heap.
*/
