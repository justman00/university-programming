package main

import (
	"fmt"
	"math/big"
	"os"

	"golang.org/x/term"
)

func main() {
	// switch stdin into 'raw' mode
	oldState, err := term.MakeRaw(int(os.Stdin.Fd()))
	if err != nil {
		fmt.Println(err)
		return
	}
	defer term.Restore(int(os.Stdin.Fd()), oldState)

	b := make([]byte, 1)
	for {
		if _, err := os.Stdin.Read(b); err != nil {
			panic(err)
		}
		charCode := int(big.NewInt(0).SetBytes(b).Uint64())

		if charCode == 3 {
			return
		}

		if string(b) != "B" {
			continue
		}

		if _, err := os.Stdout.Write([]byte{0x07}); err != nil {
			panic(err)
		}
	}
}
