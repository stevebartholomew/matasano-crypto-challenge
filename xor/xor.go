package xor

func SingleByte(a []byte, b []byte) []byte {
	bytes := make([]byte, len(a))

	for i:=0; i<len(a); i++ {
		bytes[i] = a[i] ^ b[i]
	}

	return bytes
}
