package main

import (
	"object-oriented-design-patterns-golang/factory"
	"object-oriented-design-patterns-golang/factory/guns"
)

func main() {
	ak47, _ := factory.GetGun(guns.AK_47);
	m4, _ := factory.GetGun(guns.Maverick4);
	println(ak47.GetName());
	println(m4.GetName());
}
