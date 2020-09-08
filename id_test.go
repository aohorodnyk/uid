package intl

import (
	"errors"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestReadBytesEmpty(t *testing.T) {
	bs := readBytes([]byte{}, 1, 1)
	assert.Equal(t, [][]byte{}, bs)
}
func TestReadBytesWrongSize(t *testing.T) {
	bs := readBytes([]byte{1, 2, 3}, 1, 2)
	assert.Equal(t, [][]byte{}, bs)
}
func TestNewArray(t *testing.T) {
	for idx, prov := range providerTestNewArray() {
		t.Run(fmt.Sprintf("test_new_array_%d", idx), func(t *testing.T) {
			src := make([]byte, prov.srcSize)
			act := sizeNewArray(src, prov.size)

			assert.Equal(t, prov.exp, act)
		})
	}
}

type providerTestNewArrayType struct {
	srcSize int
	size    int
	exp     int
}

func providerTestNewArray() []providerTestNewArrayType {
	return []providerTestNewArrayType{
		{
			srcSize: 1,
			size:    2,
			exp:     1,
		},
		{
			srcSize: 2,
			size:    8,
			exp:     1,
		},
		{
			srcSize: 12,
			size:    16,
			exp:     1,
		},
		{
			srcSize: 4,
			size:    4,
			exp:     1,
		},
		{
			srcSize: 6,
			size:    4,
			exp:     2,
		},
		{
			srcSize: 8,
			size:    4,
			exp:     2,
		},
		{
			srcSize: 9,
			size:    4,
			exp:     3,
		},
		{
			srcSize: 12,
			size:    4,
			exp:     3,
		},
		{
			srcSize: 12,
			size:    8,
			exp:     2,
		},
	}
}

func TestID(t *testing.T) {
	for idx, prov := range providerID() {
		t.Run(fmt.Sprintf("test_id_string_%d", idx), func(t *testing.T) {
			id := NewIDStdBase32(prov.act)
			s := id.String()
			b := id.Byte()

			assert.Equal(t, prov.act, b, "Bytes are not equal")
			assert.Equal(t, prov.expStr, s, "Base32 strings are not equal")

			i16, i16Err := id.Int16()
			assert.Equal(t, prov.expI16, i16, "Int16 values are not equal")
			assert.Equal(t, prov.expI16Err, i16Err, "Int16 errors are not equal")

			ui16, ui16Err := id.Uint16()
			assert.Equal(t, prov.expUI16, ui16, "Uint16 values are not equal")
			assert.Equal(t, prov.expUI16Err, ui16Err, "Uint16 errors are not equal")

			i32, i32Err := id.Int32()
			assert.Equal(t, prov.expI32, i32, "Int32 values are not equal")
			assert.Equal(t, prov.expI32Err, i32Err, "Int32 errors are not equal")

			ui32, ui32Err := id.Uint32()
			assert.Equal(t, prov.expUI32, ui32, "Uint32 values are not equal")
			assert.Equal(t, prov.expUI32Err, ui32Err, "Uint32 errors are not equal")

			i64, i64Err := id.Int64()
			assert.Equal(t, prov.expI64, i64, "Int64 values are not equal")
			assert.Equal(t, prov.expI64Err, i64Err, "Int64 errors are not equal")

			ui64, ui64Err := id.Uint64()
			assert.Equal(t, prov.expUI64, ui64, "Int64 values are not equal")
			assert.Equal(t, prov.expUI64Err, ui64Err, "Int64 errors are not equal")
		})
	}
}

type providerTestID struct {
	act        []byte
	expStr     string
	expI16     []int16
	expI16Err  error
	expUI16    []uint16
	expUI16Err error
	expI32     []int32
	expI32Err  error
	expUI32    []uint32
	expUI32Err error
	expI64     []int64
	expI64Err  error
	expUI64    []uint64
	expUI64Err error
}

func providerID() []providerTestID {
	return []providerTestID{
		{
			act:        []byte{0},
			expStr:     "AA",
			expI16:     nil,
			expI16Err:  &WrongSizeError{errors.New(ErrorMsgSizeDivisible2)},
			expUI16:    nil,
			expUI16Err: &WrongSizeError{errors.New(ErrorMsgSizeDivisible2)},
			expI32:     nil,
			expI32Err:  &WrongSizeError{errors.New(ErrorMsgSizeDivisible4)},
			expUI32:    nil,
			expUI32Err: &WrongSizeError{errors.New(ErrorMsgSizeDivisible4)},
			expI64:     nil,
			expI64Err:  &WrongSizeError{errors.New(ErrorMsgSizeDivisible8)},
			expUI64:    nil,
			expUI64Err: &WrongSizeError{errors.New(ErrorMsgSizeDivisible8)},
		},
		{
			act:        []byte{31},
			expStr:     "D4",
			expI16:     nil,
			expI16Err:  &WrongSizeError{errors.New(ErrorMsgSizeDivisible2)},
			expUI16:    nil,
			expUI16Err: &WrongSizeError{errors.New(ErrorMsgSizeDivisible2)},
			expI32:     nil,
			expI32Err:  &WrongSizeError{errors.New(ErrorMsgSizeDivisible4)},
			expUI32:    nil,
			expUI32Err: &WrongSizeError{errors.New(ErrorMsgSizeDivisible4)},
			expI64:     nil,
			expI64Err:  &WrongSizeError{errors.New(ErrorMsgSizeDivisible8)},
			expUI64:    nil,
			expUI64Err: &WrongSizeError{errors.New(ErrorMsgSizeDivisible8)},
		},
		{
			act:        []byte{82},
			expStr:     "KI",
			expI16:     nil,
			expI16Err:  &WrongSizeError{errors.New(ErrorMsgSizeDivisible2)},
			expUI16:    nil,
			expUI16Err: &WrongSizeError{errors.New(ErrorMsgSizeDivisible2)},
			expI32:     nil,
			expI32Err:  &WrongSizeError{errors.New(ErrorMsgSizeDivisible4)},
			expUI32:    nil,
			expUI32Err: &WrongSizeError{errors.New(ErrorMsgSizeDivisible4)},
			expI64:     nil,
			expI64Err:  &WrongSizeError{errors.New(ErrorMsgSizeDivisible8)},
			expUI64:    nil,
			expUI64Err: &WrongSizeError{errors.New(ErrorMsgSizeDivisible8)},
		},
		{
			act:        []byte{151},
			expStr:     "S4",
			expI16:     nil,
			expI16Err:  &WrongSizeError{errors.New(ErrorMsgSizeDivisible2)},
			expUI16:    nil,
			expUI16Err: &WrongSizeError{errors.New(ErrorMsgSizeDivisible2)},
			expI32:     nil,
			expI32Err:  &WrongSizeError{errors.New(ErrorMsgSizeDivisible4)},
			expUI32:    nil,
			expUI32Err: &WrongSizeError{errors.New(ErrorMsgSizeDivisible4)},
			expI64:     nil,
			expI64Err:  &WrongSizeError{errors.New(ErrorMsgSizeDivisible8)},
			expUI64:    nil,
			expUI64Err: &WrongSizeError{errors.New(ErrorMsgSizeDivisible8)},
		},
		{
			act:        []byte{255},
			expStr:     "74",
			expI16:     nil,
			expI16Err:  &WrongSizeError{errors.New(ErrorMsgSizeDivisible2)},
			expUI16:    nil,
			expUI16Err: &WrongSizeError{errors.New(ErrorMsgSizeDivisible2)},
			expI32:     nil,
			expI32Err:  &WrongSizeError{errors.New(ErrorMsgSizeDivisible4)},
			expUI32:    nil,
			expUI32Err: &WrongSizeError{errors.New(ErrorMsgSizeDivisible4)},
			expI64:     nil,
			expI64Err:  &WrongSizeError{errors.New(ErrorMsgSizeDivisible8)},
			expUI64:    nil,
			expUI64Err: &WrongSizeError{errors.New(ErrorMsgSizeDivisible8)},
		},
		{
			act:        []byte{0, 0},
			expStr:     "AAAA",
			expI16:     []int16{0},
			expI16Err:  nil,
			expUI16:    []uint16{0},
			expUI16Err: nil,
			expI32:     nil,
			expI32Err:  &WrongSizeError{errors.New(ErrorMsgSizeDivisible4)},
			expUI32:    nil,
			expUI32Err: &WrongSizeError{errors.New(ErrorMsgSizeDivisible4)},
			expI64:     nil,
			expI64Err:  &WrongSizeError{errors.New(ErrorMsgSizeDivisible8)},
			expUI64:    nil,
			expUI64Err: &WrongSizeError{errors.New(ErrorMsgSizeDivisible8)},
		},
		{
			act:        []byte{250, 159},
			expStr:     "7KPQ",
			expI16:     []int16{-24582},
			expI16Err:  nil,
			expUI16:    []uint16{40954},
			expUI16Err: nil,
			expI32:     nil,
			expI32Err:  &WrongSizeError{errors.New(ErrorMsgSizeDivisible4)},
			expUI32:    nil,
			expUI32Err: &WrongSizeError{errors.New(ErrorMsgSizeDivisible4)},
			expI64:     nil,
			expI64Err:  &WrongSizeError{errors.New(ErrorMsgSizeDivisible8)},
			expUI64:    nil,
			expUI64Err: &WrongSizeError{errors.New(ErrorMsgSizeDivisible8)},
		},
		{
			act:        []byte{223, 75},
			expStr:     "35FQ",
			expI16:     []int16{19423},
			expI16Err:  nil,
			expUI16:    []uint16{19423},
			expUI16Err: nil,
			expI32:     nil,
			expI32Err:  &WrongSizeError{errors.New(ErrorMsgSizeDivisible4)},
			expUI32:    nil,
			expUI32Err: &WrongSizeError{errors.New(ErrorMsgSizeDivisible4)},
			expI64:     nil,
			expI64Err:  &WrongSizeError{errors.New(ErrorMsgSizeDivisible8)},
			expUI64:    nil,
			expUI64Err: &WrongSizeError{errors.New(ErrorMsgSizeDivisible8)},
		},
		{
			act:        []byte{181, 41},
			expStr:     "WUUQ",
			expI16:     []int16{10677},
			expI16Err:  nil,
			expUI16:    []uint16{10677},
			expUI16Err: nil,
			expI32:     nil,
			expI32Err:  &WrongSizeError{errors.New(ErrorMsgSizeDivisible4)},
			expUI32:    nil,
			expUI32Err: &WrongSizeError{errors.New(ErrorMsgSizeDivisible4)},
			expI64:     nil,
			expI64Err:  &WrongSizeError{errors.New(ErrorMsgSizeDivisible8)},
			expUI64:    nil,
			expUI64Err: &WrongSizeError{errors.New(ErrorMsgSizeDivisible8)},
		},
		{
			act:        []byte{255, 255},
			expStr:     "777Q",
			expI16:     []int16{-1},
			expI16Err:  nil,
			expUI16:    []uint16{65535},
			expUI16Err: nil,
			expI32:     nil,
			expI32Err:  &WrongSizeError{errors.New(ErrorMsgSizeDivisible4)},
			expUI32:    nil,
			expUI32Err: &WrongSizeError{errors.New(ErrorMsgSizeDivisible4)},
			expI64:     nil,
			expI64Err:  &WrongSizeError{errors.New(ErrorMsgSizeDivisible8)},
			expUI64:    nil,
			expUI64Err: &WrongSizeError{errors.New(ErrorMsgSizeDivisible8)},
		},
		{
			act:        []byte{0, 0, 0},
			expStr:     "AAAAA",
			expI16:     nil,
			expI16Err:  &WrongSizeError{errors.New(ErrorMsgSizeDivisible2)},
			expUI16:    nil,
			expUI16Err: &WrongSizeError{errors.New(ErrorMsgSizeDivisible2)},
			expI32:     nil,
			expI32Err:  &WrongSizeError{errors.New(ErrorMsgSizeDivisible4)},
			expUI32:    nil,
			expUI32Err: &WrongSizeError{errors.New(ErrorMsgSizeDivisible4)},
			expI64:     nil,
			expI64Err:  &WrongSizeError{errors.New(ErrorMsgSizeDivisible8)},
			expUI64:    nil,
			expUI64Err: &WrongSizeError{errors.New(ErrorMsgSizeDivisible8)},
		},
		{
			act:        []byte{14, 141, 223},
			expStr:     "B2G56",
			expI16:     nil,
			expI16Err:  &WrongSizeError{errors.New(ErrorMsgSizeDivisible2)},
			expUI16:    nil,
			expUI16Err: &WrongSizeError{errors.New(ErrorMsgSizeDivisible2)},
			expI32:     nil,
			expI32Err:  &WrongSizeError{errors.New(ErrorMsgSizeDivisible4)},
			expUI32:    nil,
			expUI32Err: &WrongSizeError{errors.New(ErrorMsgSizeDivisible4)},
			expI64:     nil,
			expI64Err:  &WrongSizeError{errors.New(ErrorMsgSizeDivisible8)},
			expUI64:    nil,
			expUI64Err: &WrongSizeError{errors.New(ErrorMsgSizeDivisible8)},
		},
		{
			act:        []byte{207, 41, 92},
			expStr:     "Z4UVY",
			expI16:     nil,
			expI16Err:  &WrongSizeError{errors.New(ErrorMsgSizeDivisible2)},
			expUI16:    nil,
			expUI16Err: &WrongSizeError{errors.New(ErrorMsgSizeDivisible2)},
			expI32:     nil,
			expI32Err:  &WrongSizeError{errors.New(ErrorMsgSizeDivisible4)},
			expUI32:    nil,
			expUI32Err: &WrongSizeError{errors.New(ErrorMsgSizeDivisible4)},
			expI64:     nil,
			expI64Err:  &WrongSizeError{errors.New(ErrorMsgSizeDivisible8)},
			expUI64:    nil,
			expUI64Err: &WrongSizeError{errors.New(ErrorMsgSizeDivisible8)},
		},
		{
			act:        []byte{11, 168, 237},
			expStr:     "BOUO2",
			expI16:     nil,
			expI16Err:  &WrongSizeError{errors.New(ErrorMsgSizeDivisible2)},
			expUI16:    nil,
			expUI16Err: &WrongSizeError{errors.New(ErrorMsgSizeDivisible2)},
			expI32:     nil,
			expI32Err:  &WrongSizeError{errors.New(ErrorMsgSizeDivisible4)},
			expUI32:    nil,
			expUI32Err: &WrongSizeError{errors.New(ErrorMsgSizeDivisible4)},
			expI64:     nil,
			expI64Err:  &WrongSizeError{errors.New(ErrorMsgSizeDivisible8)},
			expUI64:    nil,
			expUI64Err: &WrongSizeError{errors.New(ErrorMsgSizeDivisible8)},
		},
		{
			act:        []byte{6, 172, 142, 121},
			expStr:     "A2WI46I",
			expI16:     []int16{-21498, 31118},
			expI16Err:  nil,
			expUI16:    []uint16{44038, 31118},
			expUI16Err: nil,
			expI32:     []int32{2039393286},
			expI32Err:  nil,
			expUI32:    []uint32{2039393286},
			expUI32Err: nil,
			expI64:     nil,
			expI64Err:  &WrongSizeError{errors.New(ErrorMsgSizeDivisible8)},
			expUI64:    nil,
			expUI64Err: &WrongSizeError{errors.New(ErrorMsgSizeDivisible8)},
		},
		{
			act:        []byte{196, 243, 233, 154},
			expStr:     "YTZ6TGQ",
			expI16:     []int16{-3132, -25879},
			expI16Err:  nil,
			expUI16:    []uint16{62404, 39657},
			expUI16Err: nil,
			expI32:     []int32{-1695943740},
			expI32Err:  nil,
			expUI32:    []uint32{2599023556},
			expUI32Err: nil,
			expI64:     nil,
			expI64Err:  &WrongSizeError{errors.New(ErrorMsgSizeDivisible8)},
			expUI64:    nil,
			expUI64Err: &WrongSizeError{errors.New(ErrorMsgSizeDivisible8)},
		},
		{
			act:        []byte{102, 194, 61, 156},
			expStr:     "M3BD3HA",
			expI16:     []int16{-15770, -25539},
			expI16Err:  nil,
			expUI16:    []uint16{49766, 39997},
			expUI16Err: nil,
			expI32:     []int32{-1673674138},
			expI32Err:  nil,
			expUI32:    []uint32{2621293158},
			expUI32Err: nil,
			expI64:     nil,
			expI64Err:  &WrongSizeError{errors.New(ErrorMsgSizeDivisible8)},
			expUI64:    nil,
			expUI64Err: &WrongSizeError{errors.New(ErrorMsgSizeDivisible8)},
		},
		{
			act:        []byte{247, 119, 29, 217, 99},
			expStr:     "653R3WLD",
			expI16:     nil,
			expI16Err:  &WrongSizeError{errors.New(ErrorMsgSizeDivisible2)},
			expUI16:    nil,
			expUI16Err: &WrongSizeError{errors.New(ErrorMsgSizeDivisible2)},
			expI32:     nil,
			expI32Err:  &WrongSizeError{errors.New(ErrorMsgSizeDivisible4)},
			expUI32:    nil,
			expUI32Err: &WrongSizeError{errors.New(ErrorMsgSizeDivisible4)},
			expI64:     nil,
			expI64Err:  &WrongSizeError{errors.New(ErrorMsgSizeDivisible8)},
			expUI64:    nil,
			expUI64Err: &WrongSizeError{errors.New(ErrorMsgSizeDivisible8)},
		},
		{
			act:        []byte{18, 90, 248, 13, 160, 130},
			expStr:     "CJNPQDNAQI",
			expI16:     []int16{23058, 3576, -32096},
			expI16Err:  nil,
			expUI16:    []uint16{23058, 3576, 33440},
			expUI16Err: nil,
			expI32:     nil,
			expI32Err:  &WrongSizeError{errors.New(ErrorMsgSizeDivisible4)},
			expUI32:    nil,
			expUI32Err: &WrongSizeError{errors.New(ErrorMsgSizeDivisible4)},
			expI64:     nil,
			expI64Err:  &WrongSizeError{errors.New(ErrorMsgSizeDivisible8)},
			expUI64:    nil,
			expUI64Err: &WrongSizeError{errors.New(ErrorMsgSizeDivisible8)},
		},
		{
			act:        []byte{122, 86, 150, 249, 88, 130, 236},
			expStr:     "PJLJN6KYQLWA",
			expI16:     nil,
			expI16Err:  &WrongSizeError{errors.New(ErrorMsgSizeDivisible2)},
			expUI16:    nil,
			expUI16Err: &WrongSizeError{errors.New(ErrorMsgSizeDivisible2)},
			expI32:     nil,
			expI32Err:  &WrongSizeError{errors.New(ErrorMsgSizeDivisible4)},
			expUI32:    nil,
			expUI32Err: &WrongSizeError{errors.New(ErrorMsgSizeDivisible4)},
			expI64:     nil,
			expI64Err:  &WrongSizeError{errors.New(ErrorMsgSizeDivisible8)},
			expUI64:    nil,
			expUI64Err: &WrongSizeError{errors.New(ErrorMsgSizeDivisible8)},
		},
		{
			act:        []byte{96, 128, 217, 59, 230, 22, 25, 224},
			expStr:     "MCANSO7GCYM6A",
			expI16:     []int16{-32672, 15321, 5862, -8167},
			expI16Err:  nil,
			expUI16:    []uint16{32864, 15321, 5862, 57369},
			expUI16Err: nil,
			expI32:     []int32{1004109920, -535226650},
			expI32Err:  nil,
			expUI32:    []uint32{1004109920, 3759740646},
			expUI32Err: nil,
			expI64:     []int64{-2298780956693528480},
			expI64Err:  nil,
			expUI64:    []uint64{16147963117016023136},
			expUI64Err: nil,
		},
		{
			act:        []byte{116, 15, 40, 20, 132, 214, 37, 84, 22},
			expStr:     "OQHSQFEE2YSVIFQ",
			expI16:     nil,
			expI16Err:  &WrongSizeError{errors.New(ErrorMsgSizeDivisible2)},
			expUI16:    nil,
			expUI16Err: &WrongSizeError{errors.New(ErrorMsgSizeDivisible2)},
			expI32:     nil,
			expI32Err:  &WrongSizeError{errors.New(ErrorMsgSizeDivisible4)},
			expUI32:    nil,
			expUI32Err: &WrongSizeError{errors.New(ErrorMsgSizeDivisible4)},
			expI64:     nil,
			expI64Err:  &WrongSizeError{errors.New(ErrorMsgSizeDivisible8)},
			expUI64:    nil,
			expUI64Err: &WrongSizeError{errors.New(ErrorMsgSizeDivisible8)},
		},
		{
			act:        []byte{45, 70, 132, 210, 205, 57, 124, 69, 85, 146},
			expStr:     "FVDIJUWNHF6EKVMS",
			expI16:     []int16{17965, -11644, 14797, 17788, -28075},
			expI16Err:  nil,
			expUI16:    []uint16{17965, 53892, 14797, 17788, 37461},
			expUI16Err: nil,
			expI32:     nil,
			expI32Err:  &WrongSizeError{errors.New(ErrorMsgSizeDivisible4)},
			expUI32:    nil,
			expUI32Err: &WrongSizeError{errors.New(ErrorMsgSizeDivisible4)},
			expI64:     nil,
			expI64Err:  &WrongSizeError{errors.New(ErrorMsgSizeDivisible8)},
			expUI64:    nil,
			expUI64Err: &WrongSizeError{errors.New(ErrorMsgSizeDivisible8)},
		},
		{
			act:        []byte{172, 255, 166, 0, 10, 143, 110, 12, 208, 93, 32, 243},
			expStr:     "VT72MAAKR5XAZUC5EDZQ",
			expI16:     []int16{-84, 166, -28918, 3182, 24016, -3296},
			expI16Err:  nil,
			expUI16:    []uint16{65452, 166, 36618, 3182, 24016, 62240},
			expUI16Err: nil,
			expI32:     []int32{10944428, 208572170, -215982640},
			expI32Err:  nil,
			expUI32:    []uint32{10944428, 208572170, 4078984656},
			expUI32Err: nil,
			expI64:     nil,
			expI64Err:  &WrongSizeError{errors.New(ErrorMsgSizeDivisible8)},
			expUI64:    nil,
			expUI64Err: &WrongSizeError{errors.New(ErrorMsgSizeDivisible8)},
		},
		{
			act:        []byte{151, 81, 2, 44, 31, 138, 170, 14, 90, 182, 24, 210, 176, 73, 176},
			expStr:     "S5IQELA7RKVA4WVWDDJLASNQ",
			expI16:     nil,
			expI16Err:  &WrongSizeError{errors.New(ErrorMsgSizeDivisible2)},
			expUI16:    nil,
			expUI16Err: &WrongSizeError{errors.New(ErrorMsgSizeDivisible2)},
			expI32:     nil,
			expI32Err:  &WrongSizeError{errors.New(ErrorMsgSizeDivisible4)},
			expUI32:    nil,
			expUI32Err: &WrongSizeError{errors.New(ErrorMsgSizeDivisible4)},
			expI64:     nil,
			expI64Err:  &WrongSizeError{errors.New(ErrorMsgSizeDivisible8)},
			expUI64:    nil,
			expUI64Err: &WrongSizeError{errors.New(ErrorMsgSizeDivisible8)},
		},
		{
			act:        []byte{170, 153, 153, 35, 130, 252, 27, 7, 162, 65, 106, 248, 89, 227, 105, 169},
			expStr:     "VKMZSI4C7QNQPISBNL4FTY3JVE",
			expI16:     []int16{-26198, 9113, -894, 1819, 16802, -1942, -7335, -22167},
			expI16Err:  nil,
			expUI16:    []uint16{39338, 9113, 64642, 1819, 16802, 63594, 58201, 43369},
			expUI16Err: nil,
			expI32:     []int32{597268906, 119274626, -127254110, -1452678311},
			expI32Err:  nil,
			expUI32:    []uint32{597268906, 119274626, 4167713186, 2842288985},
			expUI32Err: nil,
			expI64:     []int64{512280618509900202, -6239205833185803870},
			expI64Err:  nil,
			expUI64:    []uint64{512280618509900202, 12207538240523747746},
			expUI64Err: nil,
		},
		{
			act:        []byte{28, 37, 77, 21, 5, 65, 86, 50, 57, 214, 82, 83, 239, 38, 125, 77, 47, 132, 140, 7},
			expStr:     "DQSU2FIFIFLDEOOWKJJ66JT5JUXYJDAH",
			expI16:     []int16{9500, 5453, 16645, 12886, -10695, 21330, 9967, 19837, -31697, 1932},
			expI16Err:  nil,
			expUI16:    []uint16{9500, 5453, 16645, 12886, 54841, 21330, 9967, 19837, 33839, 1932},
			expUI16Err: nil,
			expI32:     []int32{357377308, 844513541, 1397937721, 1300047599, 126649391},
			expI32Err:  nil,
			expUI32:    []uint32{357377308, 844513541, 1397937721, 1300047599, 126649391},
			expUI32Err: nil,
			expI64:     nil,
			expI64Err:  &WrongSizeError{errors.New(ErrorMsgSizeDivisible8)},
			expUI64:    nil,
			expUI64Err: &WrongSizeError{errors.New(ErrorMsgSizeDivisible8)},
		},
	}
}
