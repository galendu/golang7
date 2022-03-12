package main

import (
	"fmt"
	"time"
	"container/heap"
)

// 3 用map和堆 表实现超时缓存
var cache map[int]*Node

const LIFE = 10

func init() {
	cache = make(map[int]*Node)
}

type Node struct {
	Info        string
	Deadline    int64
	Key         int
}

type TimeOutHeap []*Node

func (pq TimeOutHeap) Len() int {
	return len(pq)
}

func (pq TimeOutHeap) Less(i, j int) bool {
	return pq[i].Deadline < pq[j].Deadline
}

func (pq TimeOutHeap) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *TimeOutHeap) Push(x interface{}) {
	item := x.(*Node)
	*pq = append(*pq, item)
}

func (pq *TimeOutHeap) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]     // 数组最后一个元素
	*pq = old[0 : n-1]   // 去掉最后一个元素
	return item
}

func testTimeOutCache() {
	pq := make(TimeOutHeap, 0, 5)
	heap.Init(&pq)    // 从无序状态构建堆

	for i :=0; i < 10; i++ {
		node := &Node{Info: "54t43t45", Deadline:time.Now().UnixNano() + LIFE, Key: i}
		cache[i] = node       // 放入缓存
		heap.Push(&pq, node)  // 同时放入堆
		time.Sleep(20 * time.Millisecond)
	}

	ticker := time.NewTicker(5 * time.Millisecond) // 每隔5毫秒，检查一下是否有元素到期
	for {
		<-ticker.C
		for {
			currentTimestamp := time.Now().UnixNano()  // 取消当前时间
			if pq.Len() <= 0 {
				break
			}
			first := pq[0] // 取得小根堆顶元素
			if currentTimestamp < first.Deadline {
				break
			} else {  // 当前时间比堆顶元素小，说明堆顶已经到期，需要从缓存中删除
				delete(cache, first.Key)
				heap.Pop(&pq)
				fmt.Printf("delete %v\n", *first)
			}
		}
	}
}

func main() {
	testTimeOutCache()
}