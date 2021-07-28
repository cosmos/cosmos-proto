package testpb

import (
	"testing"

	"google.golang.org/protobuf/reflect/protoreflect"
)

func Benchmark_Get_FR(b *testing.B) {
	msg := &A{LIST: []*B{
		{
			X: "test",
		},
	}}

	fd := msg.ProtoReflect().Descriptor().Fields().ByName("LIST")

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = msg.ProtoReflect().Get(fd)
	}
}

func Benchmark_Get_SR(b *testing.B) {
	msg := &A{LIST: []*B{
		{
			X: "test",
		},
	}}

	fd := msg.ProtoReflect().Descriptor().Fields().ByName("LIST")

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = msg.slowProtoReflect().Get(fd)
	}
}

func Benchmark_WhichOneof_FR(b *testing.B) {
	msg := &A{ONEOF: &A_ONEOF_B{}}

	od := msg.ProtoReflect().Descriptor().Oneofs().ByName("ONEOF")

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = msg.ProtoReflect().WhichOneof(od)
	}
}

func Benchmark_WhichOneof_SR(b *testing.B) {
	msg := &A{ONEOF: &A_ONEOF_B{}}

	od := msg.ProtoReflect().Descriptor().Oneofs().ByName("ONEOF")

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = msg.slowProtoReflect().WhichOneof(od)
	}
}

func Benchmark_Has_FR(b *testing.B) {
	msg := &A{BYTES: nil}
	fd := (&A{}).ProtoReflect().Descriptor().Fields().ByName("BYTES")

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		msg.ProtoReflect().Has(fd)
	}
}

func Benchmark_Has_SR(b *testing.B) {
	msg := &A{BYTES: nil}
	fd := (&A{}).ProtoReflect().Descriptor().Fields().ByName("BYTES")

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = msg.slowProtoReflect().Has(fd)
	}
}

func Benchmark_Clear_FR(b *testing.B) {
	msg := &A{BYTES: []byte("test")}
	fd := (&A{}).ProtoReflect().Descriptor().Fields().ByName("BYTES")

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		msg.ProtoReflect().Clear(fd)
	}
}

func Benchmark_Clear_SR(b *testing.B) {
	msg := &A{BYTES: []byte("test")}
	fd := (&A{}).ProtoReflect().Descriptor().Fields().ByName("BYTES")

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		msg.slowProtoReflect().Clear(fd)
	}
}

func Benchmark_Set_FR(b *testing.B) {
	msg := &A{}
	fd := (&A{}).ProtoReflect().Descriptor().Fields().ByName("BYTES")
	v := protoreflect.ValueOfBytes([]byte("test"))
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		msg.ProtoReflect().Set(fd, v)
	}
}

func Benchmark_Set_SR(b *testing.B) {
	msg := &A{}
	fd := (&A{}).ProtoReflect().Descriptor().Fields().ByName("BYTES")
	v := protoreflect.ValueOfBytes([]byte("test"))
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		msg.slowProtoReflect().Set(fd, v)
	}
}

func Benchmark_NewField_FR(b *testing.B) {
	msg := &A{}
	fd := (&A{}).ProtoReflect().Descriptor().Fields().ByName("BYTES")
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = msg.ProtoReflect().NewField(fd)
	}
}

func Benchmark_NewField_SR(b *testing.B) {
	msg := &A{}
	fd := (&A{}).ProtoReflect().Descriptor().Fields().ByName("BYTES")
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = msg.slowProtoReflect().NewField(fd)
	}
}

func Benchmark_Mutable_FR(b *testing.B) {
	msg := &A{}
	fd := (&A{}).ProtoReflect().Descriptor().Fields().ByName("MESSAGE")
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = msg.ProtoReflect().Mutable(fd)
	}
}

func Benchmark_Mutable_SR(b *testing.B) {
	msg := &A{}
	fd := (&A{}).ProtoReflect().Descriptor().Fields().ByName("MESSAGE")
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = msg.slowProtoReflect().NewField(fd)
	}
}

func Benchmark_Range_FR(b *testing.B) {
	msg := &A{
		Enum:        Enumeration_Two,
		SomeBoolean: true,
		INT32:       2,
		SINT32:      3,
		UINT32:      4,
		INT64:       5,
		SING64:      6,
		UINT64:      7,
		SFIXED32:    8,
		FIXED32:     9,
		FLOAT:       10.1,
		SFIXED64:    11,
		FIXED64:     12,
		DOUBLE:      13,
		STRING:      "fourteen",
		BYTES:       []byte("fifteen"),
		MESSAGE:     &B{X: "something"},
		MAP:         map[string]*B{"a": &B{X: "aa"}},
		LIST:        []*B{{X: "list"}},
		ONEOF:       &A_ONEOF_B{ONEOF_B: &B{X: "ONEOF"}},
		LIST_ENUM:   []Enumeration{Enumeration_One},
	}

	for i := 0; i < b.N; i++ {
		msg.ProtoReflect().Range(func(_ protoreflect.FieldDescriptor, _ protoreflect.Value) bool {
			return true
		})
	}
}

func Benchmark_Range_SR(b *testing.B) {
	msg := &A{
		Enum:        Enumeration_Two,
		SomeBoolean: true,
		INT32:       2,
		SINT32:      3,
		UINT32:      4,
		INT64:       5,
		SING64:      6,
		UINT64:      7,
		SFIXED32:    8,
		FIXED32:     9,
		FLOAT:       10.1,
		SFIXED64:    11,
		FIXED64:     12,
		DOUBLE:      13,
		STRING:      "fourteen",
		BYTES:       []byte("fifteen"),
		MESSAGE:     &B{X: "something"},
		MAP:         map[string]*B{"a": &B{X: "aa"}},
		LIST:        []*B{{X: "list"}},
		ONEOF:       &A_ONEOF_B{ONEOF_B: &B{X: "ONEOF"}},
		LIST_ENUM:   []Enumeration{Enumeration_One},
	}

	for i := 0; i < b.N; i++ {
		msg.slowProtoReflect().Range(func(_ protoreflect.FieldDescriptor, _ protoreflect.Value) bool {
			return true
		})
	}
}
