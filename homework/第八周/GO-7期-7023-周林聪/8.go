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
	"strconv"
	"time"
)

//1.实现数字签名
var (
	privateKeyfilepath = "G:\\git仓库\\golang7\\homework\\第八周\\zlc_privatekey.pem"
	publicKeyfilepath  = "G:\\git仓库\\golang7\\homework\\第八周\\zlc_publickey.pem"
)

func Digital_signature(trade string) ([]byte, error) {
	sha1 := sha1.New()
	sha1.Write([]byte(trade))
	digest := sha1.Sum([]byte(""))
	privatekey_byte, err := os.ReadFile(privateKeyfilepath)
	if err != nil {
		return nil, err
	}
	block, _ := pem.Decode(privatekey_byte)
	RSAprivatekey, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	signature, err := rsa.SignPKCS1v15(nil, RSAprivatekey, crypto.Hash(0), digest)
	if err != nil {
		return nil, err
	}
	return signature, nil
}

func verify_sibnature(trade string, signature []byte) (error, bool) {
	sha1 := sha1.New()
	sha1.Write([]byte(trade))
	digest := sha1.Sum([]byte(""))
	publickey_byte, err := os.ReadFile(publicKeyfilepath)
	if err != nil {
		return err, false
	}
	block, _ := pem.Decode(publickey_byte)
	publickey_interf, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return err, false
	}
	publickey := publickey_interf.(*rsa.PublicKey)
	err = rsa.VerifyPKCS1v15(publickey, crypto.Hash(0), digest, signature)
	if err != nil {
		return err, false
	}
	return nil, err == nil
}

//2.用map和链表实现LRU缓存
var (
	Cache    = make(map[int]string, 10)
	LRU_list = list.New()
)

func LRUread(key int) string {
	value, exists := Cache[key]
	if exists {
		head := LRU_list.Front()
		if head.Value.(int) == key {
			return value
		}
		for i := 0; i < LRU_list.Len(); i++ {
			head = head.Next()
			if head.Value.(int) == key {
				LRU_list.MoveToFront(head)
			}
		}
	}
	value = "come" + strconv.Itoa(key)
	Cache[key] = value
	LRU_list.PushFront(key)
	if len(Cache) >= 10 {
		tail := LRU_list.Back()
		delete(Cache, tail.Value.(int))
		LRU_list.Remove(tail)
		fmt.Printf("因链表已满，将最近最少使用的%d从链表中删除\n", tail.Value.(int))
	}
	return value
}

func Traverlist(lst *list.List) {
	head := lst.Front()
	for head.Next() != nil {
		fmt.Printf("%v ", head.Value)
		head = head.Next()
	}
	fmt.Println(head.Value)
}

//3.用map和堆 表实现超时缓存
// var (
//
// )

type Node struct {
	key     int
	value   string
	timeout int64
}

type timeout_heap []*Node

func (th timeout_heap) Len() int {
	return len(th)
}

func (th timeout_heap) Less(i, j int) bool {
	return th[i].timeout < th[j].timeout
}

func (th timeout_heap) Swap(i, j int) {
	th[i], th[j] = th[j], th[i]
}

func (th *timeout_heap) Push(x interface{}) {
	node := x.(*Node)
	*th = append(*th, node)
}

func (th *timeout_heap) Pop() interface{} {
	old := *th
	n := len(old)
	ele := old[n-1]
	*th = old[0 : n-1]
	return ele
}

func Nodefangfa(key int) *Node {
	return &Node{
		key:     key,
		value:   "come" + strconv.Itoa(key),
		timeout: time.Now().Unix() + 26,
	}
}

func Timeout_Cache2() {
	th := make(timeout_heap, 0, 10)
	heap.Init(&th)
	Cache2 := make(map[int]*Node, 10)
	for i := 0; i < 10; i++ {
		node := Nodefangfa(i)
		fmt.Printf("node:%v    ", node)
		Cache2[i] = node
		fmt.Printf("Cache2[i]:%v    ", Cache2[i])
		heap.Push(&th, node)
		fmt.Printf("th:%v,创建成功\n", th.Len())
		time.Sleep(2 * time.Second)
	}
	ticker := time.NewTicker(5 * time.Millisecond)
	for {
		<-ticker.C
		for {
			time_now := time.Now().Unix()
			if th.Len() <= 0 {
				break
			}
			first := th[0]
			if time_now < first.timeout {
				break
			} else {
				delete(Cache2, first.key)
				heap.Pop(&th)
				fmt.Printf("将超时缓存:%v删除,%d\n", *first, time_now)
			}
		}
		if th.Len() <= 0 {
			break
		}
	}

}

func main() {
	//1
	// trade := "10:30,张三在以太坊交易了10 BIC"
	// signature, err := Digital_signature(trade)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// err2, result := verify_sibnature(trade, signature)
	// if err2 != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Println("验证数字签名结果：", result)

	//2
	// for i := 0; i < 15; i++ {
	// 	_ = LRUread(i)
	// }
	// for k, v := range Cache {
	// 	fmt.Printf("%d:%s\n", k, v)
	// }
	// Traverlist(LRU_list)

	//3
	Timeout_Cache2()
}

//go run homework/8/8.go
