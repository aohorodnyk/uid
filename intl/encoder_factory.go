package intl

import (
	"encoding/base32"
	"encoding/base64"
)

func NewEncoder() Encoder {
	return NewEncoderBase32Std()
}

func NewEncoderBase16Std() Encoder {
	return NewEncoderCustom(&Base16Encoding{})
}

func NewEncoderBase32Std() Encoder {
	return NewEncoderCustom(base32.StdEncoding.WithPadding(base32.NoPadding))
}

func NewEncoderBase32Hex() Encoder {
	return NewEncoderCustom(base32.HexEncoding.WithPadding(base32.NoPadding))
}

func NewEncoderBase64Std() Encoder {
	return NewEncoderCustom(base64.RawStdEncoding)
}

func NewEncoderBase64Url() Encoder {
	return NewEncoderCustom(base64.RawURLEncoding)
}

func NewEncoderCustom(encoder Encoder) Encoder {
	return &Encoding{
		Encoder: encoder,
	}
}
