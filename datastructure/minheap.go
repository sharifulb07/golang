package main

import (
	"container/heap"
	"fmt"
)

// int slice type heap
type intHeap []int

func (h intHeap) Len() int { return len(h) }

func (h intHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h intHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *intHeap) Push(x any) {
	*h = append(*h, x.(int))
}

func (h *intHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]

	return x
}

func main() {

	h := &intHeap{}

	heap.Init(h)

	heap.Push(h, 10)
	heap.Push(h, 14)
	heap.Push(h, 15)
	heap.Push(h, 11)
	heap.Push(h, 25)

	fmt.Println(heap.Pop(h))
	fmt.Println(heap.Pop(h))

}
