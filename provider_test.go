package uid

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProvidingGenerateError(t *testing.T) {
	p := NewProviderCustom(0, NewRand(), NewEncoder())
	b, err := p.Generate()
	assert.Equal(t, nil, b)

	_, ok := err.(*RandError)
	assert.True(t, ok)
}

func TestProvidingDecodeError(t *testing.T) {
	p := NewProvider()
	b, err := p.Parse("01")
	assert.Equal(t, nil, b)

	_, ok := err.(*EncError)
	assert.True(t, ok)
}

func TestProvidingRandom(t *testing.T) {
	for idx, size := range providerRandom() {
		t.Run(fmt.Sprintf("test_providing_random_%d", idx), func(t *testing.T) {
			prov := NewProviderSize(size)
			cmp := assert.New(t)

			idSource, err := prov.Generate()
			cmp.Nil(err)

			id1 := prov.Byte(idSource.Byte())
			compareID(t, idSource, id1)

			id2, err := prov.Parse(id1.String())
			cmp.Nil(err)
			compareID(t, idSource, id2)
			compareID(t, id1, id2)

			int16s, err := id2.Int16()
			cmp.Nil(err)
			id3 := prov.Int16(int16s)
			compareID(t, idSource, id3)
			compareID(t, id2, id3)

			uint16s, err := id3.Uint16()
			cmp.Nil(err)
			id4 := prov.Uint16(uint16s)
			compareID(t, idSource, id4)
			compareID(t, id3, id4)

			int32s, err := id4.Int32()
			cmp.Nil(err)
			id5 := prov.Int32(int32s)
			compareID(t, idSource, id5)
			compareID(t, id4, id5)

			uint32s, err := id4.Uint32()
			cmp.Nil(err)
			id6 := prov.Uint32(uint32s)
			compareID(t, idSource, id6)
			compareID(t, id5, id6)

			int64s, err := id4.Int64()
			cmp.Nil(err)
			id7 := prov.Int64(int64s)
			compareID(t, idSource, id7)
			compareID(t, id6, id7)

			uint64s, err := id4.Uint64()
			cmp.Nil(err)
			id8 := prov.Uint64(uint64s)
			compareID(t, idSource, id8)
			compareID(t, id7, id8)
		})
	}
}

func compareID(t *testing.T, id1, id2 Identifier) {
	cmp := assert.New(t)

	i16s, err := id1.Int16()
	cmp.Nil(err)
	ui16s, err := id1.Uint16()
	cmp.Nil(err)
	i32s, err := id1.Int32()
	cmp.Nil(err)
	ui32s, err := id1.Uint32()
	cmp.Nil(err)
	i64s, err := id1.Int64()
	cmp.Nil(err)
	ui64s, err := id1.Uint64()
	cmp.Nil(err)

	cmp.Equal(id1.Byte(), id2.Byte())
	cmp.Equal(id1.String(), id2.String())

	i16d, err := id2.Int16()
	cmp.Nil(err)
	cmp.Equal(i16s, i16d)

	ui16d, err := id2.Uint16()
	cmp.Nil(err)
	cmp.Equal(ui16s, ui16d)

	i32d, err := id2.Int32()
	cmp.Nil(err)
	cmp.Equal(i32s, i32d)

	ui32d, err := id2.Uint32()
	cmp.Nil(err)
	cmp.Equal(ui32s, ui32d)

	i64d, err := id2.Int64()
	cmp.Nil(err)
	cmp.Equal(i64s, i64d)

	ui64d, err := id2.Uint64()
	cmp.Nil(err)
	cmp.Equal(ui64s, ui64d)
}

func providerRandom() []uint32 {
	res := make([]uint32, 0, 16)
	for i := 8; i <= 128; i++ {
		if i%8 == 0 {
			res = append(res, uint32(i))
		}
	}

	return res
}
