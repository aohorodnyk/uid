package uidgen

import "github.com/aohorodnyk/uidgen/intl"

func NewProvider() intl.Provider {
	return intl.NewProvider()
}

func NewProviderSize(size uint32) intl.Provider {
	return intl.NewProviderSize(size)
}

func NewProviderCustom(size uint32, rand intl.Randomizer, enc intl.Encoder) intl.Provider {
	return intl.NewProviderCustom(size, rand, enc)
}

func NewID(data []byte, enc intl.Encoder) intl.Identifier {
	return intl.NewID(data, enc)
}

func NewIDStdBase32(data []byte) intl.Identifier {
	return intl.NewIDStdBase32(data)
}
