package main

import "fmt"

type Fisher interface {
	swimming()
}

type Crawler interface {
	crawling()
}

type Animal struct {
	name string
}

func (f Animal) swimming() {
	fmt.Printf("%s can swim\n", f.name)
}

func (f Animal) crawling() {
	fmt.Printf("%s can crawl\n", f.name)
}

func main() {
	frog := Animal{"frog"}
	var fisher Fisher
	var crawler Crawler
	fisher = frog
	crawler = frog
	fisher.swimming()
	crawler.crawling()
}

// 完成的不错
