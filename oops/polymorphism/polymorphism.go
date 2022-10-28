package main


type Shape interface {
	Area() float64
}

type Rectangle struct {
	length float64
	width float64
}

type Square struct {
	side float64
}

func (rectangle *Rectangle) Area() float64 {
	return rectangle.length * rectangle.width
}

func (square *Square) Area() float64 {
	return square.side * square.side
}

func main() {
	rectangle := &Rectangle{
		length: 1.0,
		width: 2.0,
	}

	square := &Square{
		side: 1,
	}

	shapes := []Shape{square, rectangle}

	totalArea := float64(0)
	for _, shape := range shapes {
		totalArea += shape.Area()
	}

	println(totalArea)

}
