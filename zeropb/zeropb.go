package zeropb

import (
	"encoding/binary"
	"fmt"

	"google.golang.org/protobuf/reflect/protoreflect"
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

func (c BufferContext) ReadUint32(offset uint32) uint32 {
	return binary.LittleEndian.Uint32(c.buf[offset : offset+4])
}

func (c BufferContext) SetUint32(offset uint32, value uint32) {
	fmt.Printf("set uint32 at offset %d to %d\n", offset+c.offset, value)
	binary.LittleEndian.PutUint32(c.buf[offset:offset+4], value)
}

func (c BufferContext) ReadUint64(offset uint32) uint64 {
	return binary.LittleEndian.Uint64(c.buf[offset : offset+8])
}

func (c BufferContext) SetUint64(offset uint32, value uint64) {
	fmt.Printf("set uint64 at offset %d to %d\n", offset+c.offset, value)
	binary.LittleEndian.PutUint64(c.buf[offset:offset+8], value)
}

func (c BufferContext) ResolvePointer(offset uint32) *BufferContext {
	ptrOffset := c.ReadUint32(offset)
	if ptrOffset == 0 {
		return nil
	}
	return c.parent.Resolve(ptrOffset)
}

func (c BufferContext) AllocPointer(offset uint32, size uint32) *BufferContext {
	ptrOffset, ctx := c.parent.Alloc(size)
	c.SetUint32(offset, ptrOffset)
	return ctx
}

func (c BufferContext) ReadString() string {
	n := c.ReadUint32(0)
	return string(c.buf[4 : 4+n])
}

func (c BufferContext) SetString(offset uint32, x string) {
	n := len(x)
	ptrOffset, ctx := c.parent.Alloc(uint32(n) + 4)
	c.SetUint32(offset, ptrOffset)
	binary.LittleEndian.PutUint32(ctx.buf[0:4], uint32(n))
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

type FieldDescriptors struct {
	protoreflect.FieldDescriptors
	fds        []protoreflect.FieldDescriptor
	byName     map[protoreflect.Name]protoreflect.FieldDescriptor
	byJSONName map[string]protoreflect.FieldDescriptor
	byTextName map[string]protoreflect.FieldDescriptor
	byNumber   map[protoreflect.FieldNumber]protoreflect.FieldDescriptor
}

func NewFieldDescriptors(parent protoreflect.FieldDescriptors, overrides []protoreflect.FieldDescriptor) *FieldDescriptors {
	fds := &FieldDescriptors{
		FieldDescriptors: parent,
		fds:              overrides,
		byName:           map[protoreflect.Name]protoreflect.FieldDescriptor{},
		byJSONName:       map[string]protoreflect.FieldDescriptor{},
		byTextName:       map[string]protoreflect.FieldDescriptor{},
		byNumber:         map[protoreflect.FieldNumber]protoreflect.FieldDescriptor{},
	}

	for _, fd := range overrides {
		fds.byName[fd.Name()] = fd
		fds.byJSONName[fd.JSONName()] = fd
		fds.byTextName[fd.TextName()] = fd
		fds.byNumber[fd.Number()] = fd
	}

	return fds
}

func (f FieldDescriptors) Len() int {
	return len(f.fds)
}

func (f FieldDescriptors) Get(i int) protoreflect.FieldDescriptor {
	return f.fds[i]
}

func (f FieldDescriptors) ByName(s protoreflect.Name) protoreflect.FieldDescriptor {
	return f.byName[s]
}

func (f FieldDescriptors) ByJSONName(s string) protoreflect.FieldDescriptor {
	return f.byJSONName[s]
}

func (f FieldDescriptors) ByTextName(s string) protoreflect.FieldDescriptor {
	return f.byTextName[s]
}

func (f FieldDescriptors) ByNumber(n protoreflect.FieldNumber) protoreflect.FieldDescriptor {
	return f.byNumber[n]
}

var _ protoreflect.FieldDescriptors = FieldDescriptors{}
