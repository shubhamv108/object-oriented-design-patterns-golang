package products

import "fmt"

type IShoe interface {
	SetLogo(logo string)
	SetSize(size int)
	GetLogo()
	GetSize()
	String() string
}

type Shoe struct {
	Logo string
	Size int
}

func (s Shoe) SetLogo(logo string) {
	s.Logo = logo
}

func (s Shoe) SetSize(size int) {
	s.Size = size
}

func (s Shoe) GetLogo() string {
	return s.Logo
}

func (s Shoe) GetSize() int {
	return s.Size
}

func (s Shoe) String() string {
	return fmt.Sprintf("{ Logo: %s, Size: %d }", s.Logo, s.Size)
}