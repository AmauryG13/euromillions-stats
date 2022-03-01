// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.19.3
// source: api/draws/v0/message.proto

package v0

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Draw_Day int32

const (
	Draw_THURSDAY Draw_Day = 0
	Draw_FRIDAY   Draw_Day = 1
)

// Enum value maps for Draw_Day.
var (
	Draw_Day_name = map[int32]string{
		0: "THURSDAY",
		1: "FRIDAY",
	}
	Draw_Day_value = map[string]int32{
		"THURSDAY": 0,
		"FRIDAY":   1,
	}
)

func (x Draw_Day) Enum() *Draw_Day {
	p := new(Draw_Day)
	*p = x
	return p
}

func (x Draw_Day) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Draw_Day) Descriptor() protoreflect.EnumDescriptor {
	return file_api_draws_v0_message_proto_enumTypes[0].Descriptor()
}

func (Draw_Day) Type() protoreflect.EnumType {
	return &file_api_draws_v0_message_proto_enumTypes[0]
}

func (x Draw_Day) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Draw_Day.Descriptor instead.
func (Draw_Day) EnumDescriptor() ([]byte, []int) {
	return file_api_draws_v0_message_proto_rawDescGZIP(), []int{0, 0}
}

type Draw struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uuid       int64                  `protobuf:"varint,1,opt,name=uuid,proto3" json:"uuid,omitempty"`
	Number     int64                  `protobuf:"varint,2,opt,name=number,proto3" json:"number,omitempty"`
	Day        Draw_Day               `protobuf:"varint,3,opt,name=day,proto3,enum=draws.message.v0.Draw_Day" json:"day,omitempty"`
	Cycle      int32                  `protobuf:"varint,4,opt,name=cycle,proto3" json:"cycle,omitempty"`
	Forclusion *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=forclusion,proto3" json:"forclusion,omitempty"`
}

func (x *Draw) Reset() {
	*x = Draw{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_draws_v0_message_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Draw) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Draw) ProtoMessage() {}

func (x *Draw) ProtoReflect() protoreflect.Message {
	mi := &file_api_draws_v0_message_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Draw.ProtoReflect.Descriptor instead.
func (*Draw) Descriptor() ([]byte, []int) {
	return file_api_draws_v0_message_proto_rawDescGZIP(), []int{0}
}

func (x *Draw) GetUuid() int64 {
	if x != nil {
		return x.Uuid
	}
	return 0
}

func (x *Draw) GetNumber() int64 {
	if x != nil {
		return x.Number
	}
	return 0
}

func (x *Draw) GetDay() Draw_Day {
	if x != nil {
		return x.Day
	}
	return Draw_THURSDAY
}

func (x *Draw) GetCycle() int32 {
	if x != nil {
		return x.Cycle
	}
	return 0
}

func (x *Draw) GetForclusion() *timestamppb.Timestamp {
	if x != nil {
		return x.Forclusion
	}
	return nil
}

var File_api_draws_v0_message_proto protoreflect.FileDescriptor

var file_api_draws_v0_message_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x61, 0x70, 0x69, 0x2f, 0x64, 0x72, 0x61, 0x77, 0x73, 0x2f, 0x76, 0x30, 0x2f, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x10, 0x64, 0x72,
	0x61, 0x77, 0x73, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x76, 0x30, 0x1a, 0x1f,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f,
	0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0xd3, 0x01, 0x0a, 0x04, 0x44, 0x72, 0x61, 0x77, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x75, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x75, 0x75, 0x69, 0x64, 0x12, 0x16, 0x0a, 0x06,
	0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x6e, 0x75,
	0x6d, 0x62, 0x65, 0x72, 0x12, 0x2c, 0x0a, 0x03, 0x64, 0x61, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x1a, 0x2e, 0x64, 0x72, 0x61, 0x77, 0x73, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x2e, 0x76, 0x30, 0x2e, 0x44, 0x72, 0x61, 0x77, 0x2e, 0x44, 0x61, 0x79, 0x52, 0x03, 0x64,
	0x61, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x79, 0x63, 0x6c, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x05, 0x63, 0x79, 0x63, 0x6c, 0x65, 0x12, 0x3a, 0x0a, 0x0a, 0x66, 0x6f, 0x72, 0x63,
	0x6c, 0x75, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0a, 0x66, 0x6f, 0x72, 0x63, 0x6c, 0x75,
	0x73, 0x69, 0x6f, 0x6e, 0x22, 0x1f, 0x0a, 0x03, 0x44, 0x61, 0x79, 0x12, 0x0c, 0x0a, 0x08, 0x54,
	0x48, 0x55, 0x52, 0x53, 0x44, 0x41, 0x59, 0x10, 0x00, 0x12, 0x0a, 0x0a, 0x06, 0x46, 0x52, 0x49,
	0x44, 0x41, 0x59, 0x10, 0x01, 0x42, 0x0e, 0x5a, 0x0c, 0x61, 0x70, 0x69, 0x2f, 0x64, 0x72, 0x61,
	0x77, 0x73, 0x2f, 0x76, 0x30, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_draws_v0_message_proto_rawDescOnce sync.Once
	file_api_draws_v0_message_proto_rawDescData = file_api_draws_v0_message_proto_rawDesc
)

func file_api_draws_v0_message_proto_rawDescGZIP() []byte {
	file_api_draws_v0_message_proto_rawDescOnce.Do(func() {
		file_api_draws_v0_message_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_draws_v0_message_proto_rawDescData)
	})
	return file_api_draws_v0_message_proto_rawDescData
}

var file_api_draws_v0_message_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_api_draws_v0_message_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_api_draws_v0_message_proto_goTypes = []interface{}{
	(Draw_Day)(0),                 // 0: draws.message.v0.Draw.Day
	(*Draw)(nil),                  // 1: draws.message.v0.Draw
	(*timestamppb.Timestamp)(nil), // 2: google.protobuf.Timestamp
}
var file_api_draws_v0_message_proto_depIdxs = []int32{
	0, // 0: draws.message.v0.Draw.day:type_name -> draws.message.v0.Draw.Day
	2, // 1: draws.message.v0.Draw.forclusion:type_name -> google.protobuf.Timestamp
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_api_draws_v0_message_proto_init() }
func file_api_draws_v0_message_proto_init() {
	if File_api_draws_v0_message_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_draws_v0_message_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Draw); i {
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
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_api_draws_v0_message_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_api_draws_v0_message_proto_goTypes,
		DependencyIndexes: file_api_draws_v0_message_proto_depIdxs,
		EnumInfos:         file_api_draws_v0_message_proto_enumTypes,
		MessageInfos:      file_api_draws_v0_message_proto_msgTypes,
	}.Build()
	File_api_draws_v0_message_proto = out.File
	file_api_draws_v0_message_proto_rawDesc = nil
	file_api_draws_v0_message_proto_goTypes = nil
	file_api_draws_v0_message_proto_depIdxs = nil
}
