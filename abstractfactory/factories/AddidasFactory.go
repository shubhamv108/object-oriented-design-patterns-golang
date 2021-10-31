package abstractfactory

import (
	"object-oriented-design-patterns-golang/abstractfactory/addidas"
	"object-oriented-design-patterns-golang/abstractfactory/products"
)

type AddidasFactory struct {}

func (a *AddidasFactory) MakeShoe() products.IShoe {
	return &addidas.AddidasShoe{
		Shoe: products.Shoe{
			Logo: "AddidasFactory",
			Size: 14,
		},
	}
}

func (a *AddidasFactory) MakeShort() products.IShort {
	return &addidas.AddidasShort{
		Short: products.Short{
			Logo: "AddidasFactory",
			Size: 14,
		},
	}
}
