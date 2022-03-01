// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.19.3
// source: api/draws/v0/service.proto

package v0

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

type GetDrawRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uuid int64 `protobuf:"varint,1,opt,name=uuid,proto3" json:"uuid,omitempty"`
}

func (x *GetDrawRequest) Reset() {
	*x = GetDrawRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_draws_v0_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetDrawRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetDrawRequest) ProtoMessage() {}

func (x *GetDrawRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_draws_v0_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetDrawRequest.ProtoReflect.Descriptor instead.
func (*GetDrawRequest) Descriptor() ([]byte, []int) {
	return file_api_draws_v0_service_proto_rawDescGZIP(), []int{0}
}

func (x *GetDrawRequest) GetUuid() int64 {
	if x != nil {
		return x.Uuid
	}
	return 0
}

type GetDrawResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Draw *Draw `protobuf:"bytes,1,opt,name=draw,proto3" json:"draw,omitempty"`
}

func (x *GetDrawResponse) Reset() {
	*x = GetDrawResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_draws_v0_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetDrawResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetDrawResponse) ProtoMessage() {}

func (x *GetDrawResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_draws_v0_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetDrawResponse.ProtoReflect.Descriptor instead.
func (*GetDrawResponse) Descriptor() ([]byte, []int) {
	return file_api_draws_v0_service_proto_rawDescGZIP(), []int{1}
}

func (x *GetDrawResponse) GetDraw() *Draw {
	if x != nil {
		return x.Draw
	}
	return nil
}

type AddDrawRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Draw *Draw `protobuf:"bytes,1,opt,name=draw,proto3" json:"draw,omitempty"`
}

func (x *AddDrawRequest) Reset() {
	*x = AddDrawRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_draws_v0_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddDrawRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddDrawRequest) ProtoMessage() {}

func (x *AddDrawRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_draws_v0_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddDrawRequest.ProtoReflect.Descriptor instead.
func (*AddDrawRequest) Descriptor() ([]byte, []int) {
	return file_api_draws_v0_service_proto_rawDescGZIP(), []int{2}
}

func (x *AddDrawRequest) GetDraw() *Draw {
	if x != nil {
		return x.Draw
	}
	return nil
}

type AddDrawResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Draw *Draw `protobuf:"bytes,1,opt,name=draw,proto3" json:"draw,omitempty"`
}

func (x *AddDrawResponse) Reset() {
	*x = AddDrawResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_draws_v0_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddDrawResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddDrawResponse) ProtoMessage() {}

func (x *AddDrawResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_draws_v0_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddDrawResponse.ProtoReflect.Descriptor instead.
func (*AddDrawResponse) Descriptor() ([]byte, []int) {
	return file_api_draws_v0_service_proto_rawDescGZIP(), []int{3}
}

func (x *AddDrawResponse) GetDraw() *Draw {
	if x != nil {
		return x.Draw
	}
	return nil
}

type ListDrawRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PageLimit int64 `protobuf:"varint,1,opt,name=pageLimit,proto3" json:"pageLimit,omitempty"`
}

func (x *ListDrawRequest) Reset() {
	*x = ListDrawRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_draws_v0_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListDrawRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListDrawRequest) ProtoMessage() {}

func (x *ListDrawRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_draws_v0_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListDrawRequest.ProtoReflect.Descriptor instead.
func (*ListDrawRequest) Descriptor() ([]byte, []int) {
	return file_api_draws_v0_service_proto_rawDescGZIP(), []int{4}
}

func (x *ListDrawRequest) GetPageLimit() int64 {
	if x != nil {
		return x.PageLimit
	}
	return 0
}

type ListDrawResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Draws []*Draw `protobuf:"bytes,1,rep,name=draws,proto3" json:"draws,omitempty"`
}

func (x *ListDrawResponse) Reset() {
	*x = ListDrawResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_draws_v0_service_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListDrawResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListDrawResponse) ProtoMessage() {}

func (x *ListDrawResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_draws_v0_service_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListDrawResponse.ProtoReflect.Descriptor instead.
func (*ListDrawResponse) Descriptor() ([]byte, []int) {
	return file_api_draws_v0_service_proto_rawDescGZIP(), []int{5}
}

func (x *ListDrawResponse) GetDraws() []*Draw {
	if x != nil {
		return x.Draws
	}
	return nil
}

var File_api_draws_v0_service_proto protoreflect.FileDescriptor

var file_api_draws_v0_service_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x61, 0x70, 0x69, 0x2f, 0x64, 0x72, 0x61, 0x77, 0x73, 0x2f, 0x76, 0x30, 0x2f, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x10, 0x64, 0x72,
	0x61, 0x77, 0x73, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x30, 0x1a, 0x1a,
	0x61, 0x70, 0x69, 0x2f, 0x64, 0x72, 0x61, 0x77, 0x73, 0x2f, 0x76, 0x30, 0x2f, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x24, 0x0a, 0x0e, 0x47, 0x65,
	0x74, 0x44, 0x72, 0x61, 0x77, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04,
	0x75, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x75, 0x75, 0x69, 0x64,
	0x22, 0x3d, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x44, 0x72, 0x61, 0x77, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x2a, 0x0a, 0x04, 0x64, 0x72, 0x61, 0x77, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x16, 0x2e, 0x64, 0x72, 0x61, 0x77, 0x73, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x2e, 0x76, 0x30, 0x2e, 0x44, 0x72, 0x61, 0x77, 0x52, 0x04, 0x64, 0x72, 0x61, 0x77, 0x22,
	0x3c, 0x0a, 0x0e, 0x41, 0x64, 0x64, 0x44, 0x72, 0x61, 0x77, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x2a, 0x0a, 0x04, 0x64, 0x72, 0x61, 0x77, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x16, 0x2e, 0x64, 0x72, 0x61, 0x77, 0x73, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e,
	0x76, 0x30, 0x2e, 0x44, 0x72, 0x61, 0x77, 0x52, 0x04, 0x64, 0x72, 0x61, 0x77, 0x22, 0x3d, 0x0a,
	0x0f, 0x41, 0x64, 0x64, 0x44, 0x72, 0x61, 0x77, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x2a, 0x0a, 0x04, 0x64, 0x72, 0x61, 0x77, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16,
	0x2e, 0x64, 0x72, 0x61, 0x77, 0x73, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x76,
	0x30, 0x2e, 0x44, 0x72, 0x61, 0x77, 0x52, 0x04, 0x64, 0x72, 0x61, 0x77, 0x22, 0x2f, 0x0a, 0x0f,
	0x4c, 0x69, 0x73, 0x74, 0x44, 0x72, 0x61, 0x77, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x1c, 0x0a, 0x09, 0x70, 0x61, 0x67, 0x65, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x09, 0x70, 0x61, 0x67, 0x65, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x22, 0x40, 0x0a,
	0x10, 0x4c, 0x69, 0x73, 0x74, 0x44, 0x72, 0x61, 0x77, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x2c, 0x0a, 0x05, 0x64, 0x72, 0x61, 0x77, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x16, 0x2e, 0x64, 0x72, 0x61, 0x77, 0x73, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x2e, 0x76, 0x30, 0x2e, 0x44, 0x72, 0x61, 0x77, 0x52, 0x05, 0x64, 0x72, 0x61, 0x77, 0x73, 0x32,
	0x87, 0x02, 0x0a, 0x0c, 0x44, 0x72, 0x61, 0x77, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x12, 0x50, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x44, 0x72, 0x61, 0x77, 0x12, 0x20, 0x2e, 0x64, 0x72,
	0x61, 0x77, 0x73, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x30, 0x2e, 0x47,
	0x65, 0x74, 0x44, 0x72, 0x61, 0x77, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e,
	0x64, 0x72, 0x61, 0x77, 0x73, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x30,
	0x2e, 0x47, 0x65, 0x74, 0x44, 0x72, 0x61, 0x77, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x12, 0x50, 0x0a, 0x07, 0x41, 0x64, 0x64, 0x44, 0x72, 0x61, 0x77, 0x12, 0x20, 0x2e,
	0x64, 0x72, 0x61, 0x77, 0x73, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x30,
	0x2e, 0x41, 0x64, 0x64, 0x44, 0x72, 0x61, 0x77, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x21, 0x2e, 0x64, 0x72, 0x61, 0x77, 0x73, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x76, 0x30, 0x2e, 0x41, 0x64, 0x64, 0x44, 0x72, 0x61, 0x77, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x12, 0x53, 0x0a, 0x08, 0x4c, 0x69, 0x73, 0x74, 0x44, 0x72, 0x61, 0x77,
	0x12, 0x21, 0x2e, 0x64, 0x72, 0x61, 0x77, 0x73, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x76, 0x30, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x44, 0x72, 0x61, 0x77, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x22, 0x2e, 0x64, 0x72, 0x61, 0x77, 0x73, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2e, 0x76, 0x30, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x44, 0x72, 0x61, 0x77, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x0e, 0x5a, 0x0c, 0x61, 0x70, 0x69,
	0x2f, 0x64, 0x72, 0x61, 0x77, 0x73, 0x2f, 0x76, 0x30, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_api_draws_v0_service_proto_rawDescOnce sync.Once
	file_api_draws_v0_service_proto_rawDescData = file_api_draws_v0_service_proto_rawDesc
)

func file_api_draws_v0_service_proto_rawDescGZIP() []byte {
	file_api_draws_v0_service_proto_rawDescOnce.Do(func() {
		file_api_draws_v0_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_draws_v0_service_proto_rawDescData)
	})
	return file_api_draws_v0_service_proto_rawDescData
}

var file_api_draws_v0_service_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_api_draws_v0_service_proto_goTypes = []interface{}{
	(*GetDrawRequest)(nil),   // 0: draws.service.v0.GetDrawRequest
	(*GetDrawResponse)(nil),  // 1: draws.service.v0.GetDrawResponse
	(*AddDrawRequest)(nil),   // 2: draws.service.v0.AddDrawRequest
	(*AddDrawResponse)(nil),  // 3: draws.service.v0.AddDrawResponse
	(*ListDrawRequest)(nil),  // 4: draws.service.v0.ListDrawRequest
	(*ListDrawResponse)(nil), // 5: draws.service.v0.ListDrawResponse
	(*Draw)(nil),             // 6: draws.message.v0.Draw
}
var file_api_draws_v0_service_proto_depIdxs = []int32{
	6, // 0: draws.service.v0.GetDrawResponse.draw:type_name -> draws.message.v0.Draw
	6, // 1: draws.service.v0.AddDrawRequest.draw:type_name -> draws.message.v0.Draw
	6, // 2: draws.service.v0.AddDrawResponse.draw:type_name -> draws.message.v0.Draw
	6, // 3: draws.service.v0.ListDrawResponse.draws:type_name -> draws.message.v0.Draw
	0, // 4: draws.service.v0.DrawsService.GetDraw:input_type -> draws.service.v0.GetDrawRequest
	2, // 5: draws.service.v0.DrawsService.AddDraw:input_type -> draws.service.v0.AddDrawRequest
	4, // 6: draws.service.v0.DrawsService.ListDraw:input_type -> draws.service.v0.ListDrawRequest
	1, // 7: draws.service.v0.DrawsService.GetDraw:output_type -> draws.service.v0.GetDrawResponse
	3, // 8: draws.service.v0.DrawsService.AddDraw:output_type -> draws.service.v0.AddDrawResponse
	5, // 9: draws.service.v0.DrawsService.ListDraw:output_type -> draws.service.v0.ListDrawResponse
	7, // [7:10] is the sub-list for method output_type
	4, // [4:7] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_api_draws_v0_service_proto_init() }
func file_api_draws_v0_service_proto_init() {
	if File_api_draws_v0_service_proto != nil {
		return
	}
	file_api_draws_v0_message_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_api_draws_v0_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetDrawRequest); i {
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
		file_api_draws_v0_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetDrawResponse); i {
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
		file_api_draws_v0_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddDrawRequest); i {
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
		file_api_draws_v0_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddDrawResponse); i {
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
		file_api_draws_v0_service_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListDrawRequest); i {
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
		file_api_draws_v0_service_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListDrawResponse); i {
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
			RawDescriptor: file_api_draws_v0_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_draws_v0_service_proto_goTypes,
		DependencyIndexes: file_api_draws_v0_service_proto_depIdxs,
		MessageInfos:      file_api_draws_v0_service_proto_msgTypes,
	}.Build()
	File_api_draws_v0_service_proto = out.File
	file_api_draws_v0_service_proto_rawDesc = nil
	file_api_draws_v0_service_proto_goTypes = nil
	file_api_draws_v0_service_proto_depIdxs = nil
}