package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	arguments := os.Args[1:]

	length := 0
	for i := range arguments {
		length = i + 1
	}
	if length != 3 {
		return
	}

	sign := 0
	if arguments[1] == "+" {
		sign = 0
	} else if arguments[1] == "-" {
		sign = 1
	} else if arguments[1] == "*" {
		sign = 2
	} else if arguments[1] == "/" {
		sign = 3
	} else if arguments[1] == "%" {
		sign = 4
	} else {
		z01.PrintRune('0')
		z01.PrintRune('\n')
		return
	}

	for i, s := range arguments[0] {
		if (s >= '0' && s <= '9') || (i == 0 && s == '-') {
			continue
		} else {
			z01.PrintRune('0')
			z01.PrintRune('\n')
			return
		}
	}

	for i, s := range arguments[2] {
		if (s >= '0' && s <= '9') || (i == 0 && s == '-') {
			continue
		} else {
			z01.PrintRune('0')
			z01.PrintRune('\n')
			return
		}
	}

	firstNbr := Atoi(arguments[0])
	secondNbr := Atoi(arguments[2])

	if secondNbr == 0 && arguments[1] == "/" {
		PrintStr("No division by 0\n")
		return
	}
	if secondNbr == 0 && arguments[1] == "%" {
		PrintStr("No Modulo by 0\n")
		return
	}

	result := 0
	arrayOfFunctions := []func(int, int) int{plus, minus, times, div, mod}
	result = apply(arrayOfFunctions[sign], firstNbr, secondNbr)

	PrintStr(string(result))
	PrintStr("\n")
}

func plus(a, b int) int {
	return a + b
}

func minus(a, b int) int {
	return a - b
}

func times(a, b int) int {
	return a * b
}

func div(a, b int) int {
	return a / b
}

func mod(a, b int) int {
	return a % b
}

func apply(f func(int, int) int, a int, b int) int {
	return f(a, b)
}

func Atoi(s string) int {
	runes := []rune(s)
	LenRune := 0
	result := 0
	for i := range runes {
		LenRune = i + 1
	}
	if LenRune == 0 {
		return 0
	}

	tens := 1
	for k := 0; k < LenRune-1; k++ {
		if runes[k] == '+' || runes[k] == '-' {
			continue
		}
		tens *= 10
	}

	for i := range runes {
		if (runes[0] == '-' || runes[0] == '+') && i == 0 {
			continue
		}
		if runes[i] < '0' || runes[i] > '9' {
			return 0
		}
		numb := 0
		for j := '0'; j < runes[i]; j++ {
			numb++
		}
		result += numb * tens
		tens /= 10
	}
	if runes[0] == '-' {
		result *= -1
	}
	return result
}

func PrintStr(s string) {
	for _, r := range s {
		z01.PrintRune(r)
	}
}
