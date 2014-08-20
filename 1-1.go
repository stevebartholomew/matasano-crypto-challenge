package main

import (
	"encoding/base64"
	"encoding/hex"
	"fmt"
)

func main() {
	srcString := "49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d"

	srcBytes, _ := hex.DecodeString(srcString)

	generatedBase64 := base64.StdEncoding.EncodeToString(srcBytes)

	fmt.Println(generatedBase64)
}
