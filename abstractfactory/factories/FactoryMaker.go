package abstractfactory

import "fmt"

func GetSportsFactory(brand string) (ISportsFactory, error) {
	switch brand {
		case "Addidas": return &AddidasFactory{}, nil
		case "Nike": return &NikeFactory{}, nil
		default: return nil, fmt.Errorf("No such shoe brand")
	}
}
