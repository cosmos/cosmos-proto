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
