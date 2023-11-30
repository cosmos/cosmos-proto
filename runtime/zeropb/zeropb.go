package zeropb

import (
	"encoding/binary"
	"fmt"
	"math"
)

type Buffer struct {
	buf       []byte
	allocated uint16
}

type Allocation struct {
	Buf    []byte
	offset uint16
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
	n16 := uint16(n)
	if int(n16) != n {
		panic(fmt.Errorf("allocation %d too large", n))
	}
	a := Allocation{
		Buf:    b.buf[:n16],
		offset: b.allocated,
	}
	b.buf = b.buf[n16:]
	b.allocated += n16
	return a
}

func (b *Buffer) AllocRel(n int, dst Allocation, offset, len uint16) Allocation {
	a := b.Alloc(n)
	bo := binary.LittleEndian
	// Write relative offset and len.
	bo.PutUint16(dst.Buf[offset:], a.offset-dst.offset-offset)
	bo.PutUint16(dst.Buf[offset+2:], len)
	return a
}

func (b *Buffer) Allocated() uint16 {
	return b.allocated
}

func (a Allocation) Slice(offset uint16) Allocation {
	return Allocation{
		Buf:    a.Buf[offset:],
		offset: a.offset + offset,
	}
}

func ReadSlice(buf []byte, offset uint16) (off, len uint16) {
	bo := binary.LittleEndian
	off = offset + bo.Uint16(buf[offset:])
	len = bo.Uint16(buf[offset+2:])
	return
}
