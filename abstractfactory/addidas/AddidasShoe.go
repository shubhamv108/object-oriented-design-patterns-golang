package addidas

import (
	"object-oriented-design-patterns-golang/abstractfactory/products"
)

type AddidasShoe struct {
	Shoe products.Shoe
}

func (n AddidasShoe) SetLogo(logo string) {
	n.Shoe.SetLogo(logo)
}

func (n AddidasShoe) SetSize(size int) {
	n.Shoe.SetSize(size)
}

func (n AddidasShoe) GetLogo() {
	n.Shoe.GetLogo()
}

func (n AddidasShoe) GetSize() {
	n.Shoe.GetSize()
}

func (s AddidasShoe) String() string {
	return s.Shoe.String()
}
