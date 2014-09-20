package onethree

import (
	"encoding/hex"
	_"encoding/binary"
	"fmt"
	"strings"
	"github.com/stevebartholomew/matasano-crypto-challenge/onetwo"
	"os/exec"
	"strconv"
	"regexp"
)

const DICT_PATH = "/usr/share/dict/words"

func Rank(phrase string) int {
	rank := 0

	words := strings.Split(phrase, " ")
	for _, word := range words {
		isAlphaNumeric, _ := regexp.Match("[A-Za-z0-9]+", []byte(word))

		if isAlphaNumeric {
			occurances, _ := exec.Command("grep", "-wciF", string(word), DICT_PATH).Output()
			i, _ := strconv.Atoi(strings.TrimSpace(string(occurances)))

			if i > 0 {
				rank += i
			}
		}
	}

	return rank
}

func Run() {
	input, _ := hex.DecodeString("1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736")
	results := make(map[string]int)

	for charIndex := 0; charIndex < 256; charIndex++ {
		keyBytes := make([]byte, len(input))

		for i := 0; i < len(input); i++ {
			keyBytes[i] = byte(charIndex)
		}

		decodedPhrase := fmt.Sprintf("%s", onetwo.FixedXOR(input, keyBytes))

		rank := Rank(decodedPhrase)

		if rank > 0 {
			results[decodedPhrase] = rank
		}
	}

	topPhrase := ""
	topRank := 0

	for phrase, rank := range results {
		if rank > topRank {
			topPhrase = phrase
			topRank = rank
		}
	}

	fmt.Println(topPhrase)
}
