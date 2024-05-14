package main

import "github.com/01-edu/z01"

func main() {
	result := '0'
	for i := 0; i <= 9; i++ {
		z01.PrintRune(result)
		result++
	}
	z01.PrintRune('\n')
}
