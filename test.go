package main

import "fmt"
import "math"

func rectangle(x, y int) (area, perimeter int, diagonal float64) {
    area = x*y
    perimeter = 2*x + 2*y
    diagonal = math.Sqrt(math.Pow(float64(x), 2) + math.Pow(float64(y), 2))
    return
}

func main() {
    x, y := 10, 12
    area, _, diagonal := rectangle(x, y)
    fmt.Printf("A %vx%v rectangle has an area of %v, and a diagonal of %v.\n", x, y, area, diagonal)
}
