package runtime

import (
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/reflect/protoregistry"
)

// NullRegistry is an implementation of the interfaces used to register generated
// types and file descriptors that does nothing. This is used to generated
// protobuf files in internal packages that are not registered with the global registry.
type NullRegistry struct{}

func (NullRegistry) RegisterMessage(protoreflect.MessageType) error     { return nil }
func (NullRegistry) RegisterEnum(protoreflect.EnumType) error           { return nil }
func (NullRegistry) RegisterExtension(protoreflect.ExtensionType) error { return nil }
func (NullRegistry) RegisterFile(protoreflect.FileDescriptor) error     { return nil }

func (NullRegistry) FindFileByPath(path string) (protoreflect.FileDescriptor, error) {
	return protoregistry.GlobalFiles.FindFileByPath(path)
}
func (NullRegistry) FindDescriptorByName(name protoreflect.FullName) (protoreflect.Descriptor, error) {
	return protoregistry.GlobalFiles.FindDescriptorByName(name)
}
