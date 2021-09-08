// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.15.7
// source: testpb/1.proto

package testpb

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

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
	ONEOF     isA_ONEOF     `protobuf_oneof:"ONEOF"`
	LIST_ENUM []Enumeration `protobuf:"varint,22,rep,packed,name=LIST_ENUM,json=LISTENUM,proto3,enum=Enumeration" json:"LIST_ENUM,omitempty"`
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

// Deprecated: Use B.ProtoReflect.Descriptor instead.

// 0: Enumeration
// 1: A
// 2: B
// 3: A.MAPEntry

// 0: A.enum:type_name -> Enumeration
// 1: A.MESSAGE:type_name -> B
// 2: A.MAP:type_name -> A.MAPEntry
// 3: A.LIST:type_name -> B
// 4: A.ONEOF_B:type_name -> B
// 5: A.LIST_ENUM:type_name -> Enumeration
// 6: A.MAPEntry.value:type_name -> B
// [7:7] is the sub-list for method output_type
// [7:7] is the sub-list for method input_type
// [7:7] is the sub-list for extension type_name
// [7:7] is the sub-list for extension extendee
// [0:7] is the sub-list for field type_name

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

func init() { file_testpb_1_proto_init() }

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
	0x22, 0xa5, 0x05, 0x0a, 0x01, 0x41, 0x12, 0x20, 0x0a, 0x04, 0x65, 0x6e, 0x75, 0x6d, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x0c, 0x2e, 0x45, 0x6e, 0x75, 0x6d, 0x65, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x04, 0x65, 0x6e, 0x75, 0x6d, 0x12, 0x21, 0x0a, 0x0c, 0x73, 0x6f, 0x6d, 0x65,
	0x5f, 0x62, 0x6f, 0x6f, 0x6c, 0x65, 0x61, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0b,
	0x73, 0x6f, 0x6d, 0x65, 0x42, 0x6f, 0x6f, 0x6c, 0x65, 0x61, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x49,
	0x4e, 0x54, 0x33, 0x32, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x49, 0x4e, 0x54, 0x33,
	0x32, 0x12, 0x16, 0x0a, 0x06, 0x53, 0x49, 0x4e, 0x54, 0x33, 0x32, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x11, 0x52, 0x06, 0x53, 0x49, 0x4e, 0x54, 0x33, 0x32, 0x12, 0x16, 0x0a, 0x06, 0x55, 0x49, 0x4e,
	0x54, 0x33, 0x32, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x55, 0x49, 0x4e, 0x54, 0x33,
	0x32, 0x12, 0x14, 0x0a, 0x05, 0x49, 0x4e, 0x54, 0x36, 0x34, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x05, 0x49, 0x4e, 0x54, 0x36, 0x34, 0x12, 0x16, 0x0a, 0x06, 0x53, 0x49, 0x4e, 0x47, 0x36,
	0x34, 0x18, 0x07, 0x20, 0x01, 0x28, 0x12, 0x52, 0x06, 0x53, 0x49, 0x4e, 0x47, 0x36, 0x34, 0x12,
	0x16, 0x0a, 0x06, 0x55, 0x49, 0x4e, 0x54, 0x36, 0x34, 0x18, 0x08, 0x20, 0x01, 0x28, 0x04, 0x52,
	0x06, 0x55, 0x49, 0x4e, 0x54, 0x36, 0x34, 0x12, 0x1a, 0x0a, 0x08, 0x53, 0x46, 0x49, 0x58, 0x45,
	0x44, 0x33, 0x32, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0f, 0x52, 0x08, 0x53, 0x46, 0x49, 0x58, 0x45,
	0x44, 0x33, 0x32, 0x12, 0x18, 0x0a, 0x07, 0x46, 0x49, 0x58, 0x45, 0x44, 0x33, 0x32, 0x18, 0x0a,
	0x20, 0x01, 0x28, 0x07, 0x52, 0x07, 0x46, 0x49, 0x58, 0x45, 0x44, 0x33, 0x32, 0x12, 0x14, 0x0a,
	0x05, 0x46, 0x4c, 0x4f, 0x41, 0x54, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x02, 0x52, 0x05, 0x46, 0x4c,
	0x4f, 0x41, 0x54, 0x12, 0x1a, 0x0a, 0x08, 0x53, 0x46, 0x49, 0x58, 0x45, 0x44, 0x36, 0x34, 0x18,
	0x0c, 0x20, 0x01, 0x28, 0x10, 0x52, 0x08, 0x53, 0x46, 0x49, 0x58, 0x45, 0x44, 0x36, 0x34, 0x12,
	0x18, 0x0a, 0x07, 0x46, 0x49, 0x58, 0x45, 0x44, 0x36, 0x34, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x06,
	0x52, 0x07, 0x46, 0x49, 0x58, 0x45, 0x44, 0x36, 0x34, 0x12, 0x16, 0x0a, 0x06, 0x44, 0x4f, 0x55,
	0x42, 0x4c, 0x45, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x01, 0x52, 0x06, 0x44, 0x4f, 0x55, 0x42, 0x4c,
	0x45, 0x12, 0x16, 0x0a, 0x06, 0x53, 0x54, 0x52, 0x49, 0x4e, 0x47, 0x18, 0x0f, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x53, 0x54, 0x52, 0x49, 0x4e, 0x47, 0x12, 0x14, 0x0a, 0x05, 0x42, 0x59, 0x54,
	0x45, 0x53, 0x18, 0x10, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x05, 0x42, 0x59, 0x54, 0x45, 0x53, 0x12,
	0x1c, 0x0a, 0x07, 0x4d, 0x45, 0x53, 0x53, 0x41, 0x47, 0x45, 0x18, 0x11, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x02, 0x2e, 0x42, 0x52, 0x07, 0x4d, 0x45, 0x53, 0x53, 0x41, 0x47, 0x45, 0x12, 0x1d, 0x0a,
	0x03, 0x4d, 0x41, 0x50, 0x18, 0x12, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x41, 0x2e, 0x4d,
	0x41, 0x50, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x03, 0x4d, 0x41, 0x50, 0x12, 0x16, 0x0a, 0x04,
	0x4c, 0x49, 0x53, 0x54, 0x18, 0x13, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x02, 0x2e, 0x42, 0x52, 0x04,
	0x4c, 0x49, 0x53, 0x54, 0x12, 0x1d, 0x0a, 0x07, 0x4f, 0x4e, 0x45, 0x4f, 0x46, 0x5f, 0x42, 0x18,
	0x14, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x02, 0x2e, 0x42, 0x48, 0x00, 0x52, 0x06, 0x4f, 0x4e, 0x45,
	0x4f, 0x46, 0x42, 0x12, 0x23, 0x0a, 0x0c, 0x4f, 0x4e, 0x45, 0x4f, 0x46, 0x5f, 0x53, 0x54, 0x52,
	0x49, 0x4e, 0x47, 0x18, 0x15, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x0b, 0x4f, 0x4e, 0x45,
	0x4f, 0x46, 0x53, 0x54, 0x52, 0x49, 0x4e, 0x47, 0x12, 0x29, 0x0a, 0x09, 0x4c, 0x49, 0x53, 0x54,
	0x5f, 0x45, 0x4e, 0x55, 0x4d, 0x18, 0x16, 0x20, 0x03, 0x28, 0x0e, 0x32, 0x0c, 0x2e, 0x45, 0x6e,
	0x75, 0x6d, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x08, 0x4c, 0x49, 0x53, 0x54, 0x45,
	0x4e, 0x55, 0x4d, 0x1a, 0x3a, 0x0a, 0x08, 0x4d, 0x41, 0x50, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12,
	0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65,
	0x79, 0x12, 0x18, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x02, 0x2e, 0x42, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x42,
	0x07, 0x0a, 0x05, 0x4f, 0x4e, 0x45, 0x4f, 0x46, 0x22, 0x11, 0x0a, 0x01, 0x42, 0x12, 0x0c, 0x0a,
	0x01, 0x78, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x01, 0x78, 0x2a, 0x1f, 0x0a, 0x0b, 0x45,
	0x6e, 0x75, 0x6d, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x07, 0x0a, 0x03, 0x4f, 0x6e,
	0x65, 0x10, 0x00, 0x12, 0x07, 0x0a, 0x03, 0x54, 0x77, 0x6f, 0x10, 0x01, 0x42, 0x27, 0x5a, 0x25,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x73, 0x6d, 0x6f,
	0x73, 0x2f, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x74,
	0x65, 0x73, 0x74, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
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
	(Enumeration)(0),
	(*A)(nil),
	(*B)(nil),
	nil,
}
var file_testpb_1_proto_depIdxs = []int32{
	0,
	2,
	3,
	2,
	2,
	0,
	2,
	7,
	7,
	7,
	7,
	0,
}
