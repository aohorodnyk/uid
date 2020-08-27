package intl

import "encoding/base32"

func NewEncoder() Encoder {
	return &Encoding{
		Encoder: base32.StdEncoding.WithPadding(base32.NoPadding),
	}
}

func NewEncoderCustom(encoder Encoder) Encoder {
	return &Encoding{
		Encoder: encoder,
	}
}
