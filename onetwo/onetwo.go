package onetwo

import (
	"encoding/hex"
)

func Run() {
	expectedString := "746865206b696420646f6e277420706c6179"

	input, _   := hex.DecodeString("1c0111001f010100061a024b53535009181c")
	against, _ := hex.DecodeString("686974207468652062756c6c277320657965")

	bytes := make([]byte, len(input))

	for i:=0; i<len(input); i++ {
		bytes[i] = against[i] ^ input[i]
	}

	if(hex.EncodeToString(bytes) == expectedString) {
		print("Hooray!")
	}
}
