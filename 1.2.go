package main

import (
	"encoding/hex"
	"github.com/stevebartholomew/matasano-crypto-challenge/xor"
)

func main() {
	expectedString := "746865206b696420646f6e277420706c6179"

	a, _   := hex.DecodeString("1c0111001f010100061a024b53535009181c")
	b, _ := hex.DecodeString("686974207468652062756c6c277320657965")

	result := xor.SingleByte(a, b)

	if(hex.EncodeToString(result) == expectedString) {
		print("Hooray!")
	}
}
