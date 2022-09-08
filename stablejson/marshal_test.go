package stablejson_test

import (
	"testing"
	"time"

	"google.golang.org/protobuf/types/known/anypb"
	"google.golang.org/protobuf/types/known/durationpb"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/fieldmaskpb"
	"google.golang.org/protobuf/types/known/structpb"
	"google.golang.org/protobuf/types/known/timestamppb"
	"google.golang.org/protobuf/types/known/wrapperspb"
	"gotest.tools/v3/assert"

	"github.com/cosmos/cosmos-proto/stablejson"
	"github.com/cosmos/cosmos-proto/stablejson/internal/testpb"
)

func TestStableJSON(t *testing.T) {
	a, err := anypb.New(&testpb.ABitOfEverything{
		I32: 10,
		Str: "abc",
	})
	assert.NilError(t, err)
	msg := &testpb.ABitOfEverything{
		Message: &testpb.NestedMessage{
			Foo: "test",
			Bar: 0, // this is the default value and should be omitted from output
		},
		Enum: testpb.AnEnum_ONE,
		StrMap: map[string]string{
			"foo": "abc",
			"bar": "def",
		},
		Int32Map: map[int32]string{
			-3: "xyz",
			0:  "abc",
			10: "qrs",
		},
		BoolMap: map[bool]string{
			true:  "T",
			false: "F",
		},
		Repeated:  []int32{3, -7, 2, 6, 4},
		Str:       `abcxyz"foo"def`,
		Bool:      true,
		Bytes:     []byte{0, 1, 2, 3},
		I32:       -15,
		F32:       1001,
		U32:       1200,
		Si32:      -376,
		Sf32:      -1000,
		I64:       14578294827584932,
		F64:       9572348124213523654,
		U64:       4759492485,
		Si64:      -59268425823934,
		Sf64:      -659101379604211154,
		Float:     1.0,
		Double:    5235.2941,
		Any:       a,
		Timestamp: timestamppb.New(time.Date(2022, 1, 1, 12, 31, 0, 0, time.UTC)),
		Duration:  durationpb.New(time.Second * 3000),
		Struct: &structpb.Struct{
			Fields: map[string]*structpb.Value{
				"null": structpb.NewNullValue(),
				"num":  structpb.NewNumberValue(3.76),
				"str":  structpb.NewStringValue("abc"),
				"bool": structpb.NewBoolValue(true),
				"nested struct": structpb.NewStructValue(&structpb.Struct{Fields: map[string]*structpb.Value{
					"a": structpb.NewStringValue("abc"),
				}}),
				"struct list": structpb.NewListValue(&structpb.ListValue{Values: []*structpb.Value{
					structpb.NewStringValue("xyz"),
					structpb.NewBoolValue(false),
					structpb.NewNumberValue(-9),
				}}),
			},
		},
		BoolValue:   &wrapperspb.BoolValue{Value: true},
		BytesValue:  &wrapperspb.BytesValue{Value: []byte{0, 1, 2, 3}},
		DoubleValue: &wrapperspb.DoubleValue{Value: 1.324},
		FloatValue:  &wrapperspb.FloatValue{Value: -1.0},
		Int32Value:  &wrapperspb.Int32Value{Value: 10},
		Int64Value:  &wrapperspb.Int64Value{Value: -376923457},
		StringValue: &wrapperspb.StringValue{Value: "gfedcba"},
		Uint32Value: &wrapperspb.UInt32Value{Value: 37492},
		Uint64Value: &wrapperspb.UInt64Value{Value: 1892409137358391},
		FieldMask:   &fieldmaskpb.FieldMask{Paths: []string{"a.b", "a.c", "b"}},
		ListValue: &structpb.ListValue{Values: []*structpb.Value{
			structpb.NewNumberValue(1.1),
			structpb.NewStringValue("qrs"),
		}},
		Value:     &structpb.Value{},
		NullValue: structpb.NullValue_NULL_VALUE,
		Empty:     &emptypb.Empty{},
	}
	bz, err := stablejson.Marshal(msg)
	assert.NilError(t, err)
	assert.Equal(t,
		`{"message":{"foo":"test"},"enum":"ONE","str_map":,"int32_map":,"bool_map":,"repeated":[3,-7,2,6,4],"str":"abcxyz\"foo\"def","bool":true,"bytes":"AAECAw==","i32":-15,"f32":1001,"u32":1200,"si32":-376,"sf32":-1000,"i64":"14578294827584932","f64":"9572348124213523654","u64":"4759492485","si64":"-59268425823934","sf64":"-659101379604211154","float":1,"double":5235.2941,"any":{"@type_url":"type.googleapis.com/testpb.ABitOfEverything","str":"abc","i32":10},"timestamp":,"duration":,"struct":{"bool":true,"nested struct":{"a":"abc"},"null":null,"num":3.76,"str":"abc","struct list":["xyz",false,-9]},"bool_value":true,"bytes_value":"AAECAw==","double_value":1.324,"float_value":-1,"int32_value":10,"int64_value":"-376923457","string_value":"gfedcba","uint32_value":37492,"uint64_value":"1892409137358391","field_mask":"a.b,a.c,b","list_value":[1.1,"qrs"],"value":,"empty":{}}`,
		string(bz))
}
