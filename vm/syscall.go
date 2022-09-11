package vm

import (
	"bufio"
	"os"
)

func GetChar() error {
	// Reads a single character (rune) from STDIN
	r := bufio.NewReader(os.Stdin)
	ch, _, err := r.ReadRune()
	if err != nil {
		panic(err)
	}
	return ch, nil
}

func WriteChar(ch rune) {
	// Write a single character to STDOUT
	w := bufio.NewWriter(os.Stdout)
	_, err := w.WriteRune(ch)
	if err != nil {
		panic(err)
	}
}
