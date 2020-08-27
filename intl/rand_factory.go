package intl

import "io"

func NewRand() Randomizer {
	return &Rand{
		reader: &CryptRandReader{},
	}
}

func NewRandCustom(r io.Reader) Randomizer {
	return &Rand{
		reader: r,
	}
}
