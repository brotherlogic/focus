// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.12.4
// source: focus.proto

package proto

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

type Focus_FocusType int32

const (
	Focus_UNKNOWN                  Focus_FocusType = 0
	Focus_FOCUS_ON_RECORD_CLEANING Focus_FocusType = 1
	Focus_FOCUS_ON_HOME_TASKS      Focus_FocusType = 2
)

// Enum value maps for Focus_FocusType.
var (
	Focus_FocusType_name = map[int32]string{
		0: "UNKNOWN",
		1: "FOCUS_ON_RECORD_CLEANING",
		2: "FOCUS_ON_HOME_TASKS",
	}
	Focus_FocusType_value = map[string]int32{
		"UNKNOWN":                  0,
		"FOCUS_ON_RECORD_CLEANING": 1,
		"FOCUS_ON_HOME_TASKS":      2,
	}
)

func (x Focus_FocusType) Enum() *Focus_FocusType {
	p := new(Focus_FocusType)
	*p = x
	return p
}

func (x Focus_FocusType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Focus_FocusType) Descriptor() protoreflect.EnumDescriptor {
	return file_focus_proto_enumTypes[0].Descriptor()
}

func (Focus_FocusType) Type() protoreflect.EnumType {
	return &file_focus_proto_enumTypes[0]
}

func (x Focus_FocusType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Focus_FocusType.Descriptor instead.
func (Focus_FocusType) EnumDescriptor() ([]byte, []int) {
	return file_focus_proto_rawDescGZIP(), []int{1, 0}
}

type Config struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Date       string           `protobuf:"bytes,1,opt,name=date,proto3" json:"date,omitempty"`
	IssueCount map[string]int32 `protobuf:"bytes,2,rep,name=issue_count,json=issueCount,proto3" json:"issue_count,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3"`
}

func (x *Config) Reset() {
	*x = Config{}
	if protoimpl.UnsafeEnabled {
		mi := &file_focus_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Config) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Config) ProtoMessage() {}

func (x *Config) ProtoReflect() protoreflect.Message {
	mi := &file_focus_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Config.ProtoReflect.Descriptor instead.
func (*Config) Descriptor() ([]byte, []int) {
	return file_focus_proto_rawDescGZIP(), []int{0}
}

func (x *Config) GetDate() string {
	if x != nil {
		return x.Date
	}
	return ""
}

func (x *Config) GetIssueCount() map[string]int32 {
	if x != nil {
		return x.IssueCount
	}
	return nil
}

type Focus struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type   Focus_FocusType `protobuf:"varint,1,opt,name=type,proto3,enum=focus.Focus_FocusType" json:"type,omitempty"`
	Detail string          `protobuf:"bytes,2,opt,name=detail,proto3" json:"detail,omitempty"`
	Link   string          `protobuf:"bytes,3,opt,name=link,proto3" json:"link,omitempty"`
}

func (x *Focus) Reset() {
	*x = Focus{}
	if protoimpl.UnsafeEnabled {
		mi := &file_focus_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Focus) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Focus) ProtoMessage() {}

func (x *Focus) ProtoReflect() protoreflect.Message {
	mi := &file_focus_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Focus.ProtoReflect.Descriptor instead.
func (*Focus) Descriptor() ([]byte, []int) {
	return file_focus_proto_rawDescGZIP(), []int{1}
}

func (x *Focus) GetType() Focus_FocusType {
	if x != nil {
		return x.Type
	}
	return Focus_UNKNOWN
}

func (x *Focus) GetDetail() string {
	if x != nil {
		return x.Detail
	}
	return ""
}

func (x *Focus) GetLink() string {
	if x != nil {
		return x.Link
	}
	return ""
}

type GetFocusRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetFocusRequest) Reset() {
	*x = GetFocusRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_focus_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetFocusRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetFocusRequest) ProtoMessage() {}

func (x *GetFocusRequest) ProtoReflect() protoreflect.Message {
	mi := &file_focus_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetFocusRequest.ProtoReflect.Descriptor instead.
func (*GetFocusRequest) Descriptor() ([]byte, []int) {
	return file_focus_proto_rawDescGZIP(), []int{2}
}

type GetFocusResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Focus *Focus `protobuf:"bytes,1,opt,name=focus,proto3" json:"focus,omitempty"`
}

func (x *GetFocusResponse) Reset() {
	*x = GetFocusResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_focus_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetFocusResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetFocusResponse) ProtoMessage() {}

func (x *GetFocusResponse) ProtoReflect() protoreflect.Message {
	mi := &file_focus_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetFocusResponse.ProtoReflect.Descriptor instead.
func (*GetFocusResponse) Descriptor() ([]byte, []int) {
	return file_focus_proto_rawDescGZIP(), []int{3}
}

func (x *GetFocusResponse) GetFocus() *Focus {
	if x != nil {
		return x.Focus
	}
	return nil
}

var File_focus_proto protoreflect.FileDescriptor

var file_focus_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x66, 0x6f, 0x63, 0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x66,
	0x6f, 0x63, 0x75, 0x73, 0x22, 0x9b, 0x01, 0x0a, 0x06, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12,
	0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x64,
	0x61, 0x74, 0x65, 0x12, 0x3e, 0x0a, 0x0b, 0x69, 0x73, 0x73, 0x75, 0x65, 0x5f, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x66, 0x6f, 0x63, 0x75, 0x73,
	0x2e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x49, 0x73, 0x73, 0x75, 0x65, 0x43, 0x6f, 0x75,
	0x6e, 0x74, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0a, 0x69, 0x73, 0x73, 0x75, 0x65, 0x43, 0x6f,
	0x75, 0x6e, 0x74, 0x1a, 0x3d, 0x0a, 0x0f, 0x49, 0x73, 0x73, 0x75, 0x65, 0x43, 0x6f, 0x75, 0x6e,
	0x74, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02,
	0x38, 0x01, 0x22, 0xb0, 0x01, 0x0a, 0x05, 0x46, 0x6f, 0x63, 0x75, 0x73, 0x12, 0x2a, 0x0a, 0x04,
	0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x16, 0x2e, 0x66, 0x6f, 0x63,
	0x75, 0x73, 0x2e, 0x46, 0x6f, 0x63, 0x75, 0x73, 0x2e, 0x46, 0x6f, 0x63, 0x75, 0x73, 0x54, 0x79,
	0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x64, 0x65, 0x74, 0x61,
	0x69, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c,
	0x12, 0x12, 0x0a, 0x04, 0x6c, 0x69, 0x6e, 0x6b, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6c, 0x69, 0x6e, 0x6b, 0x22, 0x4f, 0x0a, 0x09, 0x46, 0x6f, 0x63, 0x75, 0x73, 0x54, 0x79, 0x70,
	0x65, 0x12, 0x0b, 0x0a, 0x07, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x00, 0x12, 0x1c,
	0x0a, 0x18, 0x46, 0x4f, 0x43, 0x55, 0x53, 0x5f, 0x4f, 0x4e, 0x5f, 0x52, 0x45, 0x43, 0x4f, 0x52,
	0x44, 0x5f, 0x43, 0x4c, 0x45, 0x41, 0x4e, 0x49, 0x4e, 0x47, 0x10, 0x01, 0x12, 0x17, 0x0a, 0x13,
	0x46, 0x4f, 0x43, 0x55, 0x53, 0x5f, 0x4f, 0x4e, 0x5f, 0x48, 0x4f, 0x4d, 0x45, 0x5f, 0x54, 0x41,
	0x53, 0x4b, 0x53, 0x10, 0x02, 0x22, 0x11, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x46, 0x6f, 0x63, 0x75,
	0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x36, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x46,
	0x6f, 0x63, 0x75, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x22, 0x0a, 0x05,
	0x66, 0x6f, 0x63, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x66, 0x6f,
	0x63, 0x75, 0x73, 0x2e, 0x46, 0x6f, 0x63, 0x75, 0x73, 0x52, 0x05, 0x66, 0x6f, 0x63, 0x75, 0x73,
	0x32, 0x4b, 0x0a, 0x0c, 0x46, 0x6f, 0x63, 0x75, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x12, 0x3b, 0x0a, 0x08, 0x47, 0x65, 0x74, 0x46, 0x6f, 0x63, 0x75, 0x73, 0x12, 0x16, 0x2e, 0x66,
	0x6f, 0x63, 0x75, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x46, 0x6f, 0x63, 0x75, 0x73, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x66, 0x6f, 0x63, 0x75, 0x73, 0x2e, 0x47, 0x65, 0x74,
	0x46, 0x6f, 0x63, 0x75, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x25, 0x5a,
	0x23, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x62, 0x72, 0x6f, 0x74,
	0x68, 0x65, 0x72, 0x6c, 0x6f, 0x67, 0x69, 0x63, 0x2f, 0x66, 0x6f, 0x63, 0x75, 0x73, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_focus_proto_rawDescOnce sync.Once
	file_focus_proto_rawDescData = file_focus_proto_rawDesc
)

func file_focus_proto_rawDescGZIP() []byte {
	file_focus_proto_rawDescOnce.Do(func() {
		file_focus_proto_rawDescData = protoimpl.X.CompressGZIP(file_focus_proto_rawDescData)
	})
	return file_focus_proto_rawDescData
}

var file_focus_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_focus_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_focus_proto_goTypes = []interface{}{
	(Focus_FocusType)(0),     // 0: focus.Focus.FocusType
	(*Config)(nil),           // 1: focus.Config
	(*Focus)(nil),            // 2: focus.Focus
	(*GetFocusRequest)(nil),  // 3: focus.GetFocusRequest
	(*GetFocusResponse)(nil), // 4: focus.GetFocusResponse
	nil,                      // 5: focus.Config.IssueCountEntry
}
var file_focus_proto_depIdxs = []int32{
	5, // 0: focus.Config.issue_count:type_name -> focus.Config.IssueCountEntry
	0, // 1: focus.Focus.type:type_name -> focus.Focus.FocusType
	2, // 2: focus.GetFocusResponse.focus:type_name -> focus.Focus
	3, // 3: focus.FocusService.GetFocus:input_type -> focus.GetFocusRequest
	4, // 4: focus.FocusService.GetFocus:output_type -> focus.GetFocusResponse
	4, // [4:5] is the sub-list for method output_type
	3, // [3:4] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_focus_proto_init() }
func file_focus_proto_init() {
	if File_focus_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_focus_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Config); i {
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
		file_focus_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Focus); i {
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
		file_focus_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetFocusRequest); i {
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
		file_focus_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetFocusResponse); i {
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
			RawDescriptor: file_focus_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_focus_proto_goTypes,
		DependencyIndexes: file_focus_proto_depIdxs,
		EnumInfos:         file_focus_proto_enumTypes,
		MessageInfos:      file_focus_proto_msgTypes,
	}.Build()
	File_focus_proto = out.File
	file_focus_proto_rawDesc = nil
	file_focus_proto_goTypes = nil
	file_focus_proto_depIdxs = nil
}
