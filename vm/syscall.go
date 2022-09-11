package vm

import (
	"bufio"
	"fmt"
	"os"
)

func GetChar() (rune, error) {
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

func WriteString(bb []uint16) {
	// Write a slice of values to STDOUT
	for b := range bb {
		fmt.Printf("%c", rune(b))
	}
	fmt.Printf("\n")
}
