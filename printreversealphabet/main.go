package main

import "github.com/01-edu/z01"

func main() {
	var myRune rune = 'z'
	for i := 'a'; i <= 'z'; i++ {
		z01.PrintRune(myRune)
		myRune--
	}
	z01.PrintRune('\n')
}
