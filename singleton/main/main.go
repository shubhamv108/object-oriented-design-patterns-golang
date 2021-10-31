package main

import (
	"fmt"
	"object-oriented-design-patterns-golang/singleton"
)

func main() {
	for i := 1; i <= 100; i++ {
		println(i);
		go singleton.GetInstance();
	}

	fmt.Scanln()
}
