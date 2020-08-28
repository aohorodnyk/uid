package intl

import "encoding/binary"

type Providing struct {
	randSize uint32
	rand     Randomizer
	enc      Encoder
}

func (p Providing) Generate() (Identifier, error) {
	data, err := p.rand.Generate(p.randSize)
	if err != nil {
		return nil, &RandError{err}
	}

	return p.Byte(data), nil
}

func (p Providing) Parse(id string) (Identifier, error) {
	data, err := p.enc.DecodeString(id)
	if err != nil {
		return nil, &EncError{err}
	}

	return p.Byte(data), nil
}

func (p Providing) Byte(id []byte) Identifier {
	return NewID(id, p.enc)
}

func (p Providing) Int16(id []int16) Identifier {
	uID := make([]uint16, len(id))
	for i, val := range id {
		uID[i] = uint16(val)
	}

	return p.Uint16(uID)
}

func (p Providing) Uint16(id []uint16) Identifier {
	dataSize := SizeUint16 * len(id)
	data := make([]byte, dataSize)
	for idx, val := range id {
		curIdx := idx * SizeUint16
		dstIdx := curIdx + SizeUint16
		binary.LittleEndian.PutUint16(data[curIdx:dstIdx], val)
	}

	return p.Byte(data)
}

func (p Providing) Int32(id []int32) Identifier {
	uID := make([]uint32, len(id))
	for i, val := range id {
		uID[i] = uint32(val)
	}

	return p.Uint32(uID)
}

func (p Providing) Uint32(id []uint32) Identifier {
	dataSize := SizeUint32 * len(id)
	data := make([]byte, dataSize)
	for idx, val := range id {
		curIdx := idx * SizeUint32
		dstIdx := curIdx + SizeUint32
		binary.LittleEndian.PutUint32(data[curIdx:dstIdx], val)
	}

	return p.Byte(data)
}

func (p Providing) Int64(id []int64) Identifier {
	uID := make([]uint64, len(id))
	for i, val := range id {
		uID[i] = uint64(val)
	}

	return p.Uint64(uID)
}

func (p Providing) Uint64(id []uint64) Identifier {
	dataSize := SizeUint64 * len(id)
	data := make([]byte, dataSize)
	for idx, val := range id {
		curIdx := idx * SizeUint64
		dstIdx := curIdx + SizeUint64
		binary.LittleEndian.PutUint64(data[curIdx:dstIdx], val)
	}

	return p.Byte(data)
}
