package onethree

import (
	"encoding/hex"
	"fmt"
)

func Run() {
	//input, _ := hex.DecodeString("1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736")
	input := []byte("1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736")

	for charIndex := 0; charIndex < 256; charIndex++ {
		char := []byte(fmt.Sprintf("%c ", charIndex))

		results := make([]byte, len(input))

		for i := 0; i < len(input); i++ {
			results[i] = input[i] ^ char[0]
		}

		output := make([]byte, len(input))

		hex.Decode(output, results)

		fmt.Println(string(output))
	}
}
