package main

import "fmt"

type BarFunc func() error

type Adapter struct {
	BarFunc
}

type T struct{}

func (t T) Foo() {
	a := Adapter{t.Bar}
	b := Adapter{t.Bar2}
	a.BarFunc()
	b.BarFunc()
}

func (t T) Bar() error {
	fmt.Println(1)
	return nil
}

func (t T) Bar2() error {
	fmt.Println(2)
	return nil
}

func main() {
	x := T{}
	x.Foo()
}

// Output:
// 1
// 2
