package testpb

import (
	binary "encoding/binary"
	fmt "fmt"
	runtime "github.com/cosmos/cosmos-proto/runtime"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoiface "google.golang.org/protobuf/runtime/protoiface"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	descriptorpb "google.golang.org/protobuf/types/descriptorpb"
	io "io"
	math "math"
	reflect "reflect"
	sort "sort"
	sync "sync"
)

var _ protoreflect.Map = (*_A_18_map)(nil)

type _A_18_map struct {
	m *map[string]*B
}

func (x *_A_18_map) Len() int {
	if x.m == nil {
		return 0
	}
	return len(*x.m)
}

func (x *_A_18_map) Range(f func(protoreflect.MapKey, protoreflect.Value) bool) {
	if x.m == nil {
		return
	}
	for k, v := range *x.m {
		mapKey := (protoreflect.MapKey)(protoreflect.ValueOfString(k))
		mapValue := protoreflect.ValueOfMessage(v.ProtoReflect())
		if !f(mapKey, mapValue) {
			break
		}
	}
}

func (x *_A_18_map) Has(key protoreflect.MapKey) bool {
	if x.m == nil {
		return false
	}
	keyUnwrapped := key.String()
	concreteValue := keyUnwrapped
	_, ok := (*x.m)[concreteValue]
	return ok
}

func (x *_A_18_map) Clear(key protoreflect.MapKey) {
	if x.m == nil {
		return
	}
	keyUnwrapped := key.String()
	concreteKey := keyUnwrapped
	delete(*x.m, concreteKey)
}

func (x *_A_18_map) Get(key protoreflect.MapKey) protoreflect.Value {
	if x.m == nil {
		return protoreflect.Value{}
	}
	keyUnwrapped := key.String()
	concreteKey := keyUnwrapped
	v, ok := (*x.m)[concreteKey]
	if !ok {
		return protoreflect.Value{}
	}
	return protoreflect.ValueOfMessage(v.ProtoReflect())
}

func (x *_A_18_map) Set(key protoreflect.MapKey, value protoreflect.Value) {
	if !key.IsValid() || !value.IsValid() {
		panic("invalid key or value provided")
	}
	keyUnwrapped := key.String()
	concreteKey := keyUnwrapped
	valueUnwrapped := value.Message()
	concreteValue := valueUnwrapped.Interface().(*B)
	(*x.m)[concreteKey] = concreteValue
}

func (x *_A_18_map) Mutable(key protoreflect.MapKey) protoreflect.Value {
	keyUnwrapped := key.String()
	concreteKey := keyUnwrapped
	v, ok := (*x.m)[concreteKey]
	if ok {
		return protoreflect.ValueOfMessage(v.ProtoReflect())
	}
	newValue := new(B)
	(*x.m)[concreteKey] = newValue
	return protoreflect.ValueOfMessage(newValue.ProtoReflect())
}

func (x *_A_18_map) NewValue() protoreflect.Value {
	v := new(B)
	return protoreflect.ValueOfMessage(v.ProtoReflect())
}

func (x *_A_18_map) IsValid() bool {
	return x.m != nil
}

var _ protoreflect.List = (*_A_19_list)(nil)

type _A_19_list struct {
	list *[]*B
}

func (x *_A_19_list) Len() int {
	if x.list == nil {
		return 0
	}
	return len(*x.list)
}

func (x *_A_19_list) Get(i int) protoreflect.Value {
	return protoreflect.ValueOfMessage((*x.list)[i].ProtoReflect())
}

func (x *_A_19_list) Set(i int, value protoreflect.Value) {
	valueUnwrapped := value.Message()
	concreteValue := valueUnwrapped.Interface().(*B)
	(*x.list)[i] = concreteValue
}

func (x *_A_19_list) Append(value protoreflect.Value) {
	valueUnwrapped := value.Message()
	concreteValue := valueUnwrapped.Interface().(*B)
	*x.list = append(*x.list, concreteValue)
}

func (x *_A_19_list) AppendMutable() protoreflect.Value {
	v := new(B)
	*x.list = append(*x.list, v)
	return protoreflect.ValueOfMessage(v.ProtoReflect())
}

func (x *_A_19_list) Truncate(n int) {
	for i := n; i < len(*x.list); i++ {
		(*x.list)[i] = nil
	}
	*x.list = (*x.list)[:n]
}

func (x *_A_19_list) NewElement() protoreflect.Value {
	v := new(B)
	return protoreflect.ValueOfMessage(v.ProtoReflect())
}

func (x *_A_19_list) IsValid() bool {
	return x.list != nil
}

var _ protoreflect.List = (*_A_22_list)(nil)

type _A_22_list struct {
	list *[]Enumeration
}

func (x *_A_22_list) Len() int {
	if x.list == nil {
		return 0
	}
	return len(*x.list)
}

func (x *_A_22_list) Get(i int) protoreflect.Value {
	return protoreflect.ValueOfEnum((protoreflect.EnumNumber)((*x.list)[i]))
}

func (x *_A_22_list) Set(i int, value protoreflect.Value) {
	valueUnwrapped := value.Enum()
	concreteValue := (Enumeration)(valueUnwrapped)
	(*x.list)[i] = concreteValue
}

func (x *_A_22_list) Append(value protoreflect.Value) {
	valueUnwrapped := value.Enum()
	concreteValue := (Enumeration)(valueUnwrapped)
	*x.list = append(*x.list, concreteValue)
}

func (x *_A_22_list) AppendMutable() protoreflect.Value {
	panic(fmt.Errorf("AppendMutable can not be called on message A at list field LIST_ENUM as it is not of Message kind"))
}

func (x *_A_22_list) Truncate(n int) {
	*x.list = (*x.list)[:n]
}

func (x *_A_22_list) NewElement() protoreflect.Value {
	v := 0
	return protoreflect.ValueOfEnum((protoreflect.EnumNumber)(v))
}

func (x *_A_22_list) IsValid() bool {
	return x.list != nil
}

var (
	md_A              protoreflect.MessageDescriptor
	fd_A_enum         protoreflect.FieldDescriptor
	fd_A_some_boolean protoreflect.FieldDescriptor
	fd_A_INT32        protoreflect.FieldDescriptor
	fd_A_SINT32       protoreflect.FieldDescriptor
	fd_A_UINT32       protoreflect.FieldDescriptor
	fd_A_INT64        protoreflect.FieldDescriptor
	fd_A_SING64       protoreflect.FieldDescriptor
	fd_A_UINT64       protoreflect.FieldDescriptor
	fd_A_SFIXED32     protoreflect.FieldDescriptor
	fd_A_FIXED32      protoreflect.FieldDescriptor
	fd_A_FLOAT        protoreflect.FieldDescriptor
	fd_A_SFIXED64     protoreflect.FieldDescriptor
	fd_A_FIXED64      protoreflect.FieldDescriptor
	fd_A_DOUBLE       protoreflect.FieldDescriptor
	fd_A_STRING       protoreflect.FieldDescriptor
	fd_A_BYTES        protoreflect.FieldDescriptor
	fd_A_MESSAGE      protoreflect.FieldDescriptor
	fd_A_MAP          protoreflect.FieldDescriptor
	fd_A_LIST         protoreflect.FieldDescriptor
	fd_A_ONEOF_B      protoreflect.FieldDescriptor
	fd_A_ONEOF_STRING protoreflect.FieldDescriptor
	fd_A_LIST_ENUM    protoreflect.FieldDescriptor
	fd_A_imported     protoreflect.FieldDescriptor
)

func init() {
	file_testpb_1_proto_init()
	md_A = File_testpb_1_proto.Messages().ByName("A")
	fd_A_enum = md_A.Fields().ByName("enum")
	fd_A_some_boolean = md_A.Fields().ByName("some_boolean")
	fd_A_INT32 = md_A.Fields().ByName("INT32")
	fd_A_SINT32 = md_A.Fields().ByName("SINT32")
	fd_A_UINT32 = md_A.Fields().ByName("UINT32")
	fd_A_INT64 = md_A.Fields().ByName("INT64")
	fd_A_SING64 = md_A.Fields().ByName("SING64")
	fd_A_UINT64 = md_A.Fields().ByName("UINT64")
	fd_A_SFIXED32 = md_A.Fields().ByName("SFIXED32")
	fd_A_FIXED32 = md_A.Fields().ByName("FIXED32")
	fd_A_FLOAT = md_A.Fields().ByName("FLOAT")
	fd_A_SFIXED64 = md_A.Fields().ByName("SFIXED64")
	fd_A_FIXED64 = md_A.Fields().ByName("FIXED64")
	fd_A_DOUBLE = md_A.Fields().ByName("DOUBLE")
	fd_A_STRING = md_A.Fields().ByName("STRING")
	fd_A_BYTES = md_A.Fields().ByName("BYTES")
	fd_A_MESSAGE = md_A.Fields().ByName("MESSAGE")
	fd_A_MAP = md_A.Fields().ByName("MAP")
	fd_A_LIST = md_A.Fields().ByName("LIST")
	fd_A_ONEOF_B = md_A.Fields().ByName("ONEOF_B")
	fd_A_ONEOF_STRING = md_A.Fields().ByName("ONEOF_STRING")
	fd_A_LIST_ENUM = md_A.Fields().ByName("LIST_ENUM")
	fd_A_imported = md_A.Fields().ByName("imported")
}

var _ protoreflect.Message = (*fastReflection_A)(nil)

type fastReflection_A A

func (x *A) ProtoReflect() protoreflect.Message {
	return (*fastReflection_A)(x)
}

func (x *A) slowProtoReflect() protoreflect.Message {
	mi := &file_testpb_1_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_A_messageType fastReflection_A_messageType
var _ protoreflect.MessageType = fastReflection_A_messageType{}

type fastReflection_A_messageType struct{}

func (x fastReflection_A_messageType) Zero() protoreflect.Message {
	return (*fastReflection_A)(nil)
}
func (x fastReflection_A_messageType) New() protoreflect.Message {
	return new(fastReflection_A)
}
func (x fastReflection_A_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_A
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_A) Descriptor() protoreflect.MessageDescriptor {
	return md_A
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_A) Type() protoreflect.MessageType {
	return _fastReflection_A_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_A) New() protoreflect.Message {
	return new(fastReflection_A)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_A) Interface() protoreflect.ProtoMessage {
	return (*A)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_A) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
	if x.Enum != 0 {
		value := protoreflect.ValueOfEnum((protoreflect.EnumNumber)(x.Enum))
		if !f(fd_A_enum, value) {
			return
		}
	}
	if x.SomeBoolean != false {
		value := protoreflect.ValueOfBool(x.SomeBoolean)
		if !f(fd_A_some_boolean, value) {
			return
		}
	}
	if x.INT32 != int32(0) {
		value := protoreflect.ValueOfInt32(x.INT32)
		if !f(fd_A_INT32, value) {
			return
		}
	}
	if x.SINT32 != int32(0) {
		value := protoreflect.ValueOfInt32(x.SINT32)
		if !f(fd_A_SINT32, value) {
			return
		}
	}
	if x.UINT32 != uint32(0) {
		value := protoreflect.ValueOfUint32(x.UINT32)
		if !f(fd_A_UINT32, value) {
			return
		}
	}
	if x.INT64 != int64(0) {
		value := protoreflect.ValueOfInt64(x.INT64)
		if !f(fd_A_INT64, value) {
			return
		}
	}
	if x.SING64 != int64(0) {
		value := protoreflect.ValueOfInt64(x.SING64)
		if !f(fd_A_SING64, value) {
			return
		}
	}
	if x.UINT64 != uint64(0) {
		value := protoreflect.ValueOfUint64(x.UINT64)
		if !f(fd_A_UINT64, value) {
			return
		}
	}
	if x.SFIXED32 != int32(0) {
		value := protoreflect.ValueOfInt32(x.SFIXED32)
		if !f(fd_A_SFIXED32, value) {
			return
		}
	}
	if x.FIXED32 != uint32(0) {
		value := protoreflect.ValueOfUint32(x.FIXED32)
		if !f(fd_A_FIXED32, value) {
			return
		}
	}
	if x.FLOAT != float32(0) || math.Signbit(float64(x.FLOAT)) {
		value := protoreflect.ValueOfFloat32(x.FLOAT)
		if !f(fd_A_FLOAT, value) {
			return
		}
	}
	if x.SFIXED64 != int64(0) {
		value := protoreflect.ValueOfInt64(x.SFIXED64)
		if !f(fd_A_SFIXED64, value) {
			return
		}
	}
	if x.FIXED64 != uint64(0) {
		value := protoreflect.ValueOfUint64(x.FIXED64)
		if !f(fd_A_FIXED64, value) {
			return
		}
	}
	if x.DOUBLE != float64(0) || math.Signbit(x.DOUBLE) {
		value := protoreflect.ValueOfFloat64(x.DOUBLE)
		if !f(fd_A_DOUBLE, value) {
			return
		}
	}
	if x.STRING != "" {
		value := protoreflect.ValueOfString(x.STRING)
		if !f(fd_A_STRING, value) {
			return
		}
	}
	if len(x.BYTES) != 0 {
		value := protoreflect.ValueOfBytes(x.BYTES)
		if !f(fd_A_BYTES, value) {
			return
		}
	}
	if x.MESSAGE != nil {
		value := protoreflect.ValueOfMessage(x.MESSAGE.ProtoReflect())
		if !f(fd_A_MESSAGE, value) {
			return
		}
	}
	if len(x.MAP) != 0 {
		value := protoreflect.ValueOfMap(&_A_18_map{m: &x.MAP})
		if !f(fd_A_MAP, value) {
			return
		}
	}
	if len(x.LIST) != 0 {
		value := protoreflect.ValueOfList(&_A_19_list{list: &x.LIST})
		if !f(fd_A_LIST, value) {
			return
		}
	}
	if x.ONEOF != nil {
		switch o := x.ONEOF.(type) {
		case *A_ONEOF_B:
			v := o.ONEOF_B
			value := protoreflect.ValueOfMessage(v.ProtoReflect())
			if !f(fd_A_ONEOF_B, value) {
				return
			}
		case *A_ONEOF_STRING:
			v := o.ONEOF_STRING
			value := protoreflect.ValueOfString(v)
			if !f(fd_A_ONEOF_STRING, value) {
				return
			}
		}
	}
	if len(x.LIST_ENUM) != 0 {
		value := protoreflect.ValueOfList(&_A_22_list{list: &x.LIST_ENUM})
		if !f(fd_A_LIST_ENUM, value) {
			return
		}
	}
	if x.Imported != nil {
		value := protoreflect.ValueOfMessage(x.Imported.ProtoReflect())
		if !f(fd_A_imported, value) {
			return
		}
	}
}

// Has reports whether a field is populated.
//
// Some fields have the property of nullability where it is possible to
// distinguish between the default value of a field and whether the field
// was explicitly populated with the default value. Singular message fields,
// member fields of a oneof, and proto2 scalar fields are nullable. Such
// fields are populated only if explicitly set.
//
// In other cases (aside from the nullable cases above),
// a proto3 scalar field is populated if it contains a non-zero value, and
// a repeated field is populated if it is non-empty.
func (x *fastReflection_A) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	case "A.enum":
		return x.Enum != 0
	case "A.some_boolean":
		return x.SomeBoolean != false
	case "A.INT32":
		return x.INT32 != int32(0)
	case "A.SINT32":
		return x.SINT32 != int32(0)
	case "A.UINT32":
		return x.UINT32 != uint32(0)
	case "A.INT64":
		return x.INT64 != int64(0)
	case "A.SING64":
		return x.SING64 != int64(0)
	case "A.UINT64":
		return x.UINT64 != uint64(0)
	case "A.SFIXED32":
		return x.SFIXED32 != int32(0)
	case "A.FIXED32":
		return x.FIXED32 != uint32(0)
	case "A.FLOAT":
		return x.FLOAT != float32(0) || math.Signbit(float64(x.FLOAT))
	case "A.SFIXED64":
		return x.SFIXED64 != int64(0)
	case "A.FIXED64":
		return x.FIXED64 != uint64(0)
	case "A.DOUBLE":
		return x.DOUBLE != float64(0) || math.Signbit(x.DOUBLE)
	case "A.STRING":
		return x.STRING != ""
	case "A.BYTES":
		return len(x.BYTES) != 0
	case "A.MESSAGE":
		return x.MESSAGE != nil
	case "A.MAP":
		return len(x.MAP) != 0
	case "A.LIST":
		return len(x.LIST) != 0
	case "A.ONEOF_B":
		if x.ONEOF == nil {
			return false
		} else if _, ok := x.ONEOF.(*A_ONEOF_B); ok {
			return true
		} else {
			return false
		}
	case "A.ONEOF_STRING":
		if x.ONEOF == nil {
			return false
		} else if _, ok := x.ONEOF.(*A_ONEOF_STRING); ok {
			return true
		} else {
			return false
		}
	case "A.LIST_ENUM":
		return len(x.LIST_ENUM) != 0
	case "A.imported":
		return x.Imported != nil
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: A"))
		}
		panic(fmt.Errorf("message A does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_A) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	case "A.enum":
		x.Enum = 0
	case "A.some_boolean":
		x.SomeBoolean = false
	case "A.INT32":
		x.INT32 = int32(0)
	case "A.SINT32":
		x.SINT32 = int32(0)
	case "A.UINT32":
		x.UINT32 = uint32(0)
	case "A.INT64":
		x.INT64 = int64(0)
	case "A.SING64":
		x.SING64 = int64(0)
	case "A.UINT64":
		x.UINT64 = uint64(0)
	case "A.SFIXED32":
		x.SFIXED32 = int32(0)
	case "A.FIXED32":
		x.FIXED32 = uint32(0)
	case "A.FLOAT":
		x.FLOAT = float32(0)
	case "A.SFIXED64":
		x.SFIXED64 = int64(0)
	case "A.FIXED64":
		x.FIXED64 = uint64(0)
	case "A.DOUBLE":
		x.DOUBLE = float64(0)
	case "A.STRING":
		x.STRING = ""
	case "A.BYTES":
		x.BYTES = nil
	case "A.MESSAGE":
		x.MESSAGE = nil
	case "A.MAP":
		x.MAP = nil
	case "A.LIST":
		x.LIST = nil
	case "A.ONEOF_B":
		x.ONEOF = nil
	case "A.ONEOF_STRING":
		x.ONEOF = nil
	case "A.LIST_ENUM":
		x.LIST_ENUM = nil
	case "A.imported":
		x.Imported = nil
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: A"))
		}
		panic(fmt.Errorf("message A does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_A) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	case "A.enum":
		value := x.Enum
		return protoreflect.ValueOfEnum((protoreflect.EnumNumber)(value))
	case "A.some_boolean":
		value := x.SomeBoolean
		return protoreflect.ValueOfBool(value)
	case "A.INT32":
		value := x.INT32
		return protoreflect.ValueOfInt32(value)
	case "A.SINT32":
		value := x.SINT32
		return protoreflect.ValueOfInt32(value)
	case "A.UINT32":
		value := x.UINT32
		return protoreflect.ValueOfUint32(value)
	case "A.INT64":
		value := x.INT64
		return protoreflect.ValueOfInt64(value)
	case "A.SING64":
		value := x.SING64
		return protoreflect.ValueOfInt64(value)
	case "A.UINT64":
		value := x.UINT64
		return protoreflect.ValueOfUint64(value)
	case "A.SFIXED32":
		value := x.SFIXED32
		return protoreflect.ValueOfInt32(value)
	case "A.FIXED32":
		value := x.FIXED32
		return protoreflect.ValueOfUint32(value)
	case "A.FLOAT":
		value := x.FLOAT
		return protoreflect.ValueOfFloat32(value)
	case "A.SFIXED64":
		value := x.SFIXED64
		return protoreflect.ValueOfInt64(value)
	case "A.FIXED64":
		value := x.FIXED64
		return protoreflect.ValueOfUint64(value)
	case "A.DOUBLE":
		value := x.DOUBLE
		return protoreflect.ValueOfFloat64(value)
	case "A.STRING":
		value := x.STRING
		return protoreflect.ValueOfString(value)
	case "A.BYTES":
		value := x.BYTES
		return protoreflect.ValueOfBytes(value)
	case "A.MESSAGE":
		value := x.MESSAGE
		return protoreflect.ValueOfMessage(value.ProtoReflect())
	case "A.MAP":
		if len(x.MAP) == 0 {
			return protoreflect.ValueOfMap(&_A_18_map{})
		}
		mapValue := &_A_18_map{m: &x.MAP}
		return protoreflect.ValueOfMap(mapValue)
	case "A.LIST":
		if len(x.LIST) == 0 {
			return protoreflect.ValueOfList(&_A_19_list{})
		}
		listValue := &_A_19_list{list: &x.LIST}
		return protoreflect.ValueOfList(listValue)
	case "A.ONEOF_B":
		if x.ONEOF == nil {
			return protoreflect.ValueOfMessage((*B)(nil).ProtoReflect())
		} else if v, ok := x.ONEOF.(*A_ONEOF_B); ok {
			return protoreflect.ValueOfMessage(v.ONEOF_B.ProtoReflect())
		} else {
			return protoreflect.ValueOfMessage((*B)(nil).ProtoReflect())
		}
	case "A.ONEOF_STRING":
		if x.ONEOF == nil {
			return protoreflect.ValueOfString("")
		} else if v, ok := x.ONEOF.(*A_ONEOF_STRING); ok {
			return protoreflect.ValueOfString(v.ONEOF_STRING)
		} else {
			return protoreflect.ValueOfString("")
		}
	case "A.LIST_ENUM":
		if len(x.LIST_ENUM) == 0 {
			return protoreflect.ValueOfList(&_A_22_list{})
		}
		listValue := &_A_22_list{list: &x.LIST_ENUM}
		return protoreflect.ValueOfList(listValue)
	case "A.imported":
		value := x.Imported
		return protoreflect.ValueOfMessage(value.ProtoReflect())
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: A"))
		}
		panic(fmt.Errorf("message A does not contain field %s", descriptor.FullName()))
	}
}

// Set stores the value for a field.
//
// For a field belonging to a oneof, it implicitly clears any other field
// that may be currently set within the same oneof.
// For extension fields, it implicitly stores the provided ExtensionType.
// When setting a composite type, it is unspecified whether the stored value
// aliases the source's memory in any way. If the composite value is an
// empty, read-only value, then it panics.
//
// Set is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_A) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	case "A.enum":
		x.Enum = (Enumeration)(value.Enum())
	case "A.some_boolean":
		x.SomeBoolean = value.Bool()
	case "A.INT32":
		x.INT32 = int32(value.Int())
	case "A.SINT32":
		x.SINT32 = int32(value.Int())
	case "A.UINT32":
		x.UINT32 = uint32(value.Uint())
	case "A.INT64":
		x.INT64 = value.Int()
	case "A.SING64":
		x.SING64 = value.Int()
	case "A.UINT64":
		x.UINT64 = value.Uint()
	case "A.SFIXED32":
		x.SFIXED32 = int32(value.Int())
	case "A.FIXED32":
		x.FIXED32 = uint32(value.Uint())
	case "A.FLOAT":
		x.FLOAT = float32(value.Float())
	case "A.SFIXED64":
		x.SFIXED64 = value.Int()
	case "A.FIXED64":
		x.FIXED64 = value.Uint()
	case "A.DOUBLE":
		x.DOUBLE = value.Float()
	case "A.STRING":
		x.STRING = value.String()
	case "A.BYTES":
		x.BYTES = value.Bytes()
	case "A.MESSAGE":
		x.MESSAGE = value.Message().Interface().(*B)
	case "A.MAP":
		mv := value.Map()
		cmv := mv.(*_A_18_map)
		x.MAP = *cmv.m
	case "A.LIST":
		lv := value.List()
		clv := lv.(*_A_19_list)
		x.LIST = *clv.list
	case "A.ONEOF_B":
		cv := value.Message().Interface().(*B)
		x.ONEOF = &A_ONEOF_B{ONEOF_B: cv}
	case "A.ONEOF_STRING":
		cv := value.String()
		x.ONEOF = &A_ONEOF_STRING{ONEOF_STRING: cv}
	case "A.LIST_ENUM":
		lv := value.List()
		clv := lv.(*_A_22_list)
		x.LIST_ENUM = *clv.list
	case "A.imported":
		x.Imported = value.Message().Interface().(*descriptorpb.FileDescriptorProto)
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: A"))
		}
		panic(fmt.Errorf("message A does not contain field %s", fd.FullName()))
	}
}

// Mutable returns a mutable reference to a composite type.
//
// If the field is unpopulated, it may allocate a composite value.
// For a field belonging to a oneof, it implicitly clears any other field
// that may be currently set within the same oneof.
// For extension fields, it implicitly stores the provided ExtensionType
// if not already stored.
// It panics if the field does not contain a composite type.
//
// Mutable is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_A) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "A.MESSAGE":
		if x.MESSAGE == nil {
			x.MESSAGE = new(B)
		}
		return protoreflect.ValueOfMessage(x.MESSAGE.ProtoReflect())
	case "A.MAP":
		if x.MAP == nil {
			x.MAP = make(map[string]*B)
		}
		value := &_A_18_map{m: &x.MAP}
		return protoreflect.ValueOfMap(value)
	case "A.LIST":
		if x.LIST == nil {
			x.LIST = []*B{}
		}
		value := &_A_19_list{list: &x.LIST}
		return protoreflect.ValueOfList(value)
	case "A.ONEOF_B":
		if x.ONEOF == nil {
			value := &B{}
			oneofValue := &A_ONEOF_B{ONEOF_B: value}
			x.ONEOF = oneofValue
			return protoreflect.ValueOfMessage(value.ProtoReflect())
		}
		switch m := x.ONEOF.(type) {
		case *A_ONEOF_B:
			return protoreflect.ValueOfMessage(m.ONEOF_B.ProtoReflect())
		default:
			value := &B{}
			oneofValue := &A_ONEOF_B{ONEOF_B: value}
			x.ONEOF = oneofValue
			return protoreflect.ValueOfMessage(value.ProtoReflect())
		}
	case "A.LIST_ENUM":
		if x.LIST_ENUM == nil {
			x.LIST_ENUM = []Enumeration{}
		}
		value := &_A_22_list{list: &x.LIST_ENUM}
		return protoreflect.ValueOfList(value)
	case "A.imported":
		if x.Imported == nil {
			x.Imported = new(descriptorpb.FileDescriptorProto)
		}
		return protoreflect.ValueOfMessage(x.Imported.ProtoReflect())
	case "A.enum":
		panic(fmt.Errorf("field enum of message A is not mutable"))
	case "A.some_boolean":
		panic(fmt.Errorf("field some_boolean of message A is not mutable"))
	case "A.INT32":
		panic(fmt.Errorf("field INT32 of message A is not mutable"))
	case "A.SINT32":
		panic(fmt.Errorf("field SINT32 of message A is not mutable"))
	case "A.UINT32":
		panic(fmt.Errorf("field UINT32 of message A is not mutable"))
	case "A.INT64":
		panic(fmt.Errorf("field INT64 of message A is not mutable"))
	case "A.SING64":
		panic(fmt.Errorf("field SING64 of message A is not mutable"))
	case "A.UINT64":
		panic(fmt.Errorf("field UINT64 of message A is not mutable"))
	case "A.SFIXED32":
		panic(fmt.Errorf("field SFIXED32 of message A is not mutable"))
	case "A.FIXED32":
		panic(fmt.Errorf("field FIXED32 of message A is not mutable"))
	case "A.FLOAT":
		panic(fmt.Errorf("field FLOAT of message A is not mutable"))
	case "A.SFIXED64":
		panic(fmt.Errorf("field SFIXED64 of message A is not mutable"))
	case "A.FIXED64":
		panic(fmt.Errorf("field FIXED64 of message A is not mutable"))
	case "A.DOUBLE":
		panic(fmt.Errorf("field DOUBLE of message A is not mutable"))
	case "A.STRING":
		panic(fmt.Errorf("field STRING of message A is not mutable"))
	case "A.BYTES":
		panic(fmt.Errorf("field BYTES of message A is not mutable"))
	case "A.ONEOF_STRING":
		panic(fmt.Errorf("field ONEOF_STRING of message A is not mutable"))
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: A"))
		}
		panic(fmt.Errorf("message A does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_A) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "A.enum":
		return protoreflect.ValueOfEnum(0)
	case "A.some_boolean":
		return protoreflect.ValueOfBool(false)
	case "A.INT32":
		return protoreflect.ValueOfInt32(int32(0))
	case "A.SINT32":
		return protoreflect.ValueOfInt32(int32(0))
	case "A.UINT32":
		return protoreflect.ValueOfUint32(uint32(0))
	case "A.INT64":
		return protoreflect.ValueOfInt64(int64(0))
	case "A.SING64":
		return protoreflect.ValueOfInt64(int64(0))
	case "A.UINT64":
		return protoreflect.ValueOfUint64(uint64(0))
	case "A.SFIXED32":
		return protoreflect.ValueOfInt32(int32(0))
	case "A.FIXED32":
		return protoreflect.ValueOfUint32(uint32(0))
	case "A.FLOAT":
		return protoreflect.ValueOfFloat32(float32(0))
	case "A.SFIXED64":
		return protoreflect.ValueOfInt64(int64(0))
	case "A.FIXED64":
		return protoreflect.ValueOfUint64(uint64(0))
	case "A.DOUBLE":
		return protoreflect.ValueOfFloat64(float64(0))
	case "A.STRING":
		return protoreflect.ValueOfString("")
	case "A.BYTES":
		return protoreflect.ValueOfBytes(nil)
	case "A.MESSAGE":
		m := new(B)
		return protoreflect.ValueOfMessage(m.ProtoReflect())
	case "A.MAP":
		m := make(map[string]*B)
		return protoreflect.ValueOfMap(&_A_18_map{m: &m})
	case "A.LIST":
		list := []*B{}
		return protoreflect.ValueOfList(&_A_19_list{list: &list})
	case "A.ONEOF_B":
		value := &B{}
		return protoreflect.ValueOfMessage(value.ProtoReflect())
	case "A.ONEOF_STRING":
		return protoreflect.ValueOfString("")
	case "A.LIST_ENUM":
		list := []Enumeration{}
		return protoreflect.ValueOfList(&_A_22_list{list: &list})
	case "A.imported":
		m := new(descriptorpb.FileDescriptorProto)
		return protoreflect.ValueOfMessage(m.ProtoReflect())
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: A"))
		}
		panic(fmt.Errorf("message A does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_A) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	case "A.ONEOF":
		if x.ONEOF == nil {
			return nil
		}
		switch x.ONEOF.(type) {
		case *A_ONEOF_B:
			return x.Descriptor().Fields().ByName("ONEOF_B")
		case *A_ONEOF_STRING:
			return x.Descriptor().Fields().ByName("ONEOF_STRING")
		}
	default:
		panic(fmt.Errorf("%s is not a oneof field in A", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_A) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_A) SetUnknown(fields protoreflect.RawFields) {
	x.unknownFields = fields
}

// IsValid reports whether the message is valid.
//
// An invalid message is an empty, read-only value.
//
// An invalid message often corresponds to a nil pointer of the concrete
// message type, but the details are implementation dependent.
// Validity is not part of the protobuf data model, and may not
// be preserved in marshaling or other operations.
func (x *fastReflection_A) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_A) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*A)
		if x == nil {
			return protoiface.SizeOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Size:              0,
			}
		}
		options := runtime.SizeInputToOptions(input)
		_ = options
		var n int
		var l int
		_ = l
		if x.Enum != 0 {
			n += 1 + runtime.Sov(uint64(x.Enum))
		}
		if x.SomeBoolean {
			n += 2
		}
		if x.INT32 != 0 {
			n += 1 + runtime.Sov(uint64(x.INT32))
		}
		if x.SINT32 != 0 {
			n += 1 + runtime.Soz(uint64(x.SINT32))
		}
		if x.UINT32 != 0 {
			n += 1 + runtime.Sov(uint64(x.UINT32))
		}
		if x.INT64 != 0 {
			n += 1 + runtime.Sov(uint64(x.INT64))
		}
		if x.SING64 != 0 {
			n += 1 + runtime.Soz(uint64(x.SING64))
		}
		if x.UINT64 != 0 {
			n += 1 + runtime.Sov(uint64(x.UINT64))
		}
		if x.SFIXED32 != 0 {
			n += 5
		}
		if x.FIXED32 != 0 {
			n += 5
		}
		if x.FLOAT != 0 || math.Signbit(float64(x.FLOAT)) {
			n += 5
		}
		if x.SFIXED64 != 0 {
			n += 9
		}
		if x.FIXED64 != 0 {
			n += 9
		}
		if x.DOUBLE != 0 || math.Signbit(x.DOUBLE) {
			n += 9
		}
		l = len(x.STRING)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		l = len(x.BYTES)
		if l > 0 {
			n += 2 + l + runtime.Sov(uint64(l))
		}
		if x.MESSAGE != nil {
			l = options.Size(x.MESSAGE)
			n += 2 + l + runtime.Sov(uint64(l))
		}
		if len(x.MAP) > 0 {
			sortme := make([]string, 0, len(x.MAP))
			for k := range x.MAP {
				sortme = append(sortme, k)
			}
			sort.Strings(sortme)
			for _, k := range sortme {
				v := x.MAP[k]
				_ = k
				_ = v
				l = 0
				if v != nil {
					l = options.Size(v)
				}
				l += 1 + runtime.Sov(uint64(l))
				mapEntrySize := 1 + len(k) + runtime.Sov(uint64(len(k))) + l
				n += mapEntrySize + 2 + runtime.Sov(uint64(mapEntrySize))
			}
		}
		if len(x.LIST) > 0 {
			for _, e := range x.LIST {
				l = options.Size(e)
				n += 2 + l + runtime.Sov(uint64(l))
			}
		}
		switch x := x.ONEOF.(type) {
		case *A_ONEOF_B:
			if x == nil {
				break
			}
			if x.ONEOF_B != nil {
				l = options.Size(x.ONEOF_B)
				n += 2 + l + runtime.Sov(uint64(l))
			}
		case *A_ONEOF_STRING:
			if x == nil {
				break
			}
			l = len(x.ONEOF_STRING)
			if l > 0 {
				n += 2 + l + runtime.Sov(uint64(l))
			}
		}
		if len(x.LIST_ENUM) > 0 {
			l = 0
			for _, e := range x.LIST_ENUM {
				l += runtime.Sov(uint64(e))
			}
			n += 2 + runtime.Sov(uint64(l)) + l
		}
		if x.Imported != nil {
			l = options.Size(x.Imported)
			n += 2 + l + runtime.Sov(uint64(l))
		}
		if x.unknownFields != nil {
			n += len(x.unknownFields)
		}
		return protoiface.SizeOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Size:              n,
		}
	}

	marshal := func(input protoiface.MarshalInput) (protoiface.MarshalOutput, error) {
		x := input.Message.Interface().(*A)
		if x == nil {
			return protoiface.MarshalOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Buf:               input.Buf,
			}, nil
		}
		options := runtime.MarshalInputToOptions(input)
		_ = options
		size := options.Size(x)
		dAtA := make([]byte, size)
		i := len(dAtA)
		_ = i
		var l int
		_ = l
		if x.unknownFields != nil {
			i -= len(x.unknownFields)
			copy(dAtA[i:], x.unknownFields)
		}
		switch x := x.ONEOF.(type) {
		case *A_ONEOF_B:
			if x.ONEOF_B != nil {
				encoded, err := options.Marshal(x.ONEOF_B)
				if err != nil {
					return protoiface.MarshalOutput{
						NoUnkeyedLiterals: input.NoUnkeyedLiterals,
						Buf:               input.Buf,
					}, err
				}
				i -= len(encoded)
				copy(dAtA[i:], encoded)
				i = runtime.EncodeVarint(dAtA, i, uint64(len(encoded)))
				i--
				dAtA[i] = 0x1
				i--
				dAtA[i] = 0xa2
			}
		case *A_ONEOF_STRING:
			if len(x.ONEOF_STRING) > 0 {
				i -= len(x.ONEOF_STRING)
				copy(dAtA[i:], x.ONEOF_STRING)
				i = runtime.EncodeVarint(dAtA, i, uint64(len(x.ONEOF_STRING)))
				i--
				dAtA[i] = 0x1
				i--
				dAtA[i] = 0xaa
			}
		}
		if x.Imported != nil {
			encoded, err := options.Marshal(x.Imported)
			if err != nil {
				return protoiface.MarshalOutput{
					NoUnkeyedLiterals: input.NoUnkeyedLiterals,
					Buf:               input.Buf,
				}, err
			}
			i -= len(encoded)
			copy(dAtA[i:], encoded)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(encoded)))
			i--
			dAtA[i] = 0x1
			i--
			dAtA[i] = 0xba
		}
		if len(x.LIST_ENUM) > 0 {
			var pksize2 int
			for _, num := range x.LIST_ENUM {
				pksize2 += runtime.Sov(uint64(num))
			}
			i -= pksize2
			j1 := i
			for _, num1 := range x.LIST_ENUM {
				num := uint64(num1)
				for num >= 1<<7 {
					dAtA[j1] = uint8(uint64(num)&0x7f | 0x80)
					num >>= 7
					j1++
				}
				dAtA[j1] = uint8(num)
				j1++
			}
			i = runtime.EncodeVarint(dAtA, i, uint64(pksize2))
			i--
			dAtA[i] = 0x1
			i--
			dAtA[i] = 0xb2
		}
		if len(x.LIST) > 0 {
			for iNdEx := len(x.LIST) - 1; iNdEx >= 0; iNdEx-- {
				encoded, err := options.Marshal(x.LIST[iNdEx])
				if err != nil {
					return protoiface.MarshalOutput{
						NoUnkeyedLiterals: input.NoUnkeyedLiterals,
						Buf:               input.Buf,
					}, err
				}
				i -= len(encoded)
				copy(dAtA[i:], encoded)
				i = runtime.EncodeVarint(dAtA, i, uint64(len(encoded)))
				i--
				dAtA[i] = 0x1
				i--
				dAtA[i] = 0x9a
			}
		}
		if len(x.MAP) > 0 {
			keysForMAP := make([]string, 0, len(x.MAP))
			for k := range x.MAP {
				keysForMAP = append(keysForMAP, string(k))
			}
			sort.Slice(keysForMAP, func(i, j int) bool {
				return keysForMAP[i] < keysForMAP[j]
			})
			for iNdEx := len(keysForMAP) - 1; iNdEx >= 0; iNdEx-- {
				v := x.MAP[string(keysForMAP[iNdEx])]
				baseI := i
				encoded, err := options.Marshal(v)
				if err != nil {
					return protoiface.MarshalOutput{
						NoUnkeyedLiterals: input.NoUnkeyedLiterals,
						Buf:               input.Buf,
					}, err
				}
				i -= len(encoded)
				copy(dAtA[i:], encoded)
				i = runtime.EncodeVarint(dAtA, i, uint64(len(encoded)))
				i--
				dAtA[i] = 0x12
				i -= len(keysForMAP[iNdEx])
				copy(dAtA[i:], keysForMAP[iNdEx])
				i = runtime.EncodeVarint(dAtA, i, uint64(len(keysForMAP[iNdEx])))
				i--
				dAtA[i] = 0xa
				i = runtime.EncodeVarint(dAtA, i, uint64(baseI-i))
				i--
				dAtA[i] = 0x1
				i--
				dAtA[i] = 0x92
			}
		}
		if x.MESSAGE != nil {
			encoded, err := options.Marshal(x.MESSAGE)
			if err != nil {
				return protoiface.MarshalOutput{
					NoUnkeyedLiterals: input.NoUnkeyedLiterals,
					Buf:               input.Buf,
				}, err
			}
			i -= len(encoded)
			copy(dAtA[i:], encoded)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(encoded)))
			i--
			dAtA[i] = 0x1
			i--
			dAtA[i] = 0x8a
		}
		if len(x.BYTES) > 0 {
			i -= len(x.BYTES)
			copy(dAtA[i:], x.BYTES)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.BYTES)))
			i--
			dAtA[i] = 0x1
			i--
			dAtA[i] = 0x82
		}
		if len(x.STRING) > 0 {
			i -= len(x.STRING)
			copy(dAtA[i:], x.STRING)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.STRING)))
			i--
			dAtA[i] = 0x7a
		}
		if x.DOUBLE != 0 || math.Signbit(x.DOUBLE) {
			i -= 8
			binary.LittleEndian.PutUint64(dAtA[i:], uint64(math.Float64bits(float64(x.DOUBLE))))
			i--
			dAtA[i] = 0x71
		}
		if x.FIXED64 != 0 {
			i -= 8
			binary.LittleEndian.PutUint64(dAtA[i:], uint64(x.FIXED64))
			i--
			dAtA[i] = 0x69
		}
		if x.SFIXED64 != 0 {
			i -= 8
			binary.LittleEndian.PutUint64(dAtA[i:], uint64(x.SFIXED64))
			i--
			dAtA[i] = 0x61
		}
		if x.FLOAT != 0 || math.Signbit(float64(x.FLOAT)) {
			i -= 4
			binary.LittleEndian.PutUint32(dAtA[i:], uint32(math.Float32bits(float32(x.FLOAT))))
			i--
			dAtA[i] = 0x5d
		}
		if x.FIXED32 != 0 {
			i -= 4
			binary.LittleEndian.PutUint32(dAtA[i:], uint32(x.FIXED32))
			i--
			dAtA[i] = 0x55
		}
		if x.SFIXED32 != 0 {
			i -= 4
			binary.LittleEndian.PutUint32(dAtA[i:], uint32(x.SFIXED32))
			i--
			dAtA[i] = 0x4d
		}
		if x.UINT64 != 0 {
			i = runtime.EncodeVarint(dAtA, i, uint64(x.UINT64))
			i--
			dAtA[i] = 0x40
		}
		if x.SING64 != 0 {
			i = runtime.EncodeVarint(dAtA, i, uint64((uint64(x.SING64)<<1)^uint64((x.SING64>>63))))
			i--
			dAtA[i] = 0x38
		}
		if x.INT64 != 0 {
			i = runtime.EncodeVarint(dAtA, i, uint64(x.INT64))
			i--
			dAtA[i] = 0x30
		}
		if x.UINT32 != 0 {
			i = runtime.EncodeVarint(dAtA, i, uint64(x.UINT32))
			i--
			dAtA[i] = 0x28
		}
		if x.SINT32 != 0 {
			i = runtime.EncodeVarint(dAtA, i, uint64((uint32(x.SINT32)<<1)^uint32((x.SINT32>>31))))
			i--
			dAtA[i] = 0x20
		}
		if x.INT32 != 0 {
			i = runtime.EncodeVarint(dAtA, i, uint64(x.INT32))
			i--
			dAtA[i] = 0x18
		}
		if x.SomeBoolean {
			i--
			if x.SomeBoolean {
				dAtA[i] = 1
			} else {
				dAtA[i] = 0
			}
			i--
			dAtA[i] = 0x10
		}
		if x.Enum != 0 {
			i = runtime.EncodeVarint(dAtA, i, uint64(x.Enum))
			i--
			dAtA[i] = 0x8
		}
		if input.Buf != nil {
			input.Buf = append(input.Buf, dAtA...)
		} else {
			input.Buf = dAtA
		}
		return protoiface.MarshalOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Buf:               input.Buf,
		}, nil
	}
	unmarshal := func(input protoiface.UnmarshalInput) (protoiface.UnmarshalOutput, error) {
		x := input.Message.Interface().(*A)
		if x == nil {
			return protoiface.UnmarshalOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Flags:             input.Flags,
			}, nil
		}
		options := runtime.UnmarshalInputToOptions(input)
		_ = options
		dAtA := input.Buf
		l := len(dAtA)
		iNdEx := 0
		for iNdEx < l {
			preIndex := iNdEx
			var wire uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
				}
				if iNdEx >= l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				wire |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			fieldNum := int32(wire >> 3)
			wireType := int(wire & 0x7)
			if wireType == 4 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: A: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: A: illegal tag %d (wire type %d)", fieldNum, wire)
			}
			switch fieldNum {
			case 1:
				if wireType != 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Enum", wireType)
				}
				x.Enum = 0
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					x.Enum |= Enumeration(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
			case 2:
				if wireType != 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field SomeBoolean", wireType)
				}
				var v int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					v |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				x.SomeBoolean = bool(v != 0)
			case 3:
				if wireType != 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field INT32", wireType)
				}
				x.INT32 = 0
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					x.INT32 |= int32(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
			case 4:
				if wireType != 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field SINT32", wireType)
				}
				var v int32
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					v |= int32(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				v = int32((uint32(v) >> 1) ^ uint32(((v&1)<<31)>>31))
				x.SINT32 = v
			case 5:
				if wireType != 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field UINT32", wireType)
				}
				x.UINT32 = 0
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					x.UINT32 |= uint32(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
			case 6:
				if wireType != 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field INT64", wireType)
				}
				x.INT64 = 0
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					x.INT64 |= int64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
			case 7:
				if wireType != 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field SING64", wireType)
				}
				var v uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					v |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				v = (v >> 1) ^ uint64((int64(v&1)<<63)>>63)
				x.SING64 = int64(v)
			case 8:
				if wireType != 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field UINT64", wireType)
				}
				x.UINT64 = 0
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					x.UINT64 |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
			case 9:
				if wireType != 5 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field SFIXED32", wireType)
				}
				x.SFIXED32 = 0
				if (iNdEx + 4) > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				x.SFIXED32 = int32(binary.LittleEndian.Uint32(dAtA[iNdEx:]))
				iNdEx += 4
			case 10:
				if wireType != 5 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field FIXED32", wireType)
				}
				x.FIXED32 = 0
				if (iNdEx + 4) > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				x.FIXED32 = uint32(binary.LittleEndian.Uint32(dAtA[iNdEx:]))
				iNdEx += 4
			case 11:
				if wireType != 5 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field FLOAT", wireType)
				}
				var v uint32
				if (iNdEx + 4) > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				v = uint32(binary.LittleEndian.Uint32(dAtA[iNdEx:]))
				iNdEx += 4
				x.FLOAT = float32(math.Float32frombits(v))
			case 12:
				if wireType != 1 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field SFIXED64", wireType)
				}
				x.SFIXED64 = 0
				if (iNdEx + 8) > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				x.SFIXED64 = int64(binary.LittleEndian.Uint64(dAtA[iNdEx:]))
				iNdEx += 8
			case 13:
				if wireType != 1 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field FIXED64", wireType)
				}
				x.FIXED64 = 0
				if (iNdEx + 8) > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				x.FIXED64 = uint64(binary.LittleEndian.Uint64(dAtA[iNdEx:]))
				iNdEx += 8
			case 14:
				if wireType != 1 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field DOUBLE", wireType)
				}
				var v uint64
				if (iNdEx + 8) > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				v = uint64(binary.LittleEndian.Uint64(dAtA[iNdEx:]))
				iNdEx += 8
				x.DOUBLE = float64(math.Float64frombits(v))
			case 15:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field STRING", wireType)
				}
				var stringLen uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					stringLen |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				intStringLen := int(stringLen)
				if intStringLen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + intStringLen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				x.STRING = string(dAtA[iNdEx:postIndex])
				iNdEx = postIndex
			case 16:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field BYTES", wireType)
				}
				var byteLen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					byteLen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if byteLen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + byteLen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				x.BYTES = append(x.BYTES[:0], dAtA[iNdEx:postIndex]...)
				if x.BYTES == nil {
					x.BYTES = []byte{}
				}
				iNdEx = postIndex
			case 17:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field MESSAGE", wireType)
				}
				var msglen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					msglen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if msglen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + msglen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				if x.MESSAGE == nil {
					x.MESSAGE = &B{}
				}
				if err := options.Unmarshal(dAtA[iNdEx:postIndex], x.MESSAGE); err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				iNdEx = postIndex
			case 18:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field MAP", wireType)
				}
				var msglen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					msglen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if msglen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + msglen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				if x.MAP == nil {
					x.MAP = make(map[string]*B)
				}
				var mapkey string
				var mapvalue *B
				for iNdEx < postIndex {
					entryPreIndex := iNdEx
					var wire uint64
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
						}
						if iNdEx >= l {
							return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						wire |= uint64(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					fieldNum := int32(wire >> 3)
					if fieldNum == 1 {
						var stringLenmapkey uint64
						for shift := uint(0); ; shift += 7 {
							if shift >= 64 {
								return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
							}
							if iNdEx >= l {
								return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
							}
							b := dAtA[iNdEx]
							iNdEx++
							stringLenmapkey |= uint64(b&0x7F) << shift
							if b < 0x80 {
								break
							}
						}
						intStringLenmapkey := int(stringLenmapkey)
						if intStringLenmapkey < 0 {
							return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
						}
						postStringIndexmapkey := iNdEx + intStringLenmapkey
						if postStringIndexmapkey < 0 {
							return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
						}
						if postStringIndexmapkey > l {
							return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
						}
						mapkey = string(dAtA[iNdEx:postStringIndexmapkey])
						iNdEx = postStringIndexmapkey
					} else if fieldNum == 2 {
						var mapmsglen int
						for shift := uint(0); ; shift += 7 {
							if shift >= 64 {
								return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
							}
							if iNdEx >= l {
								return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
							}
							b := dAtA[iNdEx]
							iNdEx++
							mapmsglen |= int(b&0x7F) << shift
							if b < 0x80 {
								break
							}
						}
						if mapmsglen < 0 {
							return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
						}
						postmsgIndex := iNdEx + mapmsglen
						if postmsgIndex < 0 {
							return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
						}
						if postmsgIndex > l {
							return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
						}
						mapvalue = &B{}
						if err := options.Unmarshal(dAtA[iNdEx:postmsgIndex], mapvalue); err != nil {
							return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
						}
						iNdEx = postmsgIndex
					} else {
						iNdEx = entryPreIndex
						skippy, err := runtime.Skip(dAtA[iNdEx:])
						if err != nil {
							return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
						}
						if (skippy < 0) || (iNdEx+skippy) < 0 {
							return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
						}
						if (iNdEx + skippy) > postIndex {
							return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
						}
						iNdEx += skippy
					}
				}
				x.MAP[mapkey] = mapvalue
				iNdEx = postIndex
			case 19:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field LIST", wireType)
				}
				var msglen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					msglen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if msglen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + msglen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				x.LIST = append(x.LIST, &B{})
				if err := options.Unmarshal(dAtA[iNdEx:postIndex], x.LIST[len(x.LIST)-1]); err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				iNdEx = postIndex
			case 20:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field ONEOF_B", wireType)
				}
				var msglen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					msglen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if msglen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + msglen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				v := &B{}
				if err := options.Unmarshal(dAtA[iNdEx:postIndex], v); err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				x.ONEOF = &A_ONEOF_B{v}
				iNdEx = postIndex
			case 21:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field ONEOF_STRING", wireType)
				}
				var stringLen uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					stringLen |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				intStringLen := int(stringLen)
				if intStringLen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + intStringLen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				x.ONEOF = &A_ONEOF_STRING{string(dAtA[iNdEx:postIndex])}
				iNdEx = postIndex
			case 22:
				if wireType == 0 {
					var v Enumeration
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
						}
						if iNdEx >= l {
							return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						v |= Enumeration(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					x.LIST_ENUM = append(x.LIST_ENUM, v)
				} else if wireType == 2 {
					var packedLen int
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
						}
						if iNdEx >= l {
							return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						packedLen |= int(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					if packedLen < 0 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
					}
					postIndex := iNdEx + packedLen
					if postIndex < 0 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
					}
					if postIndex > l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					var elementCount int
					if elementCount != 0 && len(x.LIST_ENUM) == 0 {
						x.LIST_ENUM = make([]Enumeration, 0, elementCount)
					}
					for iNdEx < postIndex {
						var v Enumeration
						for shift := uint(0); ; shift += 7 {
							if shift >= 64 {
								return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
							}
							if iNdEx >= l {
								return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
							}
							b := dAtA[iNdEx]
							iNdEx++
							v |= Enumeration(b&0x7F) << shift
							if b < 0x80 {
								break
							}
						}
						x.LIST_ENUM = append(x.LIST_ENUM, v)
					}
				} else {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field LIST_ENUM", wireType)
				}
			case 23:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Imported", wireType)
				}
				var msglen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					msglen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if msglen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + msglen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				if x.Imported == nil {
					x.Imported = &descriptorpb.FileDescriptorProto{}
				}
				if err := options.Unmarshal(dAtA[iNdEx:postIndex], x.Imported); err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				iNdEx = postIndex
			default:
				iNdEx = preIndex
				skippy, err := runtime.Skip(dAtA[iNdEx:])
				if err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				if (skippy < 0) || (iNdEx+skippy) < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if (iNdEx + skippy) > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				if !options.DiscardUnknown {
					x.unknownFields = append(x.unknownFields, dAtA[iNdEx:iNdEx+skippy]...)
				}
				iNdEx += skippy
			}
		}

		if iNdEx > l {
			return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
		}
		return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, nil
	}
	return &protoiface.Methods{
		NoUnkeyedLiterals: struct{}{},
		Flags:             protoiface.SupportMarshalDeterministic,
		Size:              size,
		Marshal:           marshal,
		Unmarshal:         unmarshal,
		Merge:             nil,
		CheckInitialized:  nil,
	}
}

var (
	md_B   protoreflect.MessageDescriptor
	fd_B_x protoreflect.FieldDescriptor
)

func init() {
	file_testpb_1_proto_init()
	md_B = File_testpb_1_proto.Messages().ByName("B")
	fd_B_x = md_B.Fields().ByName("x")
}

var _ protoreflect.Message = (*fastReflection_B)(nil)

type fastReflection_B B

func (x *B) ProtoReflect() protoreflect.Message {
	return (*fastReflection_B)(x)
}

func (x *B) slowProtoReflect() protoreflect.Message {
	mi := &file_testpb_1_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_B_messageType fastReflection_B_messageType
var _ protoreflect.MessageType = fastReflection_B_messageType{}

type fastReflection_B_messageType struct{}

func (x fastReflection_B_messageType) Zero() protoreflect.Message {
	return (*fastReflection_B)(nil)
}
func (x fastReflection_B_messageType) New() protoreflect.Message {
	return new(fastReflection_B)
}
func (x fastReflection_B_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_B
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_B) Descriptor() protoreflect.MessageDescriptor {
	return md_B
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_B) Type() protoreflect.MessageType {
	return _fastReflection_B_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_B) New() protoreflect.Message {
	return new(fastReflection_B)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_B) Interface() protoreflect.ProtoMessage {
	return (*B)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_B) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
	if x.X != "" {
		value := protoreflect.ValueOfString(x.X)
		if !f(fd_B_x, value) {
			return
		}
	}
}

// Has reports whether a field is populated.
//
// Some fields have the property of nullability where it is possible to
// distinguish between the default value of a field and whether the field
// was explicitly populated with the default value. Singular message fields,
// member fields of a oneof, and proto2 scalar fields are nullable. Such
// fields are populated only if explicitly set.
//
// In other cases (aside from the nullable cases above),
// a proto3 scalar field is populated if it contains a non-zero value, and
// a repeated field is populated if it is non-empty.
func (x *fastReflection_B) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	case "B.x":
		return x.X != ""
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: B"))
		}
		panic(fmt.Errorf("message B does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_B) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	case "B.x":
		x.X = ""
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: B"))
		}
		panic(fmt.Errorf("message B does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_B) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	case "B.x":
		value := x.X
		return protoreflect.ValueOfString(value)
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: B"))
		}
		panic(fmt.Errorf("message B does not contain field %s", descriptor.FullName()))
	}
}

// Set stores the value for a field.
//
// For a field belonging to a oneof, it implicitly clears any other field
// that may be currently set within the same oneof.
// For extension fields, it implicitly stores the provided ExtensionType.
// When setting a composite type, it is unspecified whether the stored value
// aliases the source's memory in any way. If the composite value is an
// empty, read-only value, then it panics.
//
// Set is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_B) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	case "B.x":
		x.X = value.String()
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: B"))
		}
		panic(fmt.Errorf("message B does not contain field %s", fd.FullName()))
	}
}

// Mutable returns a mutable reference to a composite type.
//
// If the field is unpopulated, it may allocate a composite value.
// For a field belonging to a oneof, it implicitly clears any other field
// that may be currently set within the same oneof.
// For extension fields, it implicitly stores the provided ExtensionType
// if not already stored.
// It panics if the field does not contain a composite type.
//
// Mutable is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_B) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "B.x":
		panic(fmt.Errorf("field x of message B is not mutable"))
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: B"))
		}
		panic(fmt.Errorf("message B does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_B) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "B.x":
		return protoreflect.ValueOfString("")
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: B"))
		}
		panic(fmt.Errorf("message B does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_B) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in B", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_B) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_B) SetUnknown(fields protoreflect.RawFields) {
	x.unknownFields = fields
}

// IsValid reports whether the message is valid.
//
// An invalid message is an empty, read-only value.
//
// An invalid message often corresponds to a nil pointer of the concrete
// message type, but the details are implementation dependent.
// Validity is not part of the protobuf data model, and may not
// be preserved in marshaling or other operations.
func (x *fastReflection_B) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_B) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*B)
		if x == nil {
			return protoiface.SizeOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Size:              0,
			}
		}
		options := runtime.SizeInputToOptions(input)
		_ = options
		var n int
		var l int
		_ = l
		l = len(x.X)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		if x.unknownFields != nil {
			n += len(x.unknownFields)
		}
		return protoiface.SizeOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Size:              n,
		}
	}

	marshal := func(input protoiface.MarshalInput) (protoiface.MarshalOutput, error) {
		x := input.Message.Interface().(*B)
		if x == nil {
			return protoiface.MarshalOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Buf:               input.Buf,
			}, nil
		}
		options := runtime.MarshalInputToOptions(input)
		_ = options
		size := options.Size(x)
		dAtA := make([]byte, size)
		i := len(dAtA)
		_ = i
		var l int
		_ = l
		if x.unknownFields != nil {
			i -= len(x.unknownFields)
			copy(dAtA[i:], x.unknownFields)
		}
		if len(x.X) > 0 {
			i -= len(x.X)
			copy(dAtA[i:], x.X)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.X)))
			i--
			dAtA[i] = 0xa
		}
		if input.Buf != nil {
			input.Buf = append(input.Buf, dAtA...)
		} else {
			input.Buf = dAtA
		}
		return protoiface.MarshalOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Buf:               input.Buf,
		}, nil
	}
	unmarshal := func(input protoiface.UnmarshalInput) (protoiface.UnmarshalOutput, error) {
		x := input.Message.Interface().(*B)
		if x == nil {
			return protoiface.UnmarshalOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Flags:             input.Flags,
			}, nil
		}
		options := runtime.UnmarshalInputToOptions(input)
		_ = options
		dAtA := input.Buf
		l := len(dAtA)
		iNdEx := 0
		for iNdEx < l {
			preIndex := iNdEx
			var wire uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
				}
				if iNdEx >= l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				wire |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			fieldNum := int32(wire >> 3)
			wireType := int(wire & 0x7)
			if wireType == 4 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: B: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: B: illegal tag %d (wire type %d)", fieldNum, wire)
			}
			switch fieldNum {
			case 1:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field X", wireType)
				}
				var stringLen uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					stringLen |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				intStringLen := int(stringLen)
				if intStringLen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + intStringLen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				x.X = string(dAtA[iNdEx:postIndex])
				iNdEx = postIndex
			default:
				iNdEx = preIndex
				skippy, err := runtime.Skip(dAtA[iNdEx:])
				if err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				if (skippy < 0) || (iNdEx+skippy) < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if (iNdEx + skippy) > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				if !options.DiscardUnknown {
					x.unknownFields = append(x.unknownFields, dAtA[iNdEx:iNdEx+skippy]...)
				}
				iNdEx += skippy
			}
		}

		if iNdEx > l {
			return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
		}
		return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, nil
	}
	return &protoiface.Methods{
		NoUnkeyedLiterals: struct{}{},
		Flags:             protoiface.SupportMarshalDeterministic,
		Size:              size,
		Marshal:           marshal,
		Unmarshal:         unmarshal,
		Merge:             nil,
		CheckInitialized:  nil,
	}
}

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.0
// 	protoc        v3.15.7
// source: testpb/1.proto

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Enumeration int32

const (
	Enumeration_One Enumeration = 0
	Enumeration_Two Enumeration = 1
)

// Enum value maps for Enumeration.
var (
	Enumeration_name = map[int32]string{
		0: "One",
		1: "Two",
	}
	Enumeration_value = map[string]int32{
		"One": 0,
		"Two": 1,
	}
)

func (x Enumeration) Enum() *Enumeration {
	p := new(Enumeration)
	*p = x
	return p
}

func (x Enumeration) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Enumeration) Descriptor() protoreflect.EnumDescriptor {
	return file_testpb_1_proto_enumTypes[0].Descriptor()
}

func (Enumeration) Type() protoreflect.EnumType {
	return &file_testpb_1_proto_enumTypes[0]
}

func (x Enumeration) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Enumeration.Descriptor instead.
func (Enumeration) EnumDescriptor() ([]byte, []int) {
	return file_testpb_1_proto_rawDescGZIP(), []int{0}
}

type A struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Enum        Enumeration   `protobuf:"varint,1,opt,name=enum,proto3,enum=Enumeration" json:"enum,omitempty"`
	SomeBoolean bool          `protobuf:"varint,2,opt,name=some_boolean,json=someBoolean,proto3" json:"some_boolean,omitempty"`
	INT32       int32         `protobuf:"varint,3,opt,name=INT32,proto3" json:"INT32,omitempty"`
	SINT32      int32         `protobuf:"zigzag32,4,opt,name=SINT32,proto3" json:"SINT32,omitempty"`
	UINT32      uint32        `protobuf:"varint,5,opt,name=UINT32,proto3" json:"UINT32,omitempty"`
	INT64       int64         `protobuf:"varint,6,opt,name=INT64,proto3" json:"INT64,omitempty"`
	SING64      int64         `protobuf:"zigzag64,7,opt,name=SING64,proto3" json:"SING64,omitempty"`
	UINT64      uint64        `protobuf:"varint,8,opt,name=UINT64,proto3" json:"UINT64,omitempty"`
	SFIXED32    int32         `protobuf:"fixed32,9,opt,name=SFIXED32,proto3" json:"SFIXED32,omitempty"`
	FIXED32     uint32        `protobuf:"fixed32,10,opt,name=FIXED32,proto3" json:"FIXED32,omitempty"`
	FLOAT       float32       `protobuf:"fixed32,11,opt,name=FLOAT,proto3" json:"FLOAT,omitempty"`
	SFIXED64    int64         `protobuf:"fixed64,12,opt,name=SFIXED64,proto3" json:"SFIXED64,omitempty"`
	FIXED64     uint64        `protobuf:"fixed64,13,opt,name=FIXED64,proto3" json:"FIXED64,omitempty"`
	DOUBLE      float64       `protobuf:"fixed64,14,opt,name=DOUBLE,proto3" json:"DOUBLE,omitempty"`
	STRING      string        `protobuf:"bytes,15,opt,name=STRING,proto3" json:"STRING,omitempty"`
	BYTES       []byte        `protobuf:"bytes,16,opt,name=BYTES,proto3" json:"BYTES,omitempty"`
	MESSAGE     *B            `protobuf:"bytes,17,opt,name=MESSAGE,proto3" json:"MESSAGE,omitempty"`
	MAP         map[string]*B `protobuf:"bytes,18,rep,name=MAP,proto3" json:"MAP,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	LIST        []*B          `protobuf:"bytes,19,rep,name=LIST,proto3" json:"LIST,omitempty"`
	// Types that are assignable to ONEOF:
	//	*A_ONEOF_B
	//	*A_ONEOF_STRING
	ONEOF     isA_ONEOF                         `protobuf_oneof:"ONEOF"`
	LIST_ENUM []Enumeration                     `protobuf:"varint,22,rep,packed,name=LIST_ENUM,json=LISTENUM,proto3,enum=Enumeration" json:"LIST_ENUM,omitempty"`
	Imported  *descriptorpb.FileDescriptorProto `protobuf:"bytes,23,opt,name=imported,proto3" json:"imported,omitempty"`
}

func (x *A) Reset() {
	*x = A{}
	if protoimpl.UnsafeEnabled {
		mi := &file_testpb_1_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *A) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*A) ProtoMessage() {}

// Deprecated: Use A.ProtoReflect.Descriptor instead.
func (*A) Descriptor() ([]byte, []int) {
	return file_testpb_1_proto_rawDescGZIP(), []int{0}
}

func (x *A) GetEnum() Enumeration {
	if x != nil {
		return x.Enum
	}
	return Enumeration_One
}

func (x *A) GetSomeBoolean() bool {
	if x != nil {
		return x.SomeBoolean
	}
	return false
}

func (x *A) GetINT32() int32 {
	if x != nil {
		return x.INT32
	}
	return 0
}

func (x *A) GetSINT32() int32 {
	if x != nil {
		return x.SINT32
	}
	return 0
}

func (x *A) GetUINT32() uint32 {
	if x != nil {
		return x.UINT32
	}
	return 0
}

func (x *A) GetINT64() int64 {
	if x != nil {
		return x.INT64
	}
	return 0
}

func (x *A) GetSING64() int64 {
	if x != nil {
		return x.SING64
	}
	return 0
}

func (x *A) GetUINT64() uint64 {
	if x != nil {
		return x.UINT64
	}
	return 0
}

func (x *A) GetSFIXED32() int32 {
	if x != nil {
		return x.SFIXED32
	}
	return 0
}

func (x *A) GetFIXED32() uint32 {
	if x != nil {
		return x.FIXED32
	}
	return 0
}

func (x *A) GetFLOAT() float32 {
	if x != nil {
		return x.FLOAT
	}
	return 0
}

func (x *A) GetSFIXED64() int64 {
	if x != nil {
		return x.SFIXED64
	}
	return 0
}

func (x *A) GetFIXED64() uint64 {
	if x != nil {
		return x.FIXED64
	}
	return 0
}

func (x *A) GetDOUBLE() float64 {
	if x != nil {
		return x.DOUBLE
	}
	return 0
}

func (x *A) GetSTRING() string {
	if x != nil {
		return x.STRING
	}
	return ""
}

func (x *A) GetBYTES() []byte {
	if x != nil {
		return x.BYTES
	}
	return nil
}

func (x *A) GetMESSAGE() *B {
	if x != nil {
		return x.MESSAGE
	}
	return nil
}

func (x *A) GetMAP() map[string]*B {
	if x != nil {
		return x.MAP
	}
	return nil
}

func (x *A) GetLIST() []*B {
	if x != nil {
		return x.LIST
	}
	return nil
}

func (m *A) GetONEOF() isA_ONEOF {
	if m != nil {
		return m.ONEOF
	}
	return nil
}

func (x *A) GetONEOF_B() *B {
	if x, ok := x.GetONEOF().(*A_ONEOF_B); ok {
		return x.ONEOF_B
	}
	return nil
}

func (x *A) GetONEOF_STRING() string {
	if x, ok := x.GetONEOF().(*A_ONEOF_STRING); ok {
		return x.ONEOF_STRING
	}
	return ""
}

func (x *A) GetLIST_ENUM() []Enumeration {
	if x != nil {
		return x.LIST_ENUM
	}
	return nil
}

func (x *A) GetImported() *descriptorpb.FileDescriptorProto {
	if x != nil {
		return x.Imported
	}
	return nil
}

type isA_ONEOF interface {
	isA_ONEOF()
}

type A_ONEOF_B struct {
	ONEOF_B *B `protobuf:"bytes,20,opt,name=ONEOF_B,json=ONEOFB,proto3,oneof"`
}

type A_ONEOF_STRING struct {
	ONEOF_STRING string `protobuf:"bytes,21,opt,name=ONEOF_STRING,json=ONEOFSTRING,proto3,oneof"`
}

func (*A_ONEOF_B) isA_ONEOF() {}

func (*A_ONEOF_STRING) isA_ONEOF() {}

type B struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	X string `protobuf:"bytes,1,opt,name=x,proto3" json:"x,omitempty"`
}

func (x *B) Reset() {
	*x = B{}
	if protoimpl.UnsafeEnabled {
		mi := &file_testpb_1_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *B) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*B) ProtoMessage() {}

// Deprecated: Use B.ProtoReflect.Descriptor instead.
func (*B) Descriptor() ([]byte, []int) {
	return file_testpb_1_proto_rawDescGZIP(), []int{1}
}

func (x *B) GetX() string {
	if x != nil {
		return x.X
	}
	return ""
}

var File_testpb_1_proto protoreflect.FileDescriptor

var file_testpb_1_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x74, 0x65, 0x73, 0x74, 0x70, 0x62, 0x2f, 0x31, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x20, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0xe7, 0x05, 0x0a, 0x01, 0x41, 0x12, 0x20, 0x0a, 0x04, 0x65, 0x6e, 0x75, 0x6d,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0c, 0x2e, 0x45, 0x6e, 0x75, 0x6d, 0x65, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x04, 0x65, 0x6e, 0x75, 0x6d, 0x12, 0x21, 0x0a, 0x0c, 0x73, 0x6f,
	0x6d, 0x65, 0x5f, 0x62, 0x6f, 0x6f, 0x6c, 0x65, 0x61, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x0b, 0x73, 0x6f, 0x6d, 0x65, 0x42, 0x6f, 0x6f, 0x6c, 0x65, 0x61, 0x6e, 0x12, 0x14, 0x0a,
	0x05, 0x49, 0x4e, 0x54, 0x33, 0x32, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x49, 0x4e,
	0x54, 0x33, 0x32, 0x12, 0x16, 0x0a, 0x06, 0x53, 0x49, 0x4e, 0x54, 0x33, 0x32, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x11, 0x52, 0x06, 0x53, 0x49, 0x4e, 0x54, 0x33, 0x32, 0x12, 0x16, 0x0a, 0x06, 0x55,
	0x49, 0x4e, 0x54, 0x33, 0x32, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x55, 0x49, 0x4e,
	0x54, 0x33, 0x32, 0x12, 0x14, 0x0a, 0x05, 0x49, 0x4e, 0x54, 0x36, 0x34, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x05, 0x49, 0x4e, 0x54, 0x36, 0x34, 0x12, 0x16, 0x0a, 0x06, 0x53, 0x49, 0x4e,
	0x47, 0x36, 0x34, 0x18, 0x07, 0x20, 0x01, 0x28, 0x12, 0x52, 0x06, 0x53, 0x49, 0x4e, 0x47, 0x36,
	0x34, 0x12, 0x16, 0x0a, 0x06, 0x55, 0x49, 0x4e, 0x54, 0x36, 0x34, 0x18, 0x08, 0x20, 0x01, 0x28,
	0x04, 0x52, 0x06, 0x55, 0x49, 0x4e, 0x54, 0x36, 0x34, 0x12, 0x1a, 0x0a, 0x08, 0x53, 0x46, 0x49,
	0x58, 0x45, 0x44, 0x33, 0x32, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0f, 0x52, 0x08, 0x53, 0x46, 0x49,
	0x58, 0x45, 0x44, 0x33, 0x32, 0x12, 0x18, 0x0a, 0x07, 0x46, 0x49, 0x58, 0x45, 0x44, 0x33, 0x32,
	0x18, 0x0a, 0x20, 0x01, 0x28, 0x07, 0x52, 0x07, 0x46, 0x49, 0x58, 0x45, 0x44, 0x33, 0x32, 0x12,
	0x14, 0x0a, 0x05, 0x46, 0x4c, 0x4f, 0x41, 0x54, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x02, 0x52, 0x05,
	0x46, 0x4c, 0x4f, 0x41, 0x54, 0x12, 0x1a, 0x0a, 0x08, 0x53, 0x46, 0x49, 0x58, 0x45, 0x44, 0x36,
	0x34, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x10, 0x52, 0x08, 0x53, 0x46, 0x49, 0x58, 0x45, 0x44, 0x36,
	0x34, 0x12, 0x18, 0x0a, 0x07, 0x46, 0x49, 0x58, 0x45, 0x44, 0x36, 0x34, 0x18, 0x0d, 0x20, 0x01,
	0x28, 0x06, 0x52, 0x07, 0x46, 0x49, 0x58, 0x45, 0x44, 0x36, 0x34, 0x12, 0x16, 0x0a, 0x06, 0x44,
	0x4f, 0x55, 0x42, 0x4c, 0x45, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x01, 0x52, 0x06, 0x44, 0x4f, 0x55,
	0x42, 0x4c, 0x45, 0x12, 0x16, 0x0a, 0x06, 0x53, 0x54, 0x52, 0x49, 0x4e, 0x47, 0x18, 0x0f, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x53, 0x54, 0x52, 0x49, 0x4e, 0x47, 0x12, 0x14, 0x0a, 0x05, 0x42,
	0x59, 0x54, 0x45, 0x53, 0x18, 0x10, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x05, 0x42, 0x59, 0x54, 0x45,
	0x53, 0x12, 0x1c, 0x0a, 0x07, 0x4d, 0x45, 0x53, 0x53, 0x41, 0x47, 0x45, 0x18, 0x11, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x02, 0x2e, 0x42, 0x52, 0x07, 0x4d, 0x45, 0x53, 0x53, 0x41, 0x47, 0x45, 0x12,
	0x1d, 0x0a, 0x03, 0x4d, 0x41, 0x50, 0x18, 0x12, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x41,
	0x2e, 0x4d, 0x41, 0x50, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x03, 0x4d, 0x41, 0x50, 0x12, 0x16,
	0x0a, 0x04, 0x4c, 0x49, 0x53, 0x54, 0x18, 0x13, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x02, 0x2e, 0x42,
	0x52, 0x04, 0x4c, 0x49, 0x53, 0x54, 0x12, 0x1d, 0x0a, 0x07, 0x4f, 0x4e, 0x45, 0x4f, 0x46, 0x5f,
	0x42, 0x18, 0x14, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x02, 0x2e, 0x42, 0x48, 0x00, 0x52, 0x06, 0x4f,
	0x4e, 0x45, 0x4f, 0x46, 0x42, 0x12, 0x23, 0x0a, 0x0c, 0x4f, 0x4e, 0x45, 0x4f, 0x46, 0x5f, 0x53,
	0x54, 0x52, 0x49, 0x4e, 0x47, 0x18, 0x15, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x0b, 0x4f,
	0x4e, 0x45, 0x4f, 0x46, 0x53, 0x54, 0x52, 0x49, 0x4e, 0x47, 0x12, 0x29, 0x0a, 0x09, 0x4c, 0x49,
	0x53, 0x54, 0x5f, 0x45, 0x4e, 0x55, 0x4d, 0x18, 0x16, 0x20, 0x03, 0x28, 0x0e, 0x32, 0x0c, 0x2e,
	0x45, 0x6e, 0x75, 0x6d, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x08, 0x4c, 0x49, 0x53,
	0x54, 0x45, 0x4e, 0x55, 0x4d, 0x12, 0x40, 0x0a, 0x08, 0x69, 0x6d, 0x70, 0x6f, 0x72, 0x74, 0x65,
	0x64, 0x18, 0x17, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x46, 0x69, 0x6c, 0x65, 0x44, 0x65,
	0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x52, 0x08, 0x69,
	0x6d, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x64, 0x1a, 0x3a, 0x0a, 0x08, 0x4d, 0x41, 0x50, 0x45, 0x6e,
	0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x18, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x02, 0x2e, 0x42, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a,
	0x02, 0x38, 0x01, 0x42, 0x07, 0x0a, 0x05, 0x4f, 0x4e, 0x45, 0x4f, 0x46, 0x22, 0x11, 0x0a, 0x01,
	0x42, 0x12, 0x0c, 0x0a, 0x01, 0x78, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x01, 0x78, 0x2a,
	0x1f, 0x0a, 0x0b, 0x45, 0x6e, 0x75, 0x6d, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x07,
	0x0a, 0x03, 0x4f, 0x6e, 0x65, 0x10, 0x00, 0x12, 0x07, 0x0a, 0x03, 0x54, 0x77, 0x6f, 0x10, 0x01,
	0x42, 0x27, 0x5a, 0x25, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63,
	0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2f, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2d, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_testpb_1_proto_rawDescOnce sync.Once
	file_testpb_1_proto_rawDescData = file_testpb_1_proto_rawDesc
)

func file_testpb_1_proto_rawDescGZIP() []byte {
	file_testpb_1_proto_rawDescOnce.Do(func() {
		file_testpb_1_proto_rawDescData = protoimpl.X.CompressGZIP(file_testpb_1_proto_rawDescData)
	})
	return file_testpb_1_proto_rawDescData
}

var file_testpb_1_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_testpb_1_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_testpb_1_proto_goTypes = []interface{}{
	(Enumeration)(0),                         // 0: Enumeration
	(*A)(nil),                                // 1: A
	(*B)(nil),                                // 2: B
	nil,                                      // 3: A.MAPEntry
	(*descriptorpb.FileDescriptorProto)(nil), // 4: google.protobuf.FileDescriptorProto
}
var file_testpb_1_proto_depIdxs = []int32{
	0, // 0: A.enum:type_name -> Enumeration
	2, // 1: A.MESSAGE:type_name -> B
	3, // 2: A.MAP:type_name -> A.MAPEntry
	2, // 3: A.LIST:type_name -> B
	2, // 4: A.ONEOF_B:type_name -> B
	0, // 5: A.LIST_ENUM:type_name -> Enumeration
	4, // 6: A.imported:type_name -> google.protobuf.FileDescriptorProto
	2, // 7: A.MAPEntry.value:type_name -> B
	8, // [8:8] is the sub-list for method output_type
	8, // [8:8] is the sub-list for method input_type
	8, // [8:8] is the sub-list for extension type_name
	8, // [8:8] is the sub-list for extension extendee
	0, // [0:8] is the sub-list for field type_name
}

func init() { file_testpb_1_proto_init() }
func file_testpb_1_proto_init() {
	if File_testpb_1_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_testpb_1_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*A); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_testpb_1_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*B); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	file_testpb_1_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*A_ONEOF_B)(nil),
		(*A_ONEOF_STRING)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_testpb_1_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_testpb_1_proto_goTypes,
		DependencyIndexes: file_testpb_1_proto_depIdxs,
		EnumInfos:         file_testpb_1_proto_enumTypes,
		MessageInfos:      file_testpb_1_proto_msgTypes,
	}.Build()
	File_testpb_1_proto = out.File
	file_testpb_1_proto_rawDesc = nil
	file_testpb_1_proto_goTypes = nil
	file_testpb_1_proto_depIdxs = nil
}
