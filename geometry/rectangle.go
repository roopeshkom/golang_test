package geometry

import (
    "math"
)

func Rectangle(x, y int) (area, perimeter int, diagonal float64) {
    area = x*y
    perimeter = 2*x + 2*y
    diagonal = math.Sqrt(math.Pow(float64(x), 2) + math.Pow(float64(y), 2))
    return
}

func Circle(radius int) (area, circumference float64) {
    area = math.Pow(float64(radius), 2) * math.Pi
    circumference = 2 * math.Pi * float64(radius)
    return
}
