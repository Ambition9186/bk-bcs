// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.22.0
// source: commit.proto

package pbcommit

import (
	base "bscp.io/pkg/protocol/core/base"
	content "bscp.io/pkg/protocol/core/content"
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

// Commit source resource reference: pkg/dal/table/commit.go
type Commit struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         uint32                `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Spec       *CommitSpec           `protobuf:"bytes,2,opt,name=spec,proto3" json:"spec,omitempty"`
	Attachment *CommitAttachment     `protobuf:"bytes,3,opt,name=attachment,proto3" json:"attachment,omitempty"`
	Revision   *base.CreatedRevision `protobuf:"bytes,4,opt,name=revision,proto3" json:"revision,omitempty"`
}

func (x *Commit) Reset() {
	*x = Commit{}
	if protoimpl.UnsafeEnabled {
		mi := &file_commit_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Commit) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Commit) ProtoMessage() {}

func (x *Commit) ProtoReflect() protoreflect.Message {
	mi := &file_commit_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Commit.ProtoReflect.Descriptor instead.
func (*Commit) Descriptor() ([]byte, []int) {
	return file_commit_proto_rawDescGZIP(), []int{0}
}

func (x *Commit) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Commit) GetSpec() *CommitSpec {
	if x != nil {
		return x.Spec
	}
	return nil
}

func (x *Commit) GetAttachment() *CommitAttachment {
	if x != nil {
		return x.Attachment
	}
	return nil
}

func (x *Commit) GetRevision() *base.CreatedRevision {
	if x != nil {
		return x.Revision
	}
	return nil
}

// CommitSpec source resource reference: pkg/dal/table/commit.go
type CommitSpec struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ContentId uint32               `protobuf:"varint,1,opt,name=content_id,json=contentId,proto3" json:"content_id,omitempty"`
	Content   *content.ContentSpec `protobuf:"bytes,2,opt,name=content,proto3" json:"content,omitempty"`
	Memo      string               `protobuf:"bytes,3,opt,name=memo,proto3" json:"memo,omitempty"`
}

func (x *CommitSpec) Reset() {
	*x = CommitSpec{}
	if protoimpl.UnsafeEnabled {
		mi := &file_commit_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CommitSpec) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CommitSpec) ProtoMessage() {}

func (x *CommitSpec) ProtoReflect() protoreflect.Message {
	mi := &file_commit_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CommitSpec.ProtoReflect.Descriptor instead.
func (*CommitSpec) Descriptor() ([]byte, []int) {
	return file_commit_proto_rawDescGZIP(), []int{1}
}

func (x *CommitSpec) GetContentId() uint32 {
	if x != nil {
		return x.ContentId
	}
	return 0
}

func (x *CommitSpec) GetContent() *content.ContentSpec {
	if x != nil {
		return x.Content
	}
	return nil
}

func (x *CommitSpec) GetMemo() string {
	if x != nil {
		return x.Memo
	}
	return ""
}

// CommitAttachment source resource reference: pkg/dal/table/commit.go
type CommitAttachment struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BizId        uint32 `protobuf:"varint,1,opt,name=biz_id,json=bizId,proto3" json:"biz_id,omitempty"`
	AppId        uint32 `protobuf:"varint,2,opt,name=app_id,json=appId,proto3" json:"app_id,omitempty"`
	ConfigItemId uint32 `protobuf:"varint,3,opt,name=config_item_id,json=configItemId,proto3" json:"config_item_id,omitempty"`
}

func (x *CommitAttachment) Reset() {
	*x = CommitAttachment{}
	if protoimpl.UnsafeEnabled {
		mi := &file_commit_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CommitAttachment) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CommitAttachment) ProtoMessage() {}

func (x *CommitAttachment) ProtoReflect() protoreflect.Message {
	mi := &file_commit_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CommitAttachment.ProtoReflect.Descriptor instead.
func (*CommitAttachment) Descriptor() ([]byte, []int) {
	return file_commit_proto_rawDescGZIP(), []int{2}
}

func (x *CommitAttachment) GetBizId() uint32 {
	if x != nil {
		return x.BizId
	}
	return 0
}

func (x *CommitAttachment) GetAppId() uint32 {
	if x != nil {
		return x.AppId
	}
	return 0
}

func (x *CommitAttachment) GetConfigItemId() uint32 {
	if x != nil {
		return x.ConfigItemId
	}
	return 0
}

var File_commit_proto protoreflect.FileDescriptor

var file_commit_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x63, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08,
	0x70, 0x62, 0x63, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x1a, 0x29, 0x62, 0x73, 0x63, 0x70, 0x2e, 0x69,
	0x6f, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2f, 0x63,
	0x6f, 0x72, 0x65, 0x2f, 0x62, 0x61, 0x73, 0x65, 0x2f, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x2f, 0x62, 0x73, 0x63, 0x70, 0x2e, 0x69, 0x6f, 0x2f, 0x70, 0x6b, 0x67,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x63,
	0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x2f, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0xb3, 0x01, 0x0a, 0x06, 0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x02, 0x69, 0x64, 0x12,
	0x28, 0x0a, 0x04, 0x73, 0x70, 0x65, 0x63, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e,
	0x70, 0x62, 0x63, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x53,
	0x70, 0x65, 0x63, 0x52, 0x04, 0x73, 0x70, 0x65, 0x63, 0x12, 0x3a, 0x0a, 0x0a, 0x61, 0x74, 0x74,
	0x61, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e,
	0x70, 0x62, 0x63, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x41,
	0x74, 0x74, 0x61, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x0a, 0x61, 0x74, 0x74, 0x61, 0x63,
	0x68, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x33, 0x0a, 0x08, 0x72, 0x65, 0x76, 0x69, 0x73, 0x69, 0x6f,
	0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x70, 0x62, 0x62, 0x61, 0x73, 0x65,
	0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x52, 0x65, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e,
	0x52, 0x08, 0x72, 0x65, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x22, 0x71, 0x0a, 0x0a, 0x43, 0x6f,
	0x6d, 0x6d, 0x69, 0x74, 0x53, 0x70, 0x65, 0x63, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x6f, 0x6e, 0x74,
	0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x09, 0x63, 0x6f,
	0x6e, 0x74, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x30, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65,
	0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x70, 0x62, 0x63, 0x6f, 0x6e,
	0x74, 0x65, 0x6e, 0x74, 0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x53, 0x70, 0x65, 0x63,
	0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6d, 0x65, 0x6d,
	0x6f, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6d, 0x65, 0x6d, 0x6f, 0x22, 0x66, 0x0a,
	0x10, 0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x41, 0x74, 0x74, 0x61, 0x63, 0x68, 0x6d, 0x65, 0x6e,
	0x74, 0x12, 0x15, 0x0a, 0x06, 0x62, 0x69, 0x7a, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x05, 0x62, 0x69, 0x7a, 0x49, 0x64, 0x12, 0x15, 0x0a, 0x06, 0x61, 0x70, 0x70, 0x5f,
	0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x61, 0x70, 0x70, 0x49, 0x64, 0x12,
	0x24, 0x0a, 0x0e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x5f, 0x69, 0x74, 0x65, 0x6d, 0x5f, 0x69,
	0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0c, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x49,
	0x74, 0x65, 0x6d, 0x49, 0x64, 0x42, 0x2b, 0x5a, 0x29, 0x62, 0x73, 0x63, 0x70, 0x2e, 0x69, 0x6f,
	0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2f, 0x63, 0x6f,
	0x72, 0x65, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x3b, 0x70, 0x62, 0x63, 0x6f, 0x6d, 0x6d,
	0x69, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_commit_proto_rawDescOnce sync.Once
	file_commit_proto_rawDescData = file_commit_proto_rawDesc
)

func file_commit_proto_rawDescGZIP() []byte {
	file_commit_proto_rawDescOnce.Do(func() {
		file_commit_proto_rawDescData = protoimpl.X.CompressGZIP(file_commit_proto_rawDescData)
	})
	return file_commit_proto_rawDescData
}

var file_commit_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_commit_proto_goTypes = []interface{}{
	(*Commit)(nil),               // 0: pbcommit.Commit
	(*CommitSpec)(nil),           // 1: pbcommit.CommitSpec
	(*CommitAttachment)(nil),     // 2: pbcommit.CommitAttachment
	(*base.CreatedRevision)(nil), // 3: pbbase.CreatedRevision
	(*content.ContentSpec)(nil),  // 4: pbcontent.ContentSpec
}
var file_commit_proto_depIdxs = []int32{
	1, // 0: pbcommit.Commit.spec:type_name -> pbcommit.CommitSpec
	2, // 1: pbcommit.Commit.attachment:type_name -> pbcommit.CommitAttachment
	3, // 2: pbcommit.Commit.revision:type_name -> pbbase.CreatedRevision
	4, // 3: pbcommit.CommitSpec.content:type_name -> pbcontent.ContentSpec
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_commit_proto_init() }
func file_commit_proto_init() {
	if File_commit_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_commit_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Commit); i {
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
		file_commit_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CommitSpec); i {
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
		file_commit_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CommitAttachment); i {
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
			RawDescriptor: file_commit_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_commit_proto_goTypes,
		DependencyIndexes: file_commit_proto_depIdxs,
		MessageInfos:      file_commit_proto_msgTypes,
	}.Build()
	File_commit_proto = out.File
	file_commit_proto_rawDesc = nil
	file_commit_proto_goTypes = nil
	file_commit_proto_depIdxs = nil
}
