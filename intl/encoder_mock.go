package intl

type MockEncoding struct {
	EncodedString string
	DecodedBytes  []byte
	DecodedErr    error
}

func (e MockEncoding) EncodeToString(src []byte) string {
	return e.EncodedString
}

func (e MockEncoding) DecodeString(s string) ([]byte, error) {
	if e.DecodedErr != nil {
		return nil, e.DecodedErr
	}

	return e.DecodedBytes, nil
}
