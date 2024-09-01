package main

import (
	"fmt"

	"github.com/rudolfkova/go/shape"
)

func main() {
	var circle shape.Circle
	circle.Radius = 10
	fmt.Println(circle.Area())
	fmt.Scanln()
}
