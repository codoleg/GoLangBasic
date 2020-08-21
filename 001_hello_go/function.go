package main

import (
	"fmt"
)

func paintNeeded(width float64, height float64) {
	area := width * height
	fmt.Printf("%.2f liters needed\n", area/10.0)
}
func main() {
	paintNeeded(4.2, 3.0)
	paintNeeded(5.2, 4.0)
}
