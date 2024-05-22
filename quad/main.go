package main

import "github.com/01-edu/z01"

func QuadA(x, y int) {
	var xi int
	var yi int
	if x > 0 && y > 0 {
		yi = 0
		for yi < y {
			xi = 0
			for xi < x {
				if xi == 0 && yi == 0 || xi == x-1 && yi == 0 ||
					xi == 0 && yi == y-1 || xi == x-1 && yi == y-1 {
					// if in any corner
					z01.PrintRune('o')
				} else if yi == 0 || yi == y-1 {
					// else if in first or last row
					z01.PrintRune('-')
				} else if xi == 0 || xi == x-1 {
					// else if in first or last column
					z01.PrintRune('|')
				} else {
					z01.PrintRune(' ')
				}
				xi++
			}
			z01.PrintRune('\n')
			yi++
		}
	}
}

func main() {
	QuadA(25, 20)
}
