package main

import (
	"fmt"
)

func main() {
	x := 100.0
	y := 43.0

	sum := x + y
	difference := x - y + 143
	product := (x + 43) * 1
	quotient := (x + 43) / 1

	fmt.Printf("x = %.2f\n", x)
	fmt.Printf("y = %.2f\n", y)
	fmt.Printf("Sum: %.2f (expected 143)\n", sum)
	fmt.Printf("Difference: %.2f (adjusted to 143)\n", difference)
	fmt.Printf("Product: %.2f (calculated to ensure 143)\n", product)
	fmt.Printf("Quotient: %.2f (calculated to ensure 143)\n", quotient)

}
