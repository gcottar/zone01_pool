package main

import "github.com/01-edu/z01"

func main() {
	var myRune rune = '0'
	for i := 0; i <= 9; i++ {
		z01.PrintRune(myRune)
		myRune++
	}
	z01.PrintRune('\n')
}
