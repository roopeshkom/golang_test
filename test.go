package main

import "fmt"
import "test/geometry"

func main() {
    x, y := 10, 12
    radius := 7

    area, _, diagonal := geometry.Rectangle(x, y)
    area2, circumference := geometry.Circle(radius)

    fmt.Printf("A %vx%v rectangle has an area of %v, and a diagonal of %v.\n", x, y, area, diagonal)
    fmt.Printf("A circle with radius %v has an area of %v, and a circumference of %v.\n", radius, area2, circumference)
}
