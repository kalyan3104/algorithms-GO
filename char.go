package main

import (
	"fmt"
)

func main() {

	var char1 rune = 'A'
	var char2 rune = 'b'
	var char3 rune = '#'

	fmt.Printf("char1: %c, ASCII: %d\n", char1, char1)
	fmt.Printf("char2: %c, ASCII: %d\n", char2, char2)
	fmt.Printf("char3: %c, ASCII: %d\n", char3, char3)

	isEqual := char1 == 'A'
	fmt.Printf("Is char1 equal to 'A'? %v\n", isEqual)

	message := string(char1) + " is the first letter, and " + string(char2) + " is the second."
	fmt.Println(message)
}
