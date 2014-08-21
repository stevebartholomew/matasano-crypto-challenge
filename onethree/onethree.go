package onethree

import (
	"encoding/hex"
	"fmt"
	"strings"
	"github.com/stevebartholomew/matasano-crypto-challenge/onetwo"
)

var parts = []string{"ed", "ing", "th", "ey", "ent"}

func Rank(phrase string) int {
	rank := 0

	words := strings.Split(phrase, " ")

	for _, word := range words {
		for _, part := range parts {
			if strings.Contains(word, part) {
				print(word)
				rank++
			}
		}
	}

	return rank
}

func Run() {
	input, _ := hex.DecodeString("1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736")

	for charIndex := 0; charIndex < 256; charIndex++ {
		keyBytes := make([]byte, len(input))

		for i := 0; i < len(input); i++ {
			keyBytes[i] = byte(charIndex)
		}

		decodedPhrase := fmt.Sprintf("%s", onetwo.FixedXOR(input, keyBytes))

		rank := Rank(decodedPhrase)

		if rank > 0 {
			fmt.Println("%s\n", decodedPhrase)
		}
	}
}
