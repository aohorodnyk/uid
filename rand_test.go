package uid

import (
	"errors"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRand_GenerateCryptReader(t *testing.T) {
	rand := NewRandCustom(&CryptRandReader{})
	provSrc := providerRandGenerateCryptReader()
	for idx, prov := range provSrc {
		t.Run(fmt.Sprintf("test_rand_generate_%d", idx), func(t *testing.T) {
			act, err := rand.Generate(prov.size)

			assert.Len(t, act, int(prov.expSize), "Size of bytes are not equal")
			assert.Equal(t, prov.expErr, err, "Errors are not equal")
		})
	}
}

func TestRand_GenerateInternalError(t *testing.T) {
	ieMsg := errors.New("internal error")
	mockReader := &MockRandReader{
		Actual:    nil,
		ActualErr: ieMsg,
	}
	rand := NewRandCustom(mockReader)
	act, err := rand.Generate(4)

	expErr := &InternalError{ieMsg}

	assert.Nil(t, act, "Actual value is not nil")
	assert.Equal(t, expErr, err, "Internal error is wrong")
}

func TestRand_GenerateMockData(t *testing.T) {
	mockReader := &MockRandReader{
		Actual: []byte{5, 2, 6, 7},
	}
	rand := NewRandCustom(mockReader)
	act, err := rand.Generate(4)

	assert.Equal(t, []byte{5, 2, 6, 7}, act)
	assert.Nil(t, err, "Actual error is not nil")
}

type providerRandGenerateType struct {
	size    uint32
	expSize uint32
	expErr  error
}

func providerRandGenerateCryptReader() []providerRandGenerateType {
	const cases = 33
	prov := make([]providerRandGenerateType, cases)
	prov[0] = providerRandGenerateType{
		expErr: &WrongSizeError{errors.New(ErrorMsgWrongSizeZero)},
	}

	for i := 1; i < cases; i++ {
		prov[i] = providerRandGenerateType{
			size:    uint32(i),
			expSize: uint32(i),
		}
	}

	return prov
}
