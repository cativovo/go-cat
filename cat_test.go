package main

import (
	"bytes"
	"os"
	"testing"
)

const (
	EMPTY   = "./test/inputs/empty.txt"
	FOX     = "./test/inputs/fox.txt"
	SPIDERS = "./test/inputs/spiders.txt"
	BUSTLE  = "./test/inputs/the-bustle.txt"
)

type TestData struct {
	expectedFile string
	flags        Flags
}

func TestCatCommandBustle(t *testing.T) {
	testData := []TestData{
		{
			expectedFile: "the-bustle.txt.out",
			flags: Flags{
				number:   false,
				nonblank: false,
			},
		},
		{
			expectedFile: "the-bustle.txt.b.out",
			flags: Flags{
				number:   false,
				nonblank: true,
			},
		},
		{
			expectedFile: "the-bustle.txt.n.out",
			flags: Flags{
				number:   true,
				nonblank: false,
			},
		},
	}

	for _, v := range testData {
		var resultWriter bytes.Buffer
		var errWriter bytes.Buffer

		catCommand := NewCatCommand(&resultWriter, &errWriter, v.flags, []string{BUSTLE})
		catCommand.Run()

		result := resultWriter.String()
		expected := readExpectedFile(t, v.expectedFile)

		if expected != result {
			t.Error("error in", v.expectedFile)
			t.Error("result\n", result)
			t.Error("expected\n", expected)
		}
	}
}

func TestCatCommandBustleStdIn(t *testing.T) {
	testData := []TestData{
		{
			expectedFile: "the-bustle.txt.stdin.out",
			flags: Flags{
				number:   false,
				nonblank: false,
			},
		},
		{
			expectedFile: "the-bustle.txt.b.stdin.out",
			flags: Flags{
				number:   false,
				nonblank: true,
			},
		},
		{
			expectedFile: "the-bustle.txt.n.stdin.out",
			flags: Flags{
				number:   true,
				nonblank: false,
			},
		},
	}

	for _, v := range testData {
		var resultWriter bytes.Buffer
		var errWriter bytes.Buffer

		input := []byte(readFile(t, BUSTLE))
		r, w, err := os.Pipe()
		if err != nil {
			t.Fatal(err)
		}

		_, err = w.Write(input)
		if err != nil {
			t.Error(err)
		}
		w.Close()
		originalStdin := os.Stdin
		os.Stdin = r

		catCommand := NewCatCommand(&resultWriter, &errWriter, v.flags, []string{})
		catCommand.Run()

		result := resultWriter.String()
		expected := readExpectedFile(t, v.expectedFile)

		if expected != result {
			t.Error("error in", v.expectedFile)
			t.Error("result\n", result)
			t.Error("expected\n", expected)
		}

		// restore stdin
		os.Stdin = originalStdin
	}
}

func TestCatCommandFox(t *testing.T) {
	testData := []TestData{
		{
			expectedFile: "fox.txt.out",
			flags: Flags{
				number:   false,
				nonblank: false,
			},
		},
		{
			expectedFile: "fox.txt.b.out",
			flags: Flags{
				number:   false,
				nonblank: true,
			},
		},
		{
			expectedFile: "fox.txt.n.out",
			flags: Flags{
				number:   true,
				nonblank: false,
			},
		},
	}

	for _, v := range testData {
		var resultWriter bytes.Buffer
		var errWriter bytes.Buffer

		catCommand := NewCatCommand(&resultWriter, &errWriter, v.flags, []string{FOX})
		catCommand.Run()

		result := resultWriter.String()
		expected := readExpectedFile(t, v.expectedFile)

		if expected != result {
			t.Error("error in", v.expectedFile)
			t.Error("result\n", result)
			t.Error("expected\n", expected)
		}
	}
}

func TestCatCommandEmpty(t *testing.T) {
	testData := []TestData{
		{
			expectedFile: "empty.txt.out",
			flags: Flags{
				number:   false,
				nonblank: false,
			},
		},
		{
			expectedFile: "empty.txt.b.out",
			flags: Flags{
				number:   false,
				nonblank: true,
			},
		},
		{
			expectedFile: "empty.txt.n.out",
			flags: Flags{
				number:   true,
				nonblank: false,
			},
		},
	}

	for _, v := range testData {
		var resultWriter bytes.Buffer
		var errWriter bytes.Buffer

		catCommand := NewCatCommand(&resultWriter, &errWriter, v.flags, []string{EMPTY})
		catCommand.Run()

		result := resultWriter.String()
		expected := readExpectedFile(t, v.expectedFile)

		if expected != result {
			t.Error("error in", v.expectedFile)
			t.Error("result\n", result)
			t.Error("expected\n", expected)
		}
	}
}

func readFile(t *testing.T, path string) string {
	data, err := os.ReadFile(path)
	if err != nil {
		t.Error(path + " not found")
	}

	return string(data)
}

func readExpectedFile(t *testing.T, filename string) string {
	return readFile(t, "./test/expected/"+filename)
}
