package main

import (
	"fmt"
	"os"
)

func main() {
	fileName := "quest8.txt"

	if len(os.Args) < 2 {
		fmt.Println("File name missing")
		return
	}
	if len(os.Args) > 2 {
		fmt.Println("Too many arguments")
		return
	}
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Print("Error opening file:", err)
		return
	}
	fileInfo, err := file.Stat()
	if err != nil {
		fmt.Print("Error getting file information:", err)
		return
	}
	data := make([]byte, fileInfo.Size())
	_, err = file.Read(data)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}
	file.Close()
	fmt.Print(string(data))
}
