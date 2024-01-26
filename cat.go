package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func main() {
	flags, args := getFlagsAndArgs()
	catCommand := NewCatCommand(os.Stdout, os.Stderr, flags, args)
	catCommand.Run()
}

type Flags struct {
	number   bool
	nonblank bool
}

func getFlagsAndArgs() (Flags, []string) {
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

	return flags, flag.Args()
}

type CatCommand struct {
	resultWriter io.Writer
	errWriter    io.Writer
	args         []string
	flags        Flags
}

func NewCatCommand(resultWriter io.Writer, errWriter io.Writer, flags Flags, args []string) CatCommand {
	return CatCommand{
		flags:        flags,
		args:         args,
		resultWriter: resultWriter,
		errWriter:    errWriter,
	}
}

func (c *CatCommand) Run() {
	var result string

	if len(c.args) < 1 {
		result = c.readStdin()
	} else {
		for _, v := range c.args {
			r, err := c.readFile(v)
			if err != nil {
				fmt.Fprintln(c.errWriter, "Invalid File")
				os.Exit(1)
			}

			result += r
		}
	}

	fmt.Fprint(c.resultWriter, result)
}

func (c *CatCommand) readFile(path string) (string, error) {
	file, err := os.Open(path)
	if err != nil {
		return "", errors.New("invalid file")
	}

	defer file.Close()
	return c.getContents(file), nil
}

func (c *CatCommand) readStdin() string {
	file := os.Stdin
	defer file.Close()
	return c.getContents(file)
}

func (c *CatCommand) getContents(file *os.File) string {
	var result string
	scanner := bufio.NewScanner(file)
	line := 0

	for scanner.Scan() {
		text := scanner.Text()
		showNonblank := c.flags.nonblank && text != ""

		if c.flags.number {
			line++
		} else if showNonblank {
			line++
		}

		result += formatText(text, line, c.flags.number || showNonblank)
	}

	return result
}

func formatText(text string, line int, showLineNumber bool) string {
	if !showLineNumber {
		return fmt.Sprintf("%s\n", text)
	}

	return fmt.Sprintf("%s%d\t%s\n", strings.Repeat(" ", 6-len(strconv.Itoa(line))), line, text)
}
