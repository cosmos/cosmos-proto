package testpb

import (
	"testing"
)

func Benchmark_GetFastReflection(b *testing.B) {
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

func Benchmark_GetSlowReflection(b *testing.B) {
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
