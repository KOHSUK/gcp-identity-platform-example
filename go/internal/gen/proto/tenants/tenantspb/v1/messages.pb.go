// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        (unknown)
// source: tenants/tenantspb/v1/messages.proto

package tenantsv1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Tenant struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name          string                 `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Tenant) Reset() {
	*x = Tenant{}
	mi := &file_tenants_tenantspb_v1_messages_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Tenant) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Tenant) ProtoMessage() {}

func (x *Tenant) ProtoReflect() protoreflect.Message {
	mi := &file_tenants_tenantspb_v1_messages_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Tenant.ProtoReflect.Descriptor instead.
func (*Tenant) Descriptor() ([]byte, []int) {
	return file_tenants_tenantspb_v1_messages_proto_rawDescGZIP(), []int{0}
}

func (x *Tenant) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Tenant) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

var File_tenants_tenantspb_v1_messages_proto protoreflect.FileDescriptor

var file_tenants_tenantspb_v1_messages_proto_rawDesc = string([]byte{
	0x0a, 0x23, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x73, 0x2f, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74,
	0x73, 0x70, 0x62, 0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x14, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x73, 0x2e, 0x74,
	0x65, 0x6e, 0x61, 0x6e, 0x74, 0x73, 0x70, 0x62, 0x2e, 0x76, 0x31, 0x22, 0x2c, 0x0a, 0x06, 0x54,
	0x65, 0x6e, 0x61, 0x6e, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x42, 0x37, 0x5a, 0x35, 0x61, 0x70, 0x70,
	0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2f, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x73, 0x2f, 0x74, 0x65, 0x6e, 0x61,
	0x6e, 0x74, 0x73, 0x70, 0x62, 0x2f, 0x76, 0x31, 0x3b, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x73,
	0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_tenants_tenantspb_v1_messages_proto_rawDescOnce sync.Once
	file_tenants_tenantspb_v1_messages_proto_rawDescData []byte
)

func file_tenants_tenantspb_v1_messages_proto_rawDescGZIP() []byte {
	file_tenants_tenantspb_v1_messages_proto_rawDescOnce.Do(func() {
		file_tenants_tenantspb_v1_messages_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_tenants_tenantspb_v1_messages_proto_rawDesc), len(file_tenants_tenantspb_v1_messages_proto_rawDesc)))
	})
	return file_tenants_tenantspb_v1_messages_proto_rawDescData
}

var file_tenants_tenantspb_v1_messages_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_tenants_tenantspb_v1_messages_proto_goTypes = []any{
	(*Tenant)(nil), // 0: tenants.tenantspb.v1.Tenant
}
var file_tenants_tenantspb_v1_messages_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_tenants_tenantspb_v1_messages_proto_init() }
func file_tenants_tenantspb_v1_messages_proto_init() {
	if File_tenants_tenantspb_v1_messages_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_tenants_tenantspb_v1_messages_proto_rawDesc), len(file_tenants_tenantspb_v1_messages_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_tenants_tenantspb_v1_messages_proto_goTypes,
		DependencyIndexes: file_tenants_tenantspb_v1_messages_proto_depIdxs,
		MessageInfos:      file_tenants_tenantspb_v1_messages_proto_msgTypes,
	}.Build()
	File_tenants_tenantspb_v1_messages_proto = out.File
	file_tenants_tenantspb_v1_messages_proto_goTypes = nil
	file_tenants_tenantspb_v1_messages_proto_depIdxs = nil
}
