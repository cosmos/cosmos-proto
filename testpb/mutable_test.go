package testpb

import (
	"testing"

	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/types/dynamicpb"
)

func TestMutable(t *testing.T) {
	panics := func(f func()) (panics bool) {
		panics = true
		defer func() { recover() }()
		f()
		return false
	}

	fds := md_A.Fields()

	t.Run("panics", func(t *testing.T) {
		dyn := dynamicpb.NewMessage(md_A)
		m := &A{}

		for i := 0; i < fds.Len(); i++ {
			fd := fds.Get(i)
			// test panic cases
			if panics(func() {
				dyn.Mutable(fd)
			}) {
				// assert fields that panic for dynamicpb panic for our implementation too
				require.Panics(t, func() {
					m.ProtoReflect().Mutable(fd)
				})
			} else {
				// assert fields that do not panic for dynamipb do not panic for our implementation
				require.NotPanics(t, func() {
					m.ProtoReflect().Mutable(fd)
				})
			}
		}
	})

	/*
		t.Run("mutability", func(t *testing.T) {
			dyn := dynamicpb.NewMessage(md_A)
			m := &A{}
			for i := 0; i < fds.Len(); i++ {
				fd := fds.Get(i)
				// skip panic cases
				if panics(func() {
					dyn.Mutable(fd)
				}) {
					continue
				}

				v := m.ProtoReflect().Mutable(fd)
				dv := dyn.Mutable(fd)

				switch {
				case fd.IsMap():
				case fd.IsList():
				case fd.Kind() == protoreflect.MessageKind:
				}

			}
		})
	*/

}
