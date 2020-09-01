package intl

import (
	"errors"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBase16Encoding(t *testing.T) {
	bytes := []byte{254, 152}
	enc := NewEncoderBase16Std()
	actEnc := enc.EncodeToString(bytes)
	actDec, err := enc.DecodeString(actEnc)

	assert.Equal(t, "fe98", actEnc)
	assert.Nil(t, err)
	assert.Equal(t, bytes, actDec)
}

func TestBase32StdEncoding(t *testing.T) {
	bytes := []byte{254, 152}
	enc := NewEncoderBase32Std()
	actEnc := enc.EncodeToString(bytes)
	actDec, err := enc.DecodeString(actEnc)

	assert.Equal(t, "72MA", actEnc)
	assert.Nil(t, err)
	assert.Equal(t, bytes, actDec)
}

func TestBase32HexEncoding(t *testing.T) {
	bytes := []byte{254, 152}
	enc := NewEncoderBase32Hex()
	actEnc := enc.EncodeToString(bytes)
	actDec, err := enc.DecodeString(actEnc)

	assert.Equal(t, "VQC0", actEnc)
	assert.Nil(t, err)
	assert.Equal(t, bytes, actDec)
}

func TestBase64UrlEncoding(t *testing.T) {
	bytes := []byte{254, 152}
	enc := NewEncoderBase64Url()
	actEnc := enc.EncodeToString(bytes)
	actDec, err := enc.DecodeString(actEnc)

	assert.Equal(t, "_pg", actEnc)
	assert.Nil(t, err)
	assert.Equal(t, bytes, actDec)
}

func TestBase64StdEncoding(t *testing.T) {
	bytes := []byte{254, 152}
	enc := NewEncoderBase64Std()
	actEnc := enc.EncodeToString(bytes)
	actDec, err := enc.DecodeString(actEnc)

	assert.Equal(t, "/pg", actEnc)
	assert.Nil(t, err)
	assert.Equal(t, bytes, actDec)
}

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
