package main

import (
	"container/list"
	"fmt"
)

type LruList struct {
	lst          *list.List
	mp           map[int]*list.Element
	LruMaxLength int
}

func NewLruList() *LruList {
	lst := list.New()
	mp := map[int]*list.Element{}
	for i := 0; i < 9; i++ {
		ele := lst.PushBack(i)
		mp[i] = ele
	}
	return &LruList{lst: lst, mp: mp, LruMaxLength: 10}
}

func (lrulist LruList) TraversList() {
	lst := lrulist.lst
	head := lst.Front()
	for head.Next() != nil {
		fmt.Printf("%v ", head.Value)
		head = head.Next()
	}
	fmt.Println(head.Value)
}

func (lrulist *LruList) visitList(value int) {
	lst := lrulist.lst
	mp := lrulist.mp

	head := lst.Front()
	tail := lst.Back()
	length := lst.Len()

	var existValue *list.Element

	// 找到命中的缓存
	if value, exists := mp[value]; exists {
		existValue = value
	}

	// 如果命中缓存，从链表中找到对应的key，移到链表头部
	if existValue != nil {
		if existValue.Value == head.Value{
			return
		}
		lst.MoveToFront(existValue)
	} else {
		var ele *list.Element
		if length+1 > lrulist.LruMaxLength {
			lst.Remove(tail)
			delete(mp, tail.Value.(int))
			ele = lst.PushFront(value)
		} else {
			ele = lst.PushFront(value)
		}
		mp[value] = ele
	}

	lrulist.lst = lst
	lrulist.mp = mp
}

func testList() {
	lrulst := NewLruList()
	lrulst.TraversList()
	lrulst.visitList(5)
	lrulst.TraversList()
	lrulst.visitList(2)
	lrulst.TraversList()
	lrulst.visitList(4)
	lrulst.TraversList()
	lrulst.visitList(10)
	lrulst.TraversList()
	lrulst.visitList(11)
	lrulst.TraversList()
}

func main() {
	testList()
}
