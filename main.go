package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// first value in this slice is the path to the program
	args := os.Args[1:]

	if len(args) < 1 {
		readStdin()
	} else {
		for _, v := range args {
			readFile(v)
		}
	}
}

func readFile(path string) {
	file, err := os.Open(path)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Invalid file")
		return
	}

	defer file.Close()
	print(file)
}

func readStdin() {
	file := os.Stdin
	defer file.Close()
	print(file)
}

func print(file *os.File) {
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}
