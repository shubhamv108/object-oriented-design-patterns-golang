package abstractfactory

import (
	"object-oriented-design-patterns-golang/abstractfactory/products"
)

type ISportsFactory interface {
	MakeShoe() products.IShoe
	MakeShort() products.IShort
}