package vm

import (
	"os"

	"github.com/eiannone/keyboard"
)

func readRuneFromInput() uint16 {
	ch, k, err := keyboard.GetSingleKey()
	if err != nil {
		panic(err)
	}
	if k == keyboard.KeyEsc || k == keyboard.KeyCtrlC {
		// TODO:
		// Handle interrupt in a better way
		os.Exit(1)
	}
	return uint16(ch)
}
