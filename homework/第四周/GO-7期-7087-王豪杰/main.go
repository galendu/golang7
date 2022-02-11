package main

import (
	"day4/workhome1"
	"day4/workhome2"
	"day4/workhome3"
	"fmt"
)

func main() {

	var amphibian workhome2.Amphibian

	frog := &workhome2.Frog{Name: "青蛙王子"}
	amphibian = frog

	fmt.Println(amphibian.Swim()+"\n", amphibian.Crawl())

	s := workhome1.Square(8.9)
	fmt.Println(s)

	work, err := workhome3.Work(10.0, 100, 424.12414)
	fmt.Println(work)

	work, err = workhome3.Work(10, 0)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(work)

	work, err = workhome3.Work()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(work)
}
