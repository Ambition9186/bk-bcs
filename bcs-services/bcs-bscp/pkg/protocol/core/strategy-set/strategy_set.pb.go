// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.22.0
// source: strategy_set.proto

package pbss

import (
	base "bscp.io/pkg/protocol/core/base"
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

// StrategySet source resource reference: pkg/dal/table/strategy_set.go
type StrategySet struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         uint32                 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Spec       *StrategySetSpec       `protobuf:"bytes,2,opt,name=spec,proto3" json:"spec,omitempty"`
	State      *StrategySetState      `protobuf:"bytes,3,opt,name=state,proto3" json:"state,omitempty"`
	Attachment *StrategySetAttachment `protobuf:"bytes,4,opt,name=attachment,proto3" json:"attachment,omitempty"`
	Revision   *base.Revision         `protobuf:"bytes,5,opt,name=revision,proto3" json:"revision,omitempty"`
}

func (x *StrategySet) Reset() {
	*x = StrategySet{}
	if protoimpl.UnsafeEnabled {
		mi := &file_strategy_set_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StrategySet) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StrategySet) ProtoMessage() {}

func (x *StrategySet) ProtoReflect() protoreflect.Message {
	mi := &file_strategy_set_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StrategySet.ProtoReflect.Descriptor instead.
func (*StrategySet) Descriptor() ([]byte, []int) {
	return file_strategy_set_proto_rawDescGZIP(), []int{0}
}

func (x *StrategySet) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *StrategySet) GetSpec() *StrategySetSpec {
	if x != nil {
		return x.Spec
	}
	return nil
}

func (x *StrategySet) GetState() *StrategySetState {
	if x != nil {
		return x.State
	}
	return nil
}

func (x *StrategySet) GetAttachment() *StrategySetAttachment {
	if x != nil {
		return x.Attachment
	}
	return nil
}

func (x *StrategySet) GetRevision() *base.Revision {
	if x != nil {
		return x.Revision
	}
	return nil
}

// StrategySetSpec source resource reference: pkg/dal/table/strategy_set.go
type StrategySetSpec struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Mode string `protobuf:"bytes,2,opt,name=mode,proto3" json:"mode,omitempty"`
	Memo string `protobuf:"bytes,3,opt,name=memo,proto3" json:"memo,omitempty"`
}

func (x *StrategySetSpec) Reset() {
	*x = StrategySetSpec{}
	if protoimpl.UnsafeEnabled {
		mi := &file_strategy_set_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StrategySetSpec) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StrategySetSpec) ProtoMessage() {}

func (x *StrategySetSpec) ProtoReflect() protoreflect.Message {
	mi := &file_strategy_set_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StrategySetSpec.ProtoReflect.Descriptor instead.
func (*StrategySetSpec) Descriptor() ([]byte, []int) {
	return file_strategy_set_proto_rawDescGZIP(), []int{1}
}

func (x *StrategySetSpec) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *StrategySetSpec) GetMode() string {
	if x != nil {
		return x.Mode
	}
	return ""
}

func (x *StrategySetSpec) GetMemo() string {
	if x != nil {
		return x.Memo
	}
	return ""
}

// StrategySetState source resource reference: pkg/dal/table/strategy_set.go
type StrategySetState struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status string `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *StrategySetState) Reset() {
	*x = StrategySetState{}
	if protoimpl.UnsafeEnabled {
		mi := &file_strategy_set_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StrategySetState) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StrategySetState) ProtoMessage() {}

func (x *StrategySetState) ProtoReflect() protoreflect.Message {
	mi := &file_strategy_set_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StrategySetState.ProtoReflect.Descriptor instead.
func (*StrategySetState) Descriptor() ([]byte, []int) {
	return file_strategy_set_proto_rawDescGZIP(), []int{2}
}

func (x *StrategySetState) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

// StrategySetAttachment source resource reference: pkg/dal/table/strategy_set.go
type StrategySetAttachment struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BizId uint32 `protobuf:"varint,1,opt,name=biz_id,json=bizId,proto3" json:"biz_id,omitempty"`
	AppId uint32 `protobuf:"varint,2,opt,name=app_id,json=appId,proto3" json:"app_id,omitempty"`
}

func (x *StrategySetAttachment) Reset() {
	*x = StrategySetAttachment{}
	if protoimpl.UnsafeEnabled {
		mi := &file_strategy_set_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StrategySetAttachment) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StrategySetAttachment) ProtoMessage() {}

func (x *StrategySetAttachment) ProtoReflect() protoreflect.Message {
	mi := &file_strategy_set_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StrategySetAttachment.ProtoReflect.Descriptor instead.
func (*StrategySetAttachment) Descriptor() ([]byte, []int) {
	return file_strategy_set_proto_rawDescGZIP(), []int{3}
}

func (x *StrategySetAttachment) GetBizId() uint32 {
	if x != nil {
		return x.BizId
	}
	return 0
}

func (x *StrategySetAttachment) GetAppId() uint32 {
	if x != nil {
		return x.AppId
	}
	return 0
}

var File_strategy_set_proto protoreflect.FileDescriptor

var file_strategy_set_proto_rawDesc = []byte{
	0x0a, 0x12, 0x73, 0x74, 0x72, 0x61, 0x74, 0x65, 0x67, 0x79, 0x5f, 0x73, 0x65, 0x74, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x70, 0x62, 0x73, 0x73, 0x1a, 0x29, 0x62, 0x73, 0x63, 0x70,
	0x2e, 0x69, 0x6f, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c,
	0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x62, 0x61, 0x73, 0x65, 0x2f, 0x62, 0x61, 0x73, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xe1, 0x01, 0x0a, 0x0b, 0x53, 0x74, 0x72, 0x61, 0x74, 0x65,
	0x67, 0x79, 0x53, 0x65, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x02, 0x69, 0x64, 0x12, 0x29, 0x0a, 0x04, 0x73, 0x70, 0x65, 0x63, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x70, 0x62, 0x73, 0x73, 0x2e, 0x53, 0x74, 0x72, 0x61, 0x74,
	0x65, 0x67, 0x79, 0x53, 0x65, 0x74, 0x53, 0x70, 0x65, 0x63, 0x52, 0x04, 0x73, 0x70, 0x65, 0x63,
	0x12, 0x2c, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x16, 0x2e, 0x70, 0x62, 0x73, 0x73, 0x2e, 0x53, 0x74, 0x72, 0x61, 0x74, 0x65, 0x67, 0x79, 0x53,
	0x65, 0x74, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x12, 0x3b,
	0x0a, 0x0a, 0x61, 0x74, 0x74, 0x61, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x70, 0x62, 0x73, 0x73, 0x2e, 0x53, 0x74, 0x72, 0x61, 0x74, 0x65,
	0x67, 0x79, 0x53, 0x65, 0x74, 0x41, 0x74, 0x74, 0x61, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x52,
	0x0a, 0x61, 0x74, 0x74, 0x61, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x2c, 0x0a, 0x08, 0x72,
	0x65, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e,
	0x70, 0x62, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x52, 0x65, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x52,
	0x08, 0x72, 0x65, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x22, 0x4d, 0x0a, 0x0f, 0x53, 0x74, 0x72,
	0x61, 0x74, 0x65, 0x67, 0x79, 0x53, 0x65, 0x74, 0x53, 0x70, 0x65, 0x63, 0x12, 0x12, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x12, 0x0a, 0x04, 0x6d, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6d, 0x6f, 0x64, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6d, 0x65, 0x6d, 0x6f, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x6d, 0x65, 0x6d, 0x6f, 0x22, 0x2a, 0x0a, 0x10, 0x53, 0x74, 0x72, 0x61,
	0x74, 0x65, 0x67, 0x79, 0x53, 0x65, 0x74, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x16, 0x0a, 0x06,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x22, 0x45, 0x0a, 0x15, 0x53, 0x74, 0x72, 0x61, 0x74, 0x65, 0x67, 0x79,
	0x53, 0x65, 0x74, 0x41, 0x74, 0x74, 0x61, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x15, 0x0a,
	0x06, 0x62, 0x69, 0x7a, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x62,
	0x69, 0x7a, 0x49, 0x64, 0x12, 0x15, 0x0a, 0x06, 0x61, 0x70, 0x70, 0x5f, 0x69, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x61, 0x70, 0x70, 0x49, 0x64, 0x42, 0x2d, 0x5a, 0x2b, 0x62,
	0x73, 0x63, 0x70, 0x2e, 0x69, 0x6f, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x63, 0x6f, 0x6c, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x73, 0x74, 0x72, 0x61, 0x74, 0x65, 0x67,
	0x79, 0x2d, 0x73, 0x65, 0x74, 0x3b, 0x70, 0x62, 0x73, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_strategy_set_proto_rawDescOnce sync.Once
	file_strategy_set_proto_rawDescData = file_strategy_set_proto_rawDesc
)

func file_strategy_set_proto_rawDescGZIP() []byte {
	file_strategy_set_proto_rawDescOnce.Do(func() {
		file_strategy_set_proto_rawDescData = protoimpl.X.CompressGZIP(file_strategy_set_proto_rawDescData)
	})
	return file_strategy_set_proto_rawDescData
}

var file_strategy_set_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_strategy_set_proto_goTypes = []interface{}{
	(*StrategySet)(nil),           // 0: pbss.StrategySet
	(*StrategySetSpec)(nil),       // 1: pbss.StrategySetSpec
	(*StrategySetState)(nil),      // 2: pbss.StrategySetState
	(*StrategySetAttachment)(nil), // 3: pbss.StrategySetAttachment
	(*base.Revision)(nil),         // 4: pbbase.Revision
}
var file_strategy_set_proto_depIdxs = []int32{
	1, // 0: pbss.StrategySet.spec:type_name -> pbss.StrategySetSpec
	2, // 1: pbss.StrategySet.state:type_name -> pbss.StrategySetState
	3, // 2: pbss.StrategySet.attachment:type_name -> pbss.StrategySetAttachment
	4, // 3: pbss.StrategySet.revision:type_name -> pbbase.Revision
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_strategy_set_proto_init() }
func file_strategy_set_proto_init() {
	if File_strategy_set_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_strategy_set_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StrategySet); i {
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
		file_strategy_set_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StrategySetSpec); i {
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
		file_strategy_set_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StrategySetState); i {
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
		file_strategy_set_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StrategySetAttachment); i {
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
			RawDescriptor: file_strategy_set_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_strategy_set_proto_goTypes,
		DependencyIndexes: file_strategy_set_proto_depIdxs,
		MessageInfos:      file_strategy_set_proto_msgTypes,
	}.Build()
	File_strategy_set_proto = out.File
	file_strategy_set_proto_rawDesc = nil
	file_strategy_set_proto_goTypes = nil
	file_strategy_set_proto_depIdxs = nil
}
