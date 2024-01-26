package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	catCommand := NewCatCommand()
	catCommand.Run()
}

type Flags struct {
	number   bool
	nonblank bool
}

type CatCommand struct {
	args  []string
	flags Flags
}

func NewCatCommand() CatCommand {
	const (
		defaultNumber   = false
		numberUsage     = "Number lines"
		defaultNonblank = false
		nonblankUsage   = "Number non-blank lines"
	)

	var number bool
	flag.BoolVar(&number, "number", defaultNumber, numberUsage)
	flag.BoolVar(&number, "n", false, numberUsage+" (shorthand)")

	var nonblank bool
	flag.BoolVar(&nonblank, "number-nonblank", defaultNonblank, numberUsage)
	flag.BoolVar(&nonblank, "b", false, nonblankUsage+" (shorthand)")

	flag.Parse()
	flags := Flags{
		number:   number,
		nonblank: nonblank,
	}

	return CatCommand{
		flags: flags,
		args:  flag.Args(),
	}
}

func (c *CatCommand) Run() {
	if len(c.args) < 1 {
		c.readStdin()
	} else {
		for _, v := range c.args {
			c.readFile(v)
		}
	}
}

func (c *CatCommand) readFile(path string) {
	file, err := os.Open(path)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Invalid file")
		os.Exit(int(flag.ExitOnError))
	}

	defer file.Close()
	c.print(file)
}

func (c *CatCommand) readStdin() {
	file := os.Stdin
	defer file.Close()
	c.print(file)
}

func (c *CatCommand) print(file *os.File) {
	scanner := bufio.NewScanner(file)
	line := 0

	for scanner.Scan() {
		text := scanner.Text()

		if c.flags.number {
			line++
			printLineNumber(line)
		} else if c.flags.nonblank && text != "" {
			line++
			printLineNumber(line)
		}

		fmt.Fprintln(os.Stdout, text)
	}
}

func printLineNumber(line int) {
	fmt.Fprintf(os.Stdout, "%s%d\t", strings.Repeat(" ", 6-len(strconv.Itoa(line))), line)
}
