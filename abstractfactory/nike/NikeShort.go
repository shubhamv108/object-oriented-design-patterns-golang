package nike

import (
	"object-oriented-design-patterns-golang/abstractfactory/products"
)

type NikeShort struct {
	Short products.Short
}

func (n NikeShort) SetLogo(logo string) {
	n.Short.SetLogo(logo)
}

func (n NikeShort) SetSize(size int) {
	n.Short.SetSize(size)
}

func (n NikeShort) GetLogo() {
	n.Short.GetLogo()
}

func (n NikeShort) GetSize() {
	n.Short.GetSize()
}

func (s NikeShort) String() string {
	return s.Short.String()
}