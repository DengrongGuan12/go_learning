package main

import (
	"fmt"
	"reflect"
)

type inf interface {
	Method1()
	Method2()
}

type ss struct {
	a func()
}

func (i ss) Method1() {}
func (i ss) Method2() {}

func main() {
	s := reflect.TypeOf(ss{})
	i := reflect.TypeOf(new(inf)).Elem()

	Test(s)
	Test(i)
}

func Test(t reflect.Type) {
	if t.NumMethod() > 0 {
		fmt.Printf("\n--- %s ---\n", t)
		fmt.Println(t.Method(0).Type)
		fmt.Println(t.Method(0).Func.String())
	}
}

// func main() {
// 	fmt.Printf("Hello, world.  Sqrt(2) = %v\n", mymath2.Sqrt(2))
// 	mymath2.Adda("sdfsdfsdf")

// }
