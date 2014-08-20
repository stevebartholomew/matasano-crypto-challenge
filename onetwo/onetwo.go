package onetwo

import (
	"encoding/hex"
)

func FixedXOR(a []byte, b []byte) []byte {
	bytes := make([]byte, len(a))

	for i:=0; i<len(a); i++ {
		bytes[i] = a[i] ^ b[i]
	}

	return bytes
}

func Run() {
	expectedString := "746865206b696420646f6e277420706c6179"

	a, _   := hex.DecodeString("1c0111001f010100061a024b53535009181c")
	b, _ := hex.DecodeString("686974207468652062756c6c277320657965")

	result := FixedXOR(a, b)

	if(hex.EncodeToString(result) == expectedString) {
		print("Hooray!")
	}
}
