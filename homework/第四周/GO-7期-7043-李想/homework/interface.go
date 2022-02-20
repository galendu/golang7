package homework

import "fmt"

type Reptile interface{
	Crawl()
}

type Fish interface{
	Swim()
}

type Frog struct {
	Name string
	Varieties string
}

func(f Frog) Crawl(){
	fmt.Printf("%s在爬行，它的种类是：%s\n",f.Name,f.Varieties)
}

func(f Frog) Swim(){
	fmt.Printf("%s在游泳，它的种类是：%s\n",f.Name,f.Varieties)
}



