package intl

import (
	"errors"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestEncoderMockEncode(t *testing.T) {
	enc := &MockEncoding{
		EncodedString: "A",
	}
	e := NewEncoderCustom(enc)
	a := e.EncodeToString([]byte{})

	assert.Equal(t, "A", a)
}

func TestEncoderMockDecode(t *testing.T) {
	for idx, prov := range providerEncoderMockDecode() {
		t.Run(fmt.Sprintf("test_encoder_mock_decode_%d", idx), func(t *testing.T) {
			m := &MockEncoding{
				DecodedBytes: prov.DecodedBytes,
				DecodedErr:   prov.DecodedErr,
			}

			enc := NewEncoderCustom(m)
			b, err := enc.DecodeString("A")

			assert.Equal(t, err, prov.DecodedErr)
			assert.Equal(t, prov.DecodedBytes, b)
		})
	}
}

func providerEncoderMockDecode() []MockEncoding {
	return []MockEncoding{
		{
			DecodedErr: errors.New("test"),
		},
		{
			DecodedBytes: []byte{1, 5},
		},
	}
}
