package main

import (
	"container/heap"
	"fmt"
	"strconv"
	"time"
)

const LIFE = 2

var cache map[int]*Node

type Node struct {
	deadline int64
	key      int
	info     string
}

type TimeoutHeap []*Node

func (pq TimeoutHeap) Len() int {
	return len(pq)
}

func (pq TimeoutHeap) Less(i, j int) bool {
	return pq[i].deadline < pq[j].deadline
}

func (pq TimeoutHeap) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *TimeoutHeap) Push(x interface{}) {
	item := x.(*Node)
	*pq = append(*pq, item)
}

func (pq *TimeoutHeap) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[0 : n-1]
	return item
}

func testTimeoutHeap() {
	pq := make(TimeoutHeap, 0, 5)
	heap.Init(&pq)
	var timedelta int64

	for i := 0; i < 10; i++ {
		timedelta = int64(LIFE * (i + 1))
		node := Node{deadline: time.Now().Unix() + timedelta, key: i, info: "note:" + strconv.Itoa(i)}
		cache[i] = &node
		heap.Push(&pq, &node)
	}

	tm := time.NewTicker(1 * time.Second)
	for {
		<-tm.C
		for {
			currentTS := time.Now().Unix()
			if pq.Len() <= 0 {
				break
			}
			first := pq[0]
			if currentTS < first.deadline {
				break
			} else {
				delete(cache, first.key)
				heap.Pop(&pq)
				fmt.Printf("delete %v\n", *first)
			}
		}
		if pq.Len() <= 0 {
			break
		}
	}
}

func init() {
	cache = make(map[int]*Node)
}

func main() {
	testTimeoutHeap()
}
