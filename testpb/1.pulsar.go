package testpb

import (
	fmt "fmt"

	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoiface "google.golang.org/protobuf/runtime/protoiface"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
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
	if x.FLOAT != float32(0) {
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
	if x.DOUBLE != float64(0) {
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
			if !f(fd_A_ONEOF_B, value) {
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
		return x.FLOAT != float32(0)
	case "A.SFIXED64":
		return x.SFIXED64 != int64(0)
	case "A.FIXED64":
		return x.FIXED64 != uint64(0)
	case "A.DOUBLE":
		return x.DOUBLE != float64(0)
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
	return nil
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
	return nil
}
