package main

import "fmt"

func main() {
	test := []string{
		"test0",
		"test1",
		"test2",
		"test3",
		"test4",
		"test5",
	}
	for _, testing := range test {
		fmt.Printf("Print the caracter: %v\n", testing)
	}
}
