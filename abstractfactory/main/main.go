package main

import abstractfactory "object-oriented-design-patterns-golang/abstractfactory/factories"

func main() {
	addidasFactory, _ := abstractfactory.GetSportsFactory("Addidas");
	nikeFactory, _ := abstractfactory.GetSportsFactory("Nike");
	println(addidasFactory.MakeShoe().String());
	println(addidasFactory.MakeShort().String());
	println(nikeFactory.MakeShoe().String());
	println(nikeFactory.MakeShort().String());
}
