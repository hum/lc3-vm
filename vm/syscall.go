package vm

import (
	"bufio"
	"os"
)

func Uint16ToString(arr []uint16) string {
	var result string

	for _, v := range arr {
		result += string(rune(v))
	}
	return result
}

func GetChar() (uint16, error) {
	// Reads a single character (rune) from STDIN
	r := bufio.NewReader(os.Stdin)
	ch, _, err := r.ReadRune()
	if err != nil {
		panic(err)
	}
	return uint16(ch), nil
}

func WriteChar(ch uint16) {
	w := bufio.NewWriter(os.Stdout)
	_, err := w.WriteRune(rune(ch))
	if err != nil {
		panic(err)
	}
	w.Flush()
}

func WriteString(arr []uint16) {
	w := bufio.NewWriter(os.Stdout)
	result := Uint16ToString(arr)
	_, err := w.WriteString(result)
	if err != nil {
		panic(err)
	}
	err = w.Flush()
	if err != nil {
		panic(err)
	}
}
