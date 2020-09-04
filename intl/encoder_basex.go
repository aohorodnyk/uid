package intl

import "math/big"

type BaseXEncoding struct {
	base int
}

func (b BaseXEncoding) EncodeToString(src []byte) string {
	bi := new(big.Int)
	bi.SetBytes(src)

	return bi.Text(b.base)
}

func (b BaseXEncoding) DecodeString(s string) ([]byte, error) {
	bi := new(big.Int)
	bi.SetString(s, b.base)

	return bi.Bytes(), nil
}
