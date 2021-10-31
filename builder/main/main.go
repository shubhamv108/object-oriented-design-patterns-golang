package main

import (
	. "object-oriented-design-patterns-golang/builder/typedobject"
)

func main() {
	typeObject := Builder().
		WithAttribute1("A").
		WithAttribute2("B").
		Build();
	println(typeObject.String())

}
