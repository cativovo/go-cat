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

func TestCatCommandBustle(t *testing.T) {
	type Foo struct {
		expectedTestFile string
		flags            Flags
	}

	foo := []Foo{
		{
			expectedTestFile: "the-bustle.txt.out",
			flags: Flags{
				number:   false,
				nonblank: false,
			},
		},
		{
			expectedTestFile: "the-bustle.txt.b.out",
			flags: Flags{
				number:   false,
				nonblank: true,
			},
		},
		{
			expectedTestFile: "the-bustle.txt.n.out",
			flags: Flags{
				number:   true,
				nonblank: false,
			},
		},
	}

	for _, v := range foo {
		var resultWriter bytes.Buffer
		var errWriter bytes.Buffer

		catCommand := NewCatCommand(&resultWriter, &errWriter, v.flags, []string{BUSTLE})
		catCommand.Run()

		result := resultWriter.String()
		expected := readExpectedFile(t, v.expectedTestFile)

		if expected != result {
			t.Error("eyy mali", v.expectedTestFile)
			t.Error("result\n", result)
			t.Error("expected\n", expected)
		}
	}
}

func readExpectedFile(t *testing.T, filename string) string {
	path := "./test/expected/" + filename
	data, err := os.ReadFile(path)
	if err != nil {
		t.Fatal(path + " not found")
	}

	return string(data)
}
