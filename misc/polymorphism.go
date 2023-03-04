package main

import "fmt"

// Shape is an interface that defines a method called Area.
type Shape interface {
    Area() float64
}

// Circle is a type that satisfies the Shape interface.
type Circle struct {
    Radius float64
}

// Area calculates the area of a circle.
func (c Circle) Area() float64 {
    return 3.14 * c.Radius * c.Radius
}

// Rectangle is a type that satisfies the Shape interface.
type Rectangle struct {
    Width  float64
    Height float64
}

// Area calculates the area of a rectangle.
func (r Rectangle) Area() float64 {
    return r.Width * r.Height
}

func main() {
    // Create instances of Circle and Rectangle
    circle := Circle{Radius: 5.0}
    rectangle := Rectangle{Width: 4.0, Height: 6.0}

    // Calculate and print the areas using polymorphism
    printArea(circle)
    printArea(rectangle)
}

// printArea takes any type that satisfies the Shape interface and calculates its area.
func printArea(s Shape) {
    fmt.Printf("Area: %0.2f\n", s.Area())
}
