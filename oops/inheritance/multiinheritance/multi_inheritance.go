package main

import "fmt"


type A struct {
	val string
}

func (a *A) value() string {
	return a.val
}

type D struct {
	val string
}

func (d *D) someMethod() string {
	return d.val
}

type B struct {
	A
}

type E struct {
	val string
}

func (e *E) valueOfE() string {
	return e.val
}

type C struct {
	A
	B
	//D
	E
}

func main() {
	c := &C {
		B: B{
			A: A{
				val:"Value: CBA",
			},
		},
		A: A{
			val: "Value: CA",
		},
		//D: D{
		//	val: "Value CD",
		//},
		E: E{
			val: "Value: CE",
		},
	}

	fmt.Println(c.value()) // ambigious reference after adding d other wise A.someMethod is called
	fmt.Print(c.valueOfE())
}



