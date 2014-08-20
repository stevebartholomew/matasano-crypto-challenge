package onethree

import (
	"encoding/hex"
	"fmt"
	"github.com/stevebartholomew/matasano-crypto-challenge/onetwo"
)

func Run() {
	input, _ := hex.DecodeString("1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736")

	for charIndex := 0; charIndex < 256; charIndex++ {
		keyBytes := make([]byte, len(input))

		for i := 0; i < len(input); i++ {
			keyBytes[i] = byte(charIndex)
		}

		fmt.Printf("%s\n", onetwo.FixedXOR(input, keyBytes))
	}
}
