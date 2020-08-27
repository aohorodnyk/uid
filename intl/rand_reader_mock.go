package intl

type MockRandReader struct {
	Actual    []byte
	ActualErr error
}

func (r MockRandReader) Read(b []byte) (n int, err error) {
	if r.ActualErr != nil {
		return 0, r.ActualErr
	}

	for i := range b {
		b[i] = r.Actual[i]
	}
	return len(r.Actual), nil
}
