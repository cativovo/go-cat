package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// first value in this slice is the path to the program
	args := os.Args[1:]

	for _, v := range args {
		readFile(v)
	}
}

func readFile(path string) {
	file, err := os.Open(path)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Invalid file")
		return
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}
