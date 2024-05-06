package main

import "github.com/01-edu/z01"

func main() {
	IsNegative(-4)
	IsNegative(2)
	IsNegative(1)
}

func IsNegative(nb int) {
	if nb < 0 {
		z01.PrintRune('F')
	} else if nb > 1 {
		z01.PrintRune('T')
	} else {
		z01.PrintRune('F')
	}
	z01.PrintRune('\n')
}
