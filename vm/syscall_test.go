package vm

import (
	"io/ioutil"
	"os"
	"testing"
)

func stringToUint16(s string) []uint16 {
	result := make([]uint16, 0, len(s))

	for _, v := range s {
		result = append(result, uint16(v))
	}
	return result
}

func TestSyscallWritesString(t *testing.T) {
	input := stringToUint16("hello, world!")
	expectedOutput := "hello, world!"

	// Rewrite stdout pipe
	stdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	WriteString(input)

	// Revert to previous state
	w.Close()
	out, _ := ioutil.ReadAll(r)
	output := string(out)
	os.Stdout = stdout

	if output != expectedOutput {
		t.Fatalf("invalid output on the STDOUT, got=\"%s\", expected=\"%s\"", output, expectedOutput)
	}
}

func TestSyscallWritesChar(t *testing.T) {
	input := stringToUint16("h")
	expectedOutput := "h"

	// Rewrite stdout pipe
	stdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	WriteChar(input[0])

	// Revert to previous state
	w.Close()
	out, _ := ioutil.ReadAll(r)
	output := string(out)
	os.Stdout = stdout

	if output != expectedOutput {
		t.Fatalf("invalid output on the STDOUT, got=\"%s\", expected=\"%s\"", output, expectedOutput)
	}
}
