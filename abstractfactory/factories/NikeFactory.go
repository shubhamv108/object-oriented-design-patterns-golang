package abstractfactory

import (
	"object-oriented-design-patterns-golang/abstractfactory/nike"
	"object-oriented-design-patterns-golang/abstractfactory/products"
)

type NikeFactory struct {}

func (a *NikeFactory) MakeShoe() products.IShoe {
	return &nike.NikeShoe{
		Shoe: products.Shoe{
			Logo: "NikeFactory",
			Size: 14,
		},
	}
}

func (a *NikeFactory) MakeShort() products.IShort {
	return &nike.NikeShort{
		Short: products.Short{
			Logo: "NikeFactory",
			Size: 14,
		},
	}
}
