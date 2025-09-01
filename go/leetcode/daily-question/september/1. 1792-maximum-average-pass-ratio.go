package september

import (
	"container/heap"
	"math"
)

type Class struct {
	pass  int
	total int
	gain  float64
}

type MaxHeap []Class

func (h MaxHeap) Len() int           { return len(h) }
func (h MaxHeap) Less(i, j int) bool { return h[i].gain > h[j].gain }
func (h MaxHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MaxHeap) Push(x interface{}) {
	*h = append(*h, x.(Class))
}

func (h *MaxHeap) Pop() interface{} {
	old := *h
	n := len(old)
	item := old[n-1]
	*h = old[:n-1]
	return item
}

func calcGain(pass, total int) float64 {
	return float64(pass+1)/float64(total+1) - float64(pass)/float64(total)
}

func MaxAverageRatio(classes [][]int, extraStudents int) float64 {
	h := &MaxHeap{}
	heap.Init(h)

	for _, c := range classes {
		pass, total := c[0], c[1]
		heap.Push(h, Class{pass, total, calcGain(pass, total)})
	}

	for i := 0; i < extraStudents; i++ {
		class := heap.Pop(h).(Class)
		newPass, newTotal := class.pass+1, class.total+1
		heap.Push(h, Class{newPass, newTotal, calcGain(newPass, newTotal)})
	}

	sum := 0.0
	for _, c := range *h {
		sum += float64(c.pass) / float64(c.total)
	}
	average := sum / float64(len(*h))

	return math.Round(average*1e5) / 1e5
}
