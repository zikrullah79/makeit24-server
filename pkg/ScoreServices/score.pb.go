// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.17.3
// source: pkg/ScoreServices/score.proto

package ScoreServices

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

type GetAllScoreRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetAllScoreRequest) Reset() {
	*x = GetAllScoreRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_ScoreServices_score_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllScoreRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllScoreRequest) ProtoMessage() {}

func (x *GetAllScoreRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_ScoreServices_score_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllScoreRequest.ProtoReflect.Descriptor instead.
func (*GetAllScoreRequest) Descriptor() ([]byte, []int) {
	return file_pkg_ScoreServices_score_proto_rawDescGZIP(), []int{0}
}

type Score struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name  string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Point string `protobuf:"bytes,2,opt,name=point,proto3" json:"point,omitempty"`
}

func (x *Score) Reset() {
	*x = Score{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_ScoreServices_score_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Score) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Score) ProtoMessage() {}

func (x *Score) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_ScoreServices_score_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Score.ProtoReflect.Descriptor instead.
func (*Score) Descriptor() ([]byte, []int) {
	return file_pkg_ScoreServices_score_proto_rawDescGZIP(), []int{1}
}

func (x *Score) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Score) GetPoint() string {
	if x != nil {
		return x.Point
	}
	return ""
}

type AddNewScoreRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name  string  `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Score float32 `protobuf:"fixed32,2,opt,name=score,proto3" json:"score,omitempty"`
}

func (x *AddNewScoreRequest) Reset() {
	*x = AddNewScoreRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_ScoreServices_score_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddNewScoreRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddNewScoreRequest) ProtoMessage() {}

func (x *AddNewScoreRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_ScoreServices_score_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddNewScoreRequest.ProtoReflect.Descriptor instead.
func (*AddNewScoreRequest) Descriptor() ([]byte, []int) {
	return file_pkg_ScoreServices_score_proto_rawDescGZIP(), []int{2}
}

func (x *AddNewScoreRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *AddNewScoreRequest) GetScore() float32 {
	if x != nil {
		return x.Score
	}
	return 0
}

type ScoreResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *ScoreResponse) Reset() {
	*x = ScoreResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_ScoreServices_score_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ScoreResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ScoreResponse) ProtoMessage() {}

func (x *ScoreResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_ScoreServices_score_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ScoreResponse.ProtoReflect.Descriptor instead.
func (*ScoreResponse) Descriptor() ([]byte, []int) {
	return file_pkg_ScoreServices_score_proto_rawDescGZIP(), []int{3}
}

func (x *ScoreResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_pkg_ScoreServices_score_proto protoreflect.FileDescriptor

var file_pkg_ScoreServices_score_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x70, 0x6b, 0x67, 0x2f, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x73, 0x2f, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x0d, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x22, 0x14,
	0x0a, 0x12, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x22, 0x31, 0x0a, 0x05, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x12, 0x12, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x22, 0x3e, 0x0a, 0x12, 0x41, 0x64, 0x64, 0x4e, 0x65,
	0x77, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x02,
	0x52, 0x05, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x22, 0x29, 0x0a, 0x0d, 0x53, 0x63, 0x6f, 0x72, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x32, 0xad, 0x01, 0x0a, 0x0d, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x73, 0x12, 0x4a, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x53, 0x63,
	0x6f, 0x72, 0x65, 0x12, 0x21, 0x2e, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x22, 0x00, 0x30, 0x01,
	0x12, 0x50, 0x0a, 0x0b, 0x41, 0x64, 0x64, 0x4e, 0x65, 0x77, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x12,
	0x21, 0x2e, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e,
	0x41, 0x64, 0x64, 0x4e, 0x65, 0x77, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x73, 0x2e, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x42, 0x15, 0x5a, 0x13, 0x2e, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x53, 0x63, 0x6f, 0x72,
	0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_pkg_ScoreServices_score_proto_rawDescOnce sync.Once
	file_pkg_ScoreServices_score_proto_rawDescData = file_pkg_ScoreServices_score_proto_rawDesc
)

func file_pkg_ScoreServices_score_proto_rawDescGZIP() []byte {
	file_pkg_ScoreServices_score_proto_rawDescOnce.Do(func() {
		file_pkg_ScoreServices_score_proto_rawDescData = protoimpl.X.CompressGZIP(file_pkg_ScoreServices_score_proto_rawDescData)
	})
	return file_pkg_ScoreServices_score_proto_rawDescData
}

var file_pkg_ScoreServices_score_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_pkg_ScoreServices_score_proto_goTypes = []interface{}{
	(*GetAllScoreRequest)(nil), // 0: ScoreServices.GetAllScoreRequest
	(*Score)(nil),              // 1: ScoreServices.Score
	(*AddNewScoreRequest)(nil), // 2: ScoreServices.AddNewScoreRequest
	(*ScoreResponse)(nil),      // 3: ScoreServices.ScoreResponse
}
var file_pkg_ScoreServices_score_proto_depIdxs = []int32{
	0, // 0: ScoreServices.ScoreServices.GetAllScore:input_type -> ScoreServices.GetAllScoreRequest
	2, // 1: ScoreServices.ScoreServices.AddNewScore:input_type -> ScoreServices.AddNewScoreRequest
	1, // 2: ScoreServices.ScoreServices.GetAllScore:output_type -> ScoreServices.Score
	3, // 3: ScoreServices.ScoreServices.AddNewScore:output_type -> ScoreServices.ScoreResponse
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_pkg_ScoreServices_score_proto_init() }
func file_pkg_ScoreServices_score_proto_init() {
	if File_pkg_ScoreServices_score_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pkg_ScoreServices_score_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAllScoreRequest); i {
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
		file_pkg_ScoreServices_score_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Score); i {
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
		file_pkg_ScoreServices_score_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddNewScoreRequest); i {
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
		file_pkg_ScoreServices_score_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ScoreResponse); i {
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
			RawDescriptor: file_pkg_ScoreServices_score_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_pkg_ScoreServices_score_proto_goTypes,
		DependencyIndexes: file_pkg_ScoreServices_score_proto_depIdxs,
		MessageInfos:      file_pkg_ScoreServices_score_proto_msgTypes,
	}.Build()
	File_pkg_ScoreServices_score_proto = out.File
	file_pkg_ScoreServices_score_proto_rawDesc = nil
	file_pkg_ScoreServices_score_proto_goTypes = nil
	file_pkg_ScoreServices_score_proto_depIdxs = nil
}