package products

import "fmt"

type IShort interface {
	SetLogo(logo string)
	SetSize(size int)
	GetLogo()
	GetSize()
	String() string
}

type Short struct {
	Logo string
	Size int
}

func newShort(logo string, size int) *Shoe {
	return &Shoe{Logo: logo, Size: size}
}

func (s Short) SetLogo(logo string) {
	s.Logo = logo
}

func (s Short) SetSize(size int) {
	s.Size = size
}

func (s Short) GetLogo() string {
	return s.Logo
}

func (s Short) GetSize() int {
	return s.Size
}

func (s Short) String() string {
	return fmt.Sprintf("{ Logo: %s, Size: %d }", s.Logo, s.Size)
}

