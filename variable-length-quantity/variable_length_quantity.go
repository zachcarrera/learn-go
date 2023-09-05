package variablelengthquantity

func EncodeVarint(input []uint32) []byte {
	var encoded []byte
	for _, v := range input {
		// encode the first 7 bits
		singleEncode := []byte{byte(v % 128)}

		// shift the input number by seven bits and prepend it to singleEncode with the first bit set
		for v >>= 7; v != 0; v >>= 7 {
			singleEncode = append([]byte{128 + byte(v%128)}, singleEncode...)
		}
		encoded = append(encoded, singleEncode...)
	}
	return encoded
}

func DecodeVarint(input []byte) ([]uint32, error) {
	panic("Please implement the DecodeVarint function")
}
