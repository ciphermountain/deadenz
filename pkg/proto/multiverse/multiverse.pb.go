// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.1
// source: pkg/proto/multiverse/multiverse.proto

package service

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

type Status int32

const (
	// option allow_alias = true;
	Status_OK      Status = 0
	Status_Failure Status = 1
)

// Enum value maps for Status.
var (
	Status_name = map[int32]string{
		0: "OK",
		1: "Failure",
	}
	Status_value = map[string]int32{
		"OK":      0,
		"Failure": 1,
	}
)

func (x Status) Enum() *Status {
	p := new(Status)
	*p = x
	return p
}

func (x Status) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Status) Descriptor() protoreflect.EnumDescriptor {
	return file_pkg_proto_multiverse_multiverse_proto_enumTypes[0].Descriptor()
}

func (Status) Type() protoreflect.EnumType {
	return &file_pkg_proto_multiverse_multiverse_proto_enumTypes[0]
}

func (x Status) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Status.Descriptor instead.
func (Status) EnumDescriptor() ([]byte, []int) {
	return file_pkg_proto_multiverse_multiverse_proto_rawDescGZIP(), []int{0}
}

type GameEvent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uid  string `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`
	Data []byte `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *GameEvent) Reset() {
	*x = GameEvent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_proto_multiverse_multiverse_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GameEvent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GameEvent) ProtoMessage() {}

func (x *GameEvent) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_proto_multiverse_multiverse_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GameEvent.ProtoReflect.Descriptor instead.
func (*GameEvent) Descriptor() ([]byte, []int) {
	return file_pkg_proto_multiverse_multiverse_proto_rawDescGZIP(), []int{0}
}

func (x *GameEvent) GetUid() string {
	if x != nil {
		return x.Uid
	}
	return ""
}

func (x *GameEvent) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

type Filter struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uid        string   `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`
	Recipients []string `protobuf:"bytes,2,rep,name=recipients,proto3" json:"recipients,omitempty"`
}

func (x *Filter) Reset() {
	*x = Filter{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_proto_multiverse_multiverse_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Filter) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Filter) ProtoMessage() {}

func (x *Filter) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_proto_multiverse_multiverse_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Filter.ProtoReflect.Descriptor instead.
func (*Filter) Descriptor() ([]byte, []int) {
	return file_pkg_proto_multiverse_multiverse_proto_rawDescGZIP(), []int{1}
}

func (x *Filter) GetUid() string {
	if x != nil {
		return x.Uid
	}
	return ""
}

func (x *Filter) GetRecipients() []string {
	if x != nil {
		return x.Recipients
	}
	return nil
}

type Event struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Type:
	//
	//	*Event_CharacterDeath
	Type isEvent_Type `protobuf_oneof:"type"`
}

func (x *Event) Reset() {
	*x = Event{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_proto_multiverse_multiverse_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Event) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Event) ProtoMessage() {}

func (x *Event) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_proto_multiverse_multiverse_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Event.ProtoReflect.Descriptor instead.
func (*Event) Descriptor() ([]byte, []int) {
	return file_pkg_proto_multiverse_multiverse_proto_rawDescGZIP(), []int{2}
}

func (m *Event) GetType() isEvent_Type {
	if m != nil {
		return m.Type
	}
	return nil
}

func (x *Event) GetCharacterDeath() *DeathByCharacterType {
	if x, ok := x.GetType().(*Event_CharacterDeath); ok {
		return x.CharacterDeath
	}
	return nil
}

type isEvent_Type interface {
	isEvent_Type()
}

type Event_CharacterDeath struct {
	CharacterDeath *DeathByCharacterType `protobuf:"bytes,1,opt,name=character_death,json=characterDeath,proto3,oneof"`
}

func (*Event_CharacterDeath) isEvent_Type() {}

type DeathByCharacterType struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type uint64 `protobuf:"varint,1,opt,name=type,proto3" json:"type,omitempty"`
}

func (x *DeathByCharacterType) Reset() {
	*x = DeathByCharacterType{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_proto_multiverse_multiverse_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeathByCharacterType) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeathByCharacterType) ProtoMessage() {}

func (x *DeathByCharacterType) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_proto_multiverse_multiverse_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeathByCharacterType.ProtoReflect.Descriptor instead.
func (*DeathByCharacterType) Descriptor() ([]byte, []int) {
	return file_pkg_proto_multiverse_multiverse_proto_rawDescGZIP(), []int{3}
}

func (x *DeathByCharacterType) GetType() uint64 {
	if x != nil {
		return x.Type
	}
	return 0
}

type Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status  Status `protobuf:"varint,1,opt,name=status,proto3,enum=multiverse.Status" json:"status,omitempty"`
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *Response) Reset() {
	*x = Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_proto_multiverse_multiverse_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Response) ProtoMessage() {}

func (x *Response) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_proto_multiverse_multiverse_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Response.ProtoReflect.Descriptor instead.
func (*Response) Descriptor() ([]byte, []int) {
	return file_pkg_proto_multiverse_multiverse_proto_rawDescGZIP(), []int{4}
}

func (x *Response) GetStatus() Status {
	if x != nil {
		return x.Status
	}
	return Status_OK
}

func (x *Response) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_pkg_proto_multiverse_multiverse_proto protoreflect.FileDescriptor

var file_pkg_proto_multiverse_multiverse_proto_rawDesc = []byte{
	0x0a, 0x25, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6d, 0x75, 0x6c, 0x74,
	0x69, 0x76, 0x65, 0x72, 0x73, 0x65, 0x2f, 0x6d, 0x75, 0x6c, 0x74, 0x69, 0x76, 0x65, 0x72, 0x73,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x6d, 0x75, 0x6c, 0x74, 0x69, 0x76, 0x65,
	0x72, 0x73, 0x65, 0x22, 0x31, 0x0a, 0x09, 0x47, 0x61, 0x6d, 0x65, 0x45, 0x76, 0x65, 0x6e, 0x74,
	0x12, 0x10, 0x0a, 0x03, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75,
	0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c,
	0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x3a, 0x0a, 0x06, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72,
	0x12, 0x10, 0x0a, 0x03, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75,
	0x69, 0x64, 0x12, 0x1e, 0x0a, 0x0a, 0x72, 0x65, 0x63, 0x69, 0x70, 0x69, 0x65, 0x6e, 0x74, 0x73,
	0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0a, 0x72, 0x65, 0x63, 0x69, 0x70, 0x69, 0x65, 0x6e,
	0x74, 0x73, 0x22, 0x5c, 0x0a, 0x05, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x4b, 0x0a, 0x0f, 0x63,
	0x68, 0x61, 0x72, 0x61, 0x63, 0x74, 0x65, 0x72, 0x5f, 0x64, 0x65, 0x61, 0x74, 0x68, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x6d, 0x75, 0x6c, 0x74, 0x69, 0x76, 0x65, 0x72, 0x73,
	0x65, 0x2e, 0x44, 0x65, 0x61, 0x74, 0x68, 0x42, 0x79, 0x43, 0x68, 0x61, 0x72, 0x61, 0x63, 0x74,
	0x65, 0x72, 0x54, 0x79, 0x70, 0x65, 0x48, 0x00, 0x52, 0x0e, 0x63, 0x68, 0x61, 0x72, 0x61, 0x63,
	0x74, 0x65, 0x72, 0x44, 0x65, 0x61, 0x74, 0x68, 0x42, 0x06, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65,
	0x22, 0x2a, 0x0a, 0x14, 0x44, 0x65, 0x61, 0x74, 0x68, 0x42, 0x79, 0x43, 0x68, 0x61, 0x72, 0x61,
	0x63, 0x74, 0x65, 0x72, 0x54, 0x79, 0x70, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x22, 0x50, 0x0a, 0x08,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2a, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x12, 0x2e, 0x6d, 0x75, 0x6c, 0x74, 0x69,
	0x76, 0x65, 0x72, 0x73, 0x65, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2a, 0x1d,
	0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x06, 0x0a, 0x02, 0x4f, 0x4b, 0x10, 0x00,
	0x12, 0x0b, 0x0a, 0x07, 0x46, 0x61, 0x69, 0x6c, 0x75, 0x72, 0x65, 0x10, 0x01, 0x32, 0x84, 0x01,
	0x0a, 0x0a, 0x4d, 0x75, 0x6c, 0x74, 0x69, 0x76, 0x65, 0x72, 0x73, 0x65, 0x12, 0x41, 0x0a, 0x10,
	0x50, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x47, 0x61, 0x6d, 0x65, 0x45, 0x76, 0x65, 0x6e, 0x74,
	0x12, 0x15, 0x2e, 0x6d, 0x75, 0x6c, 0x74, 0x69, 0x76, 0x65, 0x72, 0x73, 0x65, 0x2e, 0x47, 0x61,
	0x6d, 0x65, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x1a, 0x14, 0x2e, 0x6d, 0x75, 0x6c, 0x74, 0x69, 0x76,
	0x65, 0x72, 0x73, 0x65, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12,
	0x33, 0x0a, 0x06, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x12, 0x2e, 0x6d, 0x75, 0x6c, 0x74,
	0x69, 0x76, 0x65, 0x72, 0x73, 0x65, 0x2e, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x1a, 0x11, 0x2e,
	0x6d, 0x75, 0x6c, 0x74, 0x69, 0x76, 0x65, 0x72, 0x73, 0x65, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74,
	0x22, 0x00, 0x30, 0x01, 0x42, 0x3a, 0x5a, 0x38, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x63, 0x69, 0x70, 0x68, 0x65, 0x72, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x61, 0x69,
	0x6e, 0x2f, 0x64, 0x65, 0x61, 0x64, 0x65, 0x6e, 0x7a, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x6d, 0x75,
	0x6c, 0x74, 0x69, 0x76, 0x65, 0x72, 0x73, 0x65, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pkg_proto_multiverse_multiverse_proto_rawDescOnce sync.Once
	file_pkg_proto_multiverse_multiverse_proto_rawDescData = file_pkg_proto_multiverse_multiverse_proto_rawDesc
)

func file_pkg_proto_multiverse_multiverse_proto_rawDescGZIP() []byte {
	file_pkg_proto_multiverse_multiverse_proto_rawDescOnce.Do(func() {
		file_pkg_proto_multiverse_multiverse_proto_rawDescData = protoimpl.X.CompressGZIP(file_pkg_proto_multiverse_multiverse_proto_rawDescData)
	})
	return file_pkg_proto_multiverse_multiverse_proto_rawDescData
}

var file_pkg_proto_multiverse_multiverse_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_pkg_proto_multiverse_multiverse_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_pkg_proto_multiverse_multiverse_proto_goTypes = []interface{}{
	(Status)(0),                  // 0: multiverse.Status
	(*GameEvent)(nil),            // 1: multiverse.GameEvent
	(*Filter)(nil),               // 2: multiverse.Filter
	(*Event)(nil),                // 3: multiverse.Event
	(*DeathByCharacterType)(nil), // 4: multiverse.DeathByCharacterType
	(*Response)(nil),             // 5: multiverse.Response
}
var file_pkg_proto_multiverse_multiverse_proto_depIdxs = []int32{
	4, // 0: multiverse.Event.character_death:type_name -> multiverse.DeathByCharacterType
	0, // 1: multiverse.Response.status:type_name -> multiverse.Status
	1, // 2: multiverse.Multiverse.PublishGameEvent:input_type -> multiverse.GameEvent
	2, // 3: multiverse.Multiverse.Events:input_type -> multiverse.Filter
	5, // 4: multiverse.Multiverse.PublishGameEvent:output_type -> multiverse.Response
	3, // 5: multiverse.Multiverse.Events:output_type -> multiverse.Event
	4, // [4:6] is the sub-list for method output_type
	2, // [2:4] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_pkg_proto_multiverse_multiverse_proto_init() }
func file_pkg_proto_multiverse_multiverse_proto_init() {
	if File_pkg_proto_multiverse_multiverse_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pkg_proto_multiverse_multiverse_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GameEvent); i {
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
		file_pkg_proto_multiverse_multiverse_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Filter); i {
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
		file_pkg_proto_multiverse_multiverse_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Event); i {
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
		file_pkg_proto_multiverse_multiverse_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeathByCharacterType); i {
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
		file_pkg_proto_multiverse_multiverse_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Response); i {
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
	file_pkg_proto_multiverse_multiverse_proto_msgTypes[2].OneofWrappers = []interface{}{
		(*Event_CharacterDeath)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_pkg_proto_multiverse_multiverse_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_pkg_proto_multiverse_multiverse_proto_goTypes,
		DependencyIndexes: file_pkg_proto_multiverse_multiverse_proto_depIdxs,
		EnumInfos:         file_pkg_proto_multiverse_multiverse_proto_enumTypes,
		MessageInfos:      file_pkg_proto_multiverse_multiverse_proto_msgTypes,
	}.Build()
	File_pkg_proto_multiverse_multiverse_proto = out.File
	file_pkg_proto_multiverse_multiverse_proto_rawDesc = nil
	file_pkg_proto_multiverse_multiverse_proto_goTypes = nil
	file_pkg_proto_multiverse_multiverse_proto_depIdxs = nil
}
