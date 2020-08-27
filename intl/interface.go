package intl

type Provider interface {
	Generate() (Identifier, error)
	Parse(id string) (Identifier, error)
	Byte(id []byte) Identifier
	Int16(id []int16) Identifier
	Uint16(id []uint16) Identifier
	Int32(id []int32) Identifier
	Uint32(id []uint32) Identifier
	Int64(id []int64) Identifier
	Uint64(id []uint64) Identifier
}

type Identifier interface {
	String() string
	Byte() []byte
	Int16() ([]int16, error)
	Uint16() ([]uint16, error)
	Int32() ([]int32, error)
	Uint32() ([]uint32, error)
	Int64() ([]int64, error)
	Uint64() ([]uint64, error)
}

type Randomizer interface {
	Generate(size uint32) ([]byte, error)
}

type Encoder interface {
	EncodeToString(src []byte) string
	DecodeString(s string) ([]byte, error)
}
