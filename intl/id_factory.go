package intl

func NewID(data []byte, enc Encoder) Identifier {
	return &ID{
		data: data,
		enc:  enc,
	}
}

func NewIDStdBase32(data []byte) Identifier {
	return &ID{
		data: data,
		enc:  NewEncoder(),
	}
}
