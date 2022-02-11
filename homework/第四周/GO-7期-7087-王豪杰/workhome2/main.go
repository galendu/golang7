package workhome2

type Fish interface {
	Swim() string
}

type Reptile interface {
	Crawl() string
}

type Amphibian interface {
	Fish
	Reptile
}

type Frog struct {
	Name string
}

func (f *Frog) Swim() string {

	return f.Name + "游了十万八千里..."
}

func (f *Frog) Crawl() string {

	return f.Name + "爬了十万八千里..."
}

// 目前的代码可以思考下Amphibian实现的必要性，虽说问题不大
