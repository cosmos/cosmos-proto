package testpb

import (
	"testing"
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
