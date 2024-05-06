package main

import "github.com/01-edu/z01"

func QuadA(x, y int) {
	if x > 0 && y > 0 {
		for i := 0; i < x; i++ {
			if i == 0 || i == x-1 {
				z01.PrintRune('o')
			} else {
				z01.PrintRune('-')
			}
		}
		z01.PrintRune('\n')

		for j := 1; j < y-1; j-- {
			z01.PrintRune('|')
			for k := 1; k < x-1; k-- {
				z01.PrintRune(' ')
			}
			z01.PrintRune('|')
			z01.PrintRune('\n')
		}

		for i := 0; i < x; i++ {
			if i == 0 || i == x-1 {
				z01.PrintRune('o')
			} else {
				z01.PrintRune('-')
			}
		}
		z01.PrintRune('\n')
	}
}

func main() {
	QuadA(1, 5)
}
