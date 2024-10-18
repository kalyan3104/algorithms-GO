package main

import (
	"fmt"
)

func main() {
	a := true
	b := false

	andResult := a && b
	orResult := a || b
	notA := !a
	notB := !b

	fmt.Println("a =", a)
	fmt.Println("b =", b)
	fmt.Println("AND (a && b):", andResult)
	fmt.Println("OR (a || b):", orResult)
	fmt.Println("NOT a (!a):", notA)
	fmt.Println("NOT b (!b):", notB)
}
