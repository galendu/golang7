package main

import (
	"container/heap"
	"container/list"
	"crypto"
	"crypto/rsa"
	"crypto/sha1"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"os"
)

//1.数字签名
func DigitalSignature(s string) []byte {
	sha1 := sha1.New()
	sha1.Write([]byte(s))
	digest := sha1.Sum([]byte(""))
	privateKey, _ := os.ReadFile("contr11/privatekey.pem")
	block, _ := pem.Decode(privateKey)
	priv, _ := x509.ParsePKCS1PrivateKey(block.Bytes)
	signature, _ := rsa.SignPKCS1v15(nil, priv, crypto.Hash(0), digest)
	return signature
}
func verifySignature(s string, signature []byte) bool {
	sha1 := sha1.New()
	sha1.Write([]byte(s))
	digest := sha1.Sum([]byte(""))

	publicKey, _ := os.ReadFile("contr11/publickey.pem")
	block, _ := pem.Decode(publicKey)
	pubInterface, _ := x509.ParsePKIXPublicKey(block.Bytes)
	pub := pubInterface.(*rsa.PublicKey)

	return rsa.VerifyPKCS1v15(pub, crypto.Hash(0), digest, signature) == nil
}

//2.go包中list实现LRU
//1-->2-->3-->4   3-->1-->2-->4
func visit(lst *list.List, s int) {
	if lst.Len() == 0 {
		fmt.Println("链表没有元素")
	}
	head := lst.Front()
	for {
		if head == nil {
			break
		}
		if head.Value.(int) != s {
			head = head.Next()
		} else {
			break
		}
	}
	if head == nil {
		lst.PushBack(s)
		fmt.Printf("list中不存在%d,已将%d添加到首部", s, s)
	} else {
		lst.Remove(head)
		lst.PushFront(s)
	}
}
func TravsList(lst *list.List) {
	head := lst.Front()
	for head.Next() != nil {
		fmt.Printf("%v", head.Value)
		head = head.Next()
	}
	fmt.Printf("%v ", head.Value)
}

//3.heap实现过期缓存
type Item struct {
	value    string
	priority int //优先级
}
type Queue []*Item

func (pq Queue) Len() int {
	return len(pq) //返回元素个数
}

func (pq Queue) Less(i, j int) bool {
	return pq[i].priority > pq[j].priority //默认为小根堆 ,反着定义
}
func (pq Queue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}
func (pq *Queue) Push(x interface{}) {
	item := x.(*Item)
	*pq = append(*pq, item)
}
func (pq *Queue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[0 : n-1]
	return item
}
func updateQueue(pq Queue) {
	heap.Init(&pq)
	heap.Push(&pq, &Item{"D", 6})
	for pq.Len() > 0 {
		fmt.Println(heap.Pop(&pq))
	}
}

func main() {
	// s := "zhangsan transfer 10 RMB to lisi"
	// signature := DigitalSignature(s)
	// fmt.Println("验证数字签名", verifySignature(s, signature))

	// lst := list.New()
	// lst.PushBack(1)
	// lst.PushBack(2)
	// lst.PushBack(3)
	// lst.PushBack(4)
	// TravsList(lst)
	// visit(lst, 3)
	// TravsList(lst)
	pq := make(Queue, 0, 10)
	pq.Push(&Item{"A", 3})
	pq.Push(&Item{"B", 2})
	pq.Push(&Item{"C", 4})
	updateQueue(pq)
	fmt.Println("--------")
	pq.Push(&Item{"E", 5})
	updateQueue(pq)
}
