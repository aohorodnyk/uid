package intl

import "encoding/hex"

type Base16Encoding struct{}

func (b Base16Encoding) EncodeToString(src []byte) string {
	return hex.EncodeToString(src)
}

func (b Base16Encoding) DecodeString(s string) ([]byte, error) {
	return hex.DecodeString(s)
}
