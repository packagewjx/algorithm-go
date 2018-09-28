package util

type MString string

// Convert the string to byte array. The string must be all ascii character.
func (s MString) ToBytes() []byte {
	bytes := make([]byte, len(s))

	for i, c := range s {
		bytes[i] = byte(c)
	}
	return bytes
}
