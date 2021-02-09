// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.14.0
// source: addpb.proto

package addpb

import (
	proto "github.com/golang/protobuf/proto"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type Adding struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FirstNumber  int64 `protobuf:"varint,1,opt,name=first_number,json=firstNumber,proto3" json:"first_number,omitempty"`
	SecondNumber int64 `protobuf:"varint,2,opt,name=second_number,json=secondNumber,proto3" json:"second_number,omitempty"`
}

func (x *Adding) Reset() {
	*x = Adding{}
	if protoimpl.UnsafeEnabled {
		mi := &file_addpb_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Adding) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Adding) ProtoMessage() {}

func (x *Adding) ProtoReflect() protoreflect.Message {
	mi := &file_addpb_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Adding.ProtoReflect.Descriptor instead.
func (*Adding) Descriptor() ([]byte, []int) {
	return file_addpb_proto_rawDescGZIP(), []int{0}
}

func (x *Adding) GetFirstNumber() int64 {
	if x != nil {
		return x.FirstNumber
	}
	return 0
}

func (x *Adding) GetSecondNumber() int64 {
	if x != nil {
		return x.SecondNumber
	}
	return 0
}

type AddRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Adding *Adding `protobuf:"bytes,1,opt,name=adding,proto3" json:"adding,omitempty"`
}

func (x *AddRequest) Reset() {
	*x = AddRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_addpb_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddRequest) ProtoMessage() {}

func (x *AddRequest) ProtoReflect() protoreflect.Message {
	mi := &file_addpb_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddRequest.ProtoReflect.Descriptor instead.
func (*AddRequest) Descriptor() ([]byte, []int) {
	return file_addpb_proto_rawDescGZIP(), []int{1}
}

func (x *AddRequest) GetAdding() *Adding {
	if x != nil {
		return x.Adding
	}
	return nil
}

type AddResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Sum int64 `protobuf:"varint,1,opt,name=sum,proto3" json:"sum,omitempty"`
}

func (x *AddResponse) Reset() {
	*x = AddResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_addpb_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddResponse) ProtoMessage() {}

func (x *AddResponse) ProtoReflect() protoreflect.Message {
	mi := &file_addpb_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddResponse.ProtoReflect.Descriptor instead.
func (*AddResponse) Descriptor() ([]byte, []int) {
	return file_addpb_proto_rawDescGZIP(), []int{2}
}

func (x *AddResponse) GetSum() int64 {
	if x != nil {
		return x.Sum
	}
	return 0
}

var File_addpb_proto protoreflect.FileDescriptor

var file_addpb_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x61, 0x64, 0x64, 0x70, 0x62, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x03, 0x61,
	0x64, 0x64, 0x22, 0x50, 0x0a, 0x06, 0x41, 0x64, 0x64, 0x69, 0x6e, 0x67, 0x12, 0x21, 0x0a, 0x0c,
	0x66, 0x69, 0x72, 0x73, 0x74, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x0b, 0x66, 0x69, 0x72, 0x73, 0x74, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12,
	0x23, 0x0a, 0x0d, 0x73, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0c, 0x73, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x4e, 0x75,
	0x6d, 0x62, 0x65, 0x72, 0x22, 0x31, 0x0a, 0x0a, 0x41, 0x64, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x23, 0x0a, 0x06, 0x61, 0x64, 0x64, 0x69, 0x6e, 0x67, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x61, 0x64, 0x64, 0x2e, 0x41, 0x64, 0x64, 0x69, 0x6e, 0x67, 0x52,
	0x06, 0x61, 0x64, 0x64, 0x69, 0x6e, 0x67, 0x22, 0x1f, 0x0a, 0x0b, 0x41, 0x64, 0x64, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x73, 0x75, 0x6d, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x03, 0x73, 0x75, 0x6d, 0x32, 0x38, 0x0a, 0x0a, 0x41, 0x64, 0x64, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x2a, 0x0a, 0x03, 0x41, 0x64, 0x64, 0x12, 0x0f, 0x2e,
	0x61, 0x64, 0x64, 0x2e, 0x41, 0x64, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x10,
	0x2e, 0x61, 0x64, 0x64, 0x2e, 0x41, 0x64, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x42, 0x2b, 0x5a, 0x29, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x53, 0x2d, 0x69, 0x67, 0x6e, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2d, 0x67, 0x6f, 0x2d, 0x63,
	0x6f, 0x75, 0x72, 0x73, 0x65, 0x2f, 0x61, 0x64, 0x64, 0x2f, 0x61, 0x64, 0x64, 0x70, 0x62, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_addpb_proto_rawDescOnce sync.Once
	file_addpb_proto_rawDescData = file_addpb_proto_rawDesc
)

func file_addpb_proto_rawDescGZIP() []byte {
	file_addpb_proto_rawDescOnce.Do(func() {
		file_addpb_proto_rawDescData = protoimpl.X.CompressGZIP(file_addpb_proto_rawDescData)
	})
	return file_addpb_proto_rawDescData
}

var file_addpb_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_addpb_proto_goTypes = []interface{}{
	(*Adding)(nil),      // 0: add.Adding
	(*AddRequest)(nil),  // 1: add.AddRequest
	(*AddResponse)(nil), // 2: add.AddResponse
}
var file_addpb_proto_depIdxs = []int32{
	0, // 0: add.AddRequest.adding:type_name -> add.Adding
	1, // 1: add.AddService.Add:input_type -> add.AddRequest
	2, // 2: add.AddService.Add:output_type -> add.AddResponse
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_addpb_proto_init() }
func file_addpb_proto_init() {
	if File_addpb_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_addpb_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Adding); i {
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
		file_addpb_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddRequest); i {
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
		file_addpb_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddResponse); i {
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
			RawDescriptor: file_addpb_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_addpb_proto_goTypes,
		DependencyIndexes: file_addpb_proto_depIdxs,
		MessageInfos:      file_addpb_proto_msgTypes,
	}.Build()
	File_addpb_proto = out.File
	file_addpb_proto_rawDesc = nil
	file_addpb_proto_goTypes = nil
	file_addpb_proto_depIdxs = nil
}
