package main

import "fmt"

type iA interface {
	method1()
	method2()
}

type A struct {
	name string
	method2 func()
}

func (a *A) method1() {
	fmt.Println("A Method1")
	a.method2()
}

//func (a *A) method2() {
//	fmt.Println("A Method1")
//	a.method2()
//}


type B struct {
	A
}

func (b *B) method2() {
	fmt.Println("B Method2")
	// b.A.method2()
	b.method1()
}

func main() {
	a := A{
		name: "AbstractClass",
	}
	b := B{
		A: a,
	}
	b.A.method2 = b.method2

	b.method2()
}
