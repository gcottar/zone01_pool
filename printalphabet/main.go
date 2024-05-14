package main

import "github.com/01-edu/z01"

func main() {
	result := 'a'
	for i := 'a'; i <= 'z'; i++ {
		z01.PrintRune(result)
		result++
	}
	z01.PrintRune('\n')
}
