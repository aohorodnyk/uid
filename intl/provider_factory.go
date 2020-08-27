package intl

func NewProvider() Provider {
	return &Providing{
		randSize: 20,
		rand:     NewRand(),
		enc:      NewEncoder(),
	}
}

func NewProviderSize(size uint32) Provider {
	return &Providing{
		randSize: size,
		rand:     NewRand(),
		enc:      NewEncoder(),
	}
}

func NewProviderCustom(size uint32, rand Randomizer, enc Encoder) Provider {
	return &Providing{
		randSize: size,
		rand:     rand,
		enc:      enc,
	}
}
