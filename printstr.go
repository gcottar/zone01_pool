package main

import "github.com/01-edu/z01"

func main() {
	PrintStr("Hello Wolrd")
}

func PrintStr(s string) {
	hello := []rune(s)
	for i := 0; i < len(s); i++ {
		z01.PrintRune(hello[i])
	}
	z01.PrintRune('\n')
}
