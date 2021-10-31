package addidas

import (
	"object-oriented-design-patterns-golang/abstractfactory/products"
)

type AddidasShort struct {
	Short products.Short
}

func (n AddidasShort) SetLogo(logo string) {
	n.Short.SetLogo(logo)
}

func (n AddidasShort) SetSize(size int) {
	n.Short.SetSize(size)
}

func (n AddidasShort) GetLogo() {
	n.Short.GetLogo()
}

func (n AddidasShort) GetSize() {
	n.Short.GetSize()
}

func (s AddidasShort) String() string {
	return s.Short.String()
}