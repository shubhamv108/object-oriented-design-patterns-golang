package nike

import (
	"object-oriented-design-patterns-golang/abstractfactory/products"
)

type NikeShoe struct {
	Shoe products.Shoe
}

func (n NikeShoe) SetLogo(logo string) {
	n.SetLogo(logo)
}

func (n NikeShoe) SetSize(size int) {
	n.SetSize(size)
}

func (n NikeShoe) GetLogo() {
	n.GetLogo()
}

func (n NikeShoe) GetSize() {
	n.GetSize()
}

func (s NikeShoe) String() string {
	return s.Shoe.String()
}



