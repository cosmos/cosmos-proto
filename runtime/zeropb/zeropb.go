package zeropb

import (
	"encoding/binary"
	"math"
)

type Buffer struct {
	buf       []byte
	allocated uint16
}

type Allocation struct {
	Buf    []byte
	Offset uint16
}

func NewBuffer(b []byte) *Buffer {
	if len(b) > math.MaxUint16 {
		b = b[:math.MaxUint16]
	}
	return &Buffer{
		buf: b,
	}
}

func (b *Buffer) Alloc(n int) Allocation {
	a := Allocation{
		Buf:    b.buf[:n],
		Offset: b.allocated,
	}
	// The slice above would have paniced if n would not fit a uint16.
	n16 := uint16(n)
	b.buf = b.buf[n16:]
	b.allocated += n16
	return a
}

func (b *Buffer) AllocRel(n int, dst Allocation, offset, len uint16) Allocation {
	a := b.Alloc(n)
	bo := binary.LittleEndian
	// Write relative offset and len.
	bo.PutUint16(dst.Buf[offset:], a.Offset-dst.Offset-offset)
	bo.PutUint16(dst.Buf[offset+2:], len)
	return a
}

func (b *Buffer) Allocated() uint16 {
	return b.allocated
}

func (a Allocation) Slice(offset uint16) Allocation {
	return Allocation{
		Buf:    a.Buf[offset:],
		Offset: a.Offset + offset,
	}
}

func ReadSlice(buf []byte, offset uint16) (off, len uint16) {
	bo := binary.LittleEndian
	off = ReadOffset(buf, offset)
	len = bo.Uint16(buf[offset+2:])
	return
}

func ReadOffset(buf []byte, offset uint16) (off uint16) {
	bo := binary.LittleEndian
	off = offset + bo.Uint16(buf[offset:])
	return
}
