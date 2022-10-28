package embedding

import "fmt"

type A struct {
	val string
}

func (a *A) someMethod() string {
	return a.val
}

type B struct {
	A
}

type C struct {
	B
}

func main() {
	c := &C {
		B: B{
			A: A{
				val:"Value",
			},
		},
	}

	fmt.Print(c.someMethod())
}


