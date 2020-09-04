package uidgen

import "github.com/aohorodnyk/uidgen/intl"

func NewProvider() intl.Provider {
	return intl.NewProvider()
}

func NewProviderSize(size uint32) intl.Provider {
	return intl.NewProviderSize(size)
}

func NewProvider36() intl.Provider {
	return intl.NewProviderCustom(intl.SizeDefault, intl.NewRand(), intl.NewEncoderBase36())
}

func NewProvider36Size(size uint32) intl.Provider {
	return intl.NewProviderCustom(size, intl.NewRand(), intl.NewEncoderBase36())
}

func NewProvider62() intl.Provider {
	return intl.NewProviderCustom(intl.SizeDefault, intl.NewRand(), intl.NewEncoderBase62())
}

func NewProvider62Size(size uint32) intl.Provider {
	return intl.NewProviderCustom(size, intl.NewRand(), intl.NewEncoderBase62())
}

func NewProviderUrl64() intl.Provider {
	return intl.NewProviderCustom(intl.SizeDefault, intl.NewRand(), intl.NewEncoderBase64Url())
}

func NewProviderUrl64Size(size uint32) intl.Provider {
	return intl.NewProviderCustom(size, intl.NewRand(), intl.NewEncoderBase64Url())
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

func NewIDBase36(data []byte) intl.Identifier {
	return intl.NewID(data, intl.NewEncoderBase36())
}

func NewIDBase62(data []byte) intl.Identifier {
	return intl.NewID(data, intl.NewEncoderBase62())
}

func NewIDBase64(data []byte) intl.Identifier {
	return intl.NewID(data, intl.NewEncoderBase62())
}
