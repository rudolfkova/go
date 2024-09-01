package main

import (
	"fmt"
	"github.com/rudolfkova/go/shape"
)

type Circle struct {
	area float64
}

func main() {
	circle := Circle{Radius: 5}
	rectangle := Rectangle{Width: 4, Height: 3}

	fmt.Println("Circle area:", circle.Area())
	fmt.Println("Circle perimeter:", circle.Perimeter())
	fmt.Println("Rectangle area:", rectangle.Area())
	fmt.Println("Rectangle perimeter:", rectangle.Perimeter())
}
