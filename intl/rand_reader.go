package intl

import cryptRand "crypto/rand"

type CryptRandReader struct{}

func (r CryptRandReader) Read(b []byte) (n int, err error) {
	return cryptRand.Read(b)
}
