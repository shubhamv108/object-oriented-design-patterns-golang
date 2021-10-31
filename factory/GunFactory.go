package factory

import (
	"fmt"
	"object-oriented-design-patterns-golang/factory/guns"
)

func GetGun(gunType guns.GunType) (guns.IGun, error) {
	if gunType == guns.AK_47 {
		return guns.NewAK47(), nil
	}
	if gunType == guns.Maverick4 {
		return guns.NewM4(), nil
	}
	return nil, fmt.Errorf("Invalid gun type")
}