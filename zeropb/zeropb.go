package zeropb

import (
	"encoding/binary"
	"fmt"
	"math"
)

type Serializable[T any] interface {
	WithBufferContext(*BufferContext) *T
	BufferContext() *BufferContext
	Size() uint32
}

type Root struct {
	buf []byte
}

type BufferContext struct {
	parent *Root
	offset uint32
	buf    []byte
}

const DefaultBufferCapacity = 1024

func (b *Root) Alloc(n uint32) (offset uint32, ctx *BufferContext) {
	if b.buf == nil {
		b.buf = make([]byte, 0, DefaultBufferCapacity)
	}

	offset = uint32(len(b.buf))
	for i := uint32(0); i < n; i++ {
		b.buf = append(b.buf, 0)
	}
	fmt.Printf("alloc %d bytes at offset %d\n", n, offset)
	ctx = &BufferContext{
		parent: b,
		offset: offset,
		buf:    b.buf[offset : offset+n],
	}
	return
}

func (b *Root) Resolve(offset uint32) *BufferContext {
	ctx := &BufferContext{
		parent: b,
		buf:    b.buf[offset:],
		offset: offset,
	}
	return ctx
}

func (c BufferContext) ReadBool(offset uint32) bool {
	return c.buf[offset] != 0
}

func (c BufferContext) SetBool(offset uint32, value bool) {
	if value {
		c.buf[offset] = 1
	} else {
		c.buf[offset] = 0
	}
}

func (c BufferContext) ReadUint8(offset uint32) uint8 {
	return c.buf[offset]
}

func (c BufferContext) SetUint8(offset uint32, value uint8) {
	c.buf[offset] = value
}

func (c BufferContext) ReadInt32(offset uint32) int32 {
	return int32(c.ReadUint32(offset))
}

func (c BufferContext) SetInt32(offset uint32, value int32) {
	c.SetUint32(offset, uint32(value))
}

func (c BufferContext) ReadUint16(offset uint32) uint16 {
	return binary.LittleEndian.Uint16(c.buf[offset : offset+2])
}

func (c BufferContext) SetUint16(offset uint32, value uint16) {
	binary.LittleEndian.PutUint16(c.buf[offset:offset+2], value)
}

func (c BufferContext) ReadUint32(offset uint32) uint32 {
	return binary.LittleEndian.Uint32(c.buf[offset : offset+4])
}

func (c BufferContext) SetUint32(offset uint32, value uint32) {
	fmt.Printf("set uint32 at offset %d to %d\n", offset+c.offset, value)
	binary.LittleEndian.PutUint32(c.buf[offset:offset+4], value)
}

func (c BufferContext) ReadInt64(offset uint32) int64 {
	return int64(c.ReadUint64(offset))
}

func (c BufferContext) SetInt64(offset uint32, value int64) {
	c.SetUint64(offset, uint64(value))
}

func (c BufferContext) ReadUint64(offset uint32) uint64 {
	return binary.LittleEndian.Uint64(c.buf[offset : offset+8])
}

func (c BufferContext) SetUint64(offset uint32, value uint64) {
	fmt.Printf("set uint64 at offset %d to %d\n", offset+c.offset, value)
	binary.LittleEndian.PutUint64(c.buf[offset:offset+8], value)
}

func (c BufferContext) ReadFloat32(offset uint32) float32 {
	return math.Float32frombits(c.ReadUint32(offset))
}

func (c BufferContext) SetFloat32(offset uint32, value float32) {
	c.SetUint32(offset, math.Float32bits(value))
}

func (c BufferContext) ReadFloat64(offset uint32) float64 {
	return math.Float64frombits(c.ReadUint64(offset))
}

func (c BufferContext) SetFloat64(offset uint32, value float64) {
	c.SetUint64(offset, math.Float64bits(value))
}

func (c BufferContext) ResolvePointer(offset uint32) *BufferContext {
	ptrOffset := c.ReadUint16(offset)
	if ptrOffset == 0 {
		return nil
	}
	return c.parent.Resolve(uint32(ptrOffset))
}

func (c BufferContext) AllocPointer(offset uint32, size uint32) *BufferContext {
	ptrOffset, ctx := c.parent.Alloc(size)
	c.SetUint16(offset, uint16(ptrOffset))
	return ctx
}

func (c BufferContext) ClearPointer(offset uint32) {
	c.SetUint16(offset, 0)
}

func (c BufferContext) AsString() string {
	return string(c.AsBytes())
}

func (c BufferContext) AsBytes() []byte {
	n := c.ReadUint16(0)
	return c.buf[4 : 4+n]
}

func (c BufferContext) ReadString(offset uint32) string {
	return c.ResolvePointer(offset).AsString()
}

func (c BufferContext) SetString(offset uint32, x string) {
	c.SetBytes(offset, []byte(x))
}

func (c BufferContext) ReadBytes(offset uint32) []byte {
	return c.ResolvePointer(offset).AsBytes()
}

func (c BufferContext) SetBytes(offset uint32, x []byte) {
	n := len(x)
	if n > math.MaxUint16 {
		panic("byte array too long")
	}
	ptrOffset, ctx := c.parent.Alloc(uint32(n) + 4)
	c.SetUint16(offset, uint16(ptrOffset))
	binary.LittleEndian.PutUint16(ctx.buf[0:4], uint16(uint32(n)))
	copy(ctx.buf[4:4+n], x)
}

func ReadArray[T any, PT interface {
	Serializable[T]
	*T
}](ctx *BufferContext, offset uint32) Array[*T] {
	var elem PT
	size := elem.Size()
	arrayCtx := ctx.ResolvePointer(offset)
	if arrayCtx == nil {
		return Array[*T]{array: nil}
	}

	n := arrayCtx.ReadUint32(0)
	array := make([]*T, n)
	elemOffset := uint32(4) + arrayCtx.offset
	for i := uint32(0); i < n; i++ {
		elem := new(T)
		pt := PT(elem)
		parent := ctx.parent
		pt.WithBufferContext(&BufferContext{
			parent: parent,
			offset: elemOffset,
			buf:    parent.buf[elemOffset : elemOffset+size],
		})
		array[i] = elem
		elemOffset += size
	}

	return Array[*T]{array: array}
}

func InitArray[T any, PT interface {
	Serializable[T]
	*T
}](ctx *BufferContext, offset uint32, n int) Array[*T] {
	var elem PT
	size := elem.Size()
	arrayCtx := ctx.AllocPointer(offset, 4+uint32(n)*size)
	arrayCtx.SetUint32(0, uint32(n))
	array := make([]*T, n)
	elemOffset := uint32(4) + arrayCtx.offset
	for i := 0; i < n; i++ {
		elem := new(T)
		pt := PT(elem)
		parent := ctx.parent
		pt.WithBufferContext(&BufferContext{
			parent: parent,
			offset: elemOffset,
			buf:    parent.buf[elemOffset : elemOffset+size],
		})
		array[i] = elem
		elemOffset += size
	}
	return Array[*T]{array: array}
}

// Array is a fixed length array. Its elements, however, should be assumed to be mutable.
type Array[T any] struct {
	array []T
}

func (a Array[T]) Len() int {
	return len(a.array)
}

func (a Array[T]) Get(i int) T {
	return a.array[i]
}
