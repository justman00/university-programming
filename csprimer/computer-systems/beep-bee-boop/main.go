package main

import (
	"fmt"
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
		data := int(b[0])

		if data == 3 {
			return
		}

		if data != 66 {
			continue
		}

		if _, err := os.Stdout.Write([]byte{0x07}); err != nil {
			panic(err)
		}
	}
}
