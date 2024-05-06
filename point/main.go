package main

import (
	"github.com/01-edu/z01"
)

type point struct {
	x int
	y int
}

func setPoint(ptr *point) {
	ptr.x = 42
	ptr.y = 21
}

func main() {
	points := &point{}

	setPoint(points)

	printRune("x = ", points.x)
	printRune(", y = ", points.y)
	z01.PrintRune('\n')
}

func printRune(s string, n int) {
	for _, r := range s {
		z01.PrintRune(r)
	}
	printIntAsRune(n)
}

func printIntAsRune(n int) {
	if n < 0 {
		z01.PrintRune('-')
		n *= -1
	}
	printPositiveIntAsRune(n)
}

func printPositiveIntAsRune(n int) {
	if n >= 10 {
		printPositiveIntAsRune(n / 10)
	}
	z01.PrintRune(rune('0' + n%10))
}
