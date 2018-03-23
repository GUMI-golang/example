package main

import "fmt"

type If interface {
	Fn()
}
type Test struct {
	T int
}

func (*Test) Fn()  {

}

func main() {
	var t If = &Test{T: 3}

	fmt.Println(t)
	CH(&t)
	fmt.Println(t)
}
func CH(ptif *If) {
	if v, ok := (*ptif).(*Test); ok {
		v.T = 11
	}
}