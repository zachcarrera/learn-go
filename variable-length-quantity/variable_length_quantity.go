package variablelengthquantity

import "errors"

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
	var decoded []uint32
	var currentDecode uint32
	var length int
	for _, byte := range input {
		currentDecode += uint32(byte & 0x7F)
		if byte&0x80 == 0 {
			decoded = append(decoded, currentDecode)
			currentDecode, length = 0, 0
			continue
		}
		currentDecode <<= 7
		length++
		if length > 4 {
			return nil, errors.New("Sequence is too long")
		}
	}
	if length != 0 {
		return nil, errors.New("Incomplete byte sequence")
	}
	return decoded, nil
}
