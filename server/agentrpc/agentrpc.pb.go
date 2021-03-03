// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.14.0
// source: agentrpc.proto

package agentrpc

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

type CloneRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token     string `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	External  bool   `protobuf:"varint,2,opt,name=external,proto3" json:"external,omitempty"`
	DonorHost string `protobuf:"bytes,3,opt,name=donor_host,json=donorHost,proto3" json:"donor_host,omitempty"`
	DonorPort int32  `protobuf:"varint,4,opt,name=donor_port,json=donorPort,proto3" json:"donor_port,omitempty"`
}

func (x *CloneRequest) Reset() {
	*x = CloneRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_agentrpc_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CloneRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CloneRequest) ProtoMessage() {}

func (x *CloneRequest) ProtoReflect() protoreflect.Message {
	mi := &file_agentrpc_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CloneRequest.ProtoReflect.Descriptor instead.
func (*CloneRequest) Descriptor() ([]byte, []int) {
	return file_agentrpc_proto_rawDescGZIP(), []int{0}
}

func (x *CloneRequest) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *CloneRequest) GetExternal() bool {
	if x != nil {
		return x.External
	}
	return false
}

func (x *CloneRequest) GetDonorHost() string {
	if x != nil {
		return x.DonorHost
	}
	return ""
}

func (x *CloneRequest) GetDonorPort() int32 {
	if x != nil {
		return x.DonorPort
	}
	return 0
}

type CloneResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *CloneResponse) Reset() {
	*x = CloneResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_agentrpc_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CloneResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CloneResponse) ProtoMessage() {}

func (x *CloneResponse) ProtoReflect() protoreflect.Message {
	mi := &file_agentrpc_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CloneResponse.ProtoReflect.Descriptor instead.
func (*CloneResponse) Descriptor() ([]byte, []int) {
	return file_agentrpc_proto_rawDescGZIP(), []int{1}
}

type FlushAndBackupBinlogRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token           string `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	BackupId        string `protobuf:"bytes,2,opt,name=backup_id,json=backupId,proto3" json:"backup_id,omitempty"`
	BucketHost      string `protobuf:"bytes,3,opt,name=bucket_host,json=bucketHost,proto3" json:"bucket_host,omitempty"`
	BucketPort      int32  `protobuf:"varint,4,opt,name=bucket_port,json=bucketPort,proto3" json:"bucket_port,omitempty"`
	BucketName      string `protobuf:"bytes,5,opt,name=bucket_name,json=bucketName,proto3" json:"bucket_name,omitempty"`
	BucketRegion    string `protobuf:"bytes,6,opt,name=bucket_region,json=bucketRegion,proto3" json:"bucket_region,omitempty"`
	AccessKeyId     string `protobuf:"bytes,7,opt,name=access_key_id,json=accessKeyId,proto3" json:"access_key_id,omitempty"`
	SecretAccessKey string `protobuf:"bytes,8,opt,name=secret_access_key,json=secretAccessKey,proto3" json:"secret_access_key,omitempty"`
}

func (x *FlushAndBackupBinlogRequest) Reset() {
	*x = FlushAndBackupBinlogRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_agentrpc_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FlushAndBackupBinlogRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FlushAndBackupBinlogRequest) ProtoMessage() {}

func (x *FlushAndBackupBinlogRequest) ProtoReflect() protoreflect.Message {
	mi := &file_agentrpc_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FlushAndBackupBinlogRequest.ProtoReflect.Descriptor instead.
func (*FlushAndBackupBinlogRequest) Descriptor() ([]byte, []int) {
	return file_agentrpc_proto_rawDescGZIP(), []int{2}
}

func (x *FlushAndBackupBinlogRequest) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *FlushAndBackupBinlogRequest) GetBackupId() string {
	if x != nil {
		return x.BackupId
	}
	return ""
}

func (x *FlushAndBackupBinlogRequest) GetBucketHost() string {
	if x != nil {
		return x.BucketHost
	}
	return ""
}

func (x *FlushAndBackupBinlogRequest) GetBucketPort() int32 {
	if x != nil {
		return x.BucketPort
	}
	return 0
}

func (x *FlushAndBackupBinlogRequest) GetBucketName() string {
	if x != nil {
		return x.BucketName
	}
	return ""
}

func (x *FlushAndBackupBinlogRequest) GetBucketRegion() string {
	if x != nil {
		return x.BucketRegion
	}
	return ""
}

func (x *FlushAndBackupBinlogRequest) GetAccessKeyId() string {
	if x != nil {
		return x.AccessKeyId
	}
	return ""
}

func (x *FlushAndBackupBinlogRequest) GetSecretAccessKey() string {
	if x != nil {
		return x.SecretAccessKey
	}
	return ""
}

type FlushAndBackupBinlogResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *FlushAndBackupBinlogResponse) Reset() {
	*x = FlushAndBackupBinlogResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_agentrpc_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FlushAndBackupBinlogResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FlushAndBackupBinlogResponse) ProtoMessage() {}

func (x *FlushAndBackupBinlogResponse) ProtoReflect() protoreflect.Message {
	mi := &file_agentrpc_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FlushAndBackupBinlogResponse.ProtoReflect.Descriptor instead.
func (*FlushAndBackupBinlogResponse) Descriptor() ([]byte, []int) {
	return file_agentrpc_proto_rawDescGZIP(), []int{3}
}

type FlushBinlogRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token  string `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	Delete bool   `protobuf:"varint,2,opt,name=delete,proto3" json:"delete,omitempty"`
}

func (x *FlushBinlogRequest) Reset() {
	*x = FlushBinlogRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_agentrpc_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FlushBinlogRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FlushBinlogRequest) ProtoMessage() {}

func (x *FlushBinlogRequest) ProtoReflect() protoreflect.Message {
	mi := &file_agentrpc_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FlushBinlogRequest.ProtoReflect.Descriptor instead.
func (*FlushBinlogRequest) Descriptor() ([]byte, []int) {
	return file_agentrpc_proto_rawDescGZIP(), []int{4}
}

func (x *FlushBinlogRequest) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *FlushBinlogRequest) GetDelete() bool {
	if x != nil {
		return x.Delete
	}
	return false
}

type FlushBinlogResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *FlushBinlogResponse) Reset() {
	*x = FlushBinlogResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_agentrpc_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FlushBinlogResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FlushBinlogResponse) ProtoMessage() {}

func (x *FlushBinlogResponse) ProtoReflect() protoreflect.Message {
	mi := &file_agentrpc_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FlushBinlogResponse.ProtoReflect.Descriptor instead.
func (*FlushBinlogResponse) Descriptor() ([]byte, []int) {
	return file_agentrpc_proto_rawDescGZIP(), []int{5}
}

var File_agentrpc_proto protoreflect.FileDescriptor

var file_agentrpc_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x72, 0x70, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x7e, 0x0a, 0x0c, 0x43, 0x6c, 0x6f, 0x6e, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x1a, 0x0a, 0x08, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e,
	0x61, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e,
	0x61, 0x6c, 0x12, 0x1d, 0x0a, 0x0a, 0x64, 0x6f, 0x6e, 0x6f, 0x72, 0x5f, 0x68, 0x6f, 0x73, 0x74,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x64, 0x6f, 0x6e, 0x6f, 0x72, 0x48, 0x6f, 0x73,
	0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x64, 0x6f, 0x6e, 0x6f, 0x72, 0x5f, 0x70, 0x6f, 0x72, 0x74, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x64, 0x6f, 0x6e, 0x6f, 0x72, 0x50, 0x6f, 0x72, 0x74,
	0x22, 0x0f, 0x0a, 0x0d, 0x43, 0x6c, 0x6f, 0x6e, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0xa8, 0x02, 0x0a, 0x1b, 0x46, 0x6c, 0x75, 0x73, 0x68, 0x41, 0x6e, 0x64, 0x42, 0x61,
	0x63, 0x6b, 0x75, 0x70, 0x42, 0x69, 0x6e, 0x6c, 0x6f, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x1b, 0x0a, 0x09, 0x62, 0x61, 0x63, 0x6b, 0x75,
	0x70, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x62, 0x61, 0x63, 0x6b,
	0x75, 0x70, 0x49, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x62, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x5f, 0x68,
	0x6f, 0x73, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x62, 0x75, 0x63, 0x6b, 0x65,
	0x74, 0x48, 0x6f, 0x73, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x62, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x5f,
	0x70, 0x6f, 0x72, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x62, 0x75, 0x63, 0x6b,
	0x65, 0x74, 0x50, 0x6f, 0x72, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x62, 0x75, 0x63, 0x6b, 0x65, 0x74,
	0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x62, 0x75, 0x63,
	0x6b, 0x65, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x62, 0x75, 0x63, 0x6b, 0x65,
	0x74, 0x5f, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c,
	0x62, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x52, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x12, 0x22, 0x0a, 0x0d,
	0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x5f, 0x6b, 0x65, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x07, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0b, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x4b, 0x65, 0x79, 0x49, 0x64,
	0x12, 0x2a, 0x0a, 0x11, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x5f, 0x61, 0x63, 0x63, 0x65, 0x73,
	0x73, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x73, 0x65, 0x63,
	0x72, 0x65, 0x74, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x4b, 0x65, 0x79, 0x22, 0x1e, 0x0a, 0x1c,
	0x46, 0x6c, 0x75, 0x73, 0x68, 0x41, 0x6e, 0x64, 0x42, 0x61, 0x63, 0x6b, 0x75, 0x70, 0x42, 0x69,
	0x6e, 0x6c, 0x6f, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x42, 0x0a, 0x12,
	0x46, 0x6c, 0x75, 0x73, 0x68, 0x42, 0x69, 0x6e, 0x6c, 0x6f, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x64, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x22, 0x15, 0x0a, 0x13, 0x46, 0x6c, 0x75, 0x73, 0x68, 0x42, 0x69, 0x6e, 0x6c, 0x6f, 0x67, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x32, 0x36, 0x0a, 0x0c, 0x43, 0x6c, 0x6f, 0x6e, 0x65,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x26, 0x0a, 0x05, 0x43, 0x6c, 0x6f, 0x6e, 0x65,
	0x12, 0x0d, 0x2e, 0x43, 0x6c, 0x6f, 0x6e, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x0e, 0x2e, 0x43, 0x6c, 0x6f, 0x6e, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x32,
	0xa4, 0x01, 0x0a, 0x13, 0x42, 0x61, 0x63, 0x6b, 0x75, 0x70, 0x42, 0x69, 0x6e, 0x6c, 0x6f, 0x67,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x53, 0x0a, 0x14, 0x46, 0x6c, 0x75, 0x73, 0x68,
	0x41, 0x6e, 0x64, 0x42, 0x61, 0x63, 0x6b, 0x75, 0x70, 0x42, 0x69, 0x6e, 0x6c, 0x6f, 0x67, 0x12,
	0x1c, 0x2e, 0x46, 0x6c, 0x75, 0x73, 0x68, 0x41, 0x6e, 0x64, 0x42, 0x61, 0x63, 0x6b, 0x75, 0x70,
	0x42, 0x69, 0x6e, 0x6c, 0x6f, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e,
	0x46, 0x6c, 0x75, 0x73, 0x68, 0x41, 0x6e, 0x64, 0x42, 0x61, 0x63, 0x6b, 0x75, 0x70, 0x42, 0x69,
	0x6e, 0x6c, 0x6f, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x38, 0x0a, 0x0b,
	0x46, 0x6c, 0x75, 0x73, 0x68, 0x42, 0x69, 0x6e, 0x6c, 0x6f, 0x67, 0x12, 0x13, 0x2e, 0x46, 0x6c,
	0x75, 0x73, 0x68, 0x42, 0x69, 0x6e, 0x6c, 0x6f, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x14, 0x2e, 0x46, 0x6c, 0x75, 0x73, 0x68, 0x42, 0x69, 0x6e, 0x6c, 0x6f, 0x67, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x31, 0x5a, 0x2f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x79, 0x62, 0x6f, 0x7a, 0x75, 0x2d, 0x67, 0x6f, 0x2f, 0x6d,
	0x6f, 0x63, 0x6f, 0x2d, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x2f, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x72, 0x70, 0x63, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_agentrpc_proto_rawDescOnce sync.Once
	file_agentrpc_proto_rawDescData = file_agentrpc_proto_rawDesc
)

func file_agentrpc_proto_rawDescGZIP() []byte {
	file_agentrpc_proto_rawDescOnce.Do(func() {
		file_agentrpc_proto_rawDescData = protoimpl.X.CompressGZIP(file_agentrpc_proto_rawDescData)
	})
	return file_agentrpc_proto_rawDescData
}

var file_agentrpc_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_agentrpc_proto_goTypes = []interface{}{
	(*CloneRequest)(nil),                 // 0: CloneRequest
	(*CloneResponse)(nil),                // 1: CloneResponse
	(*FlushAndBackupBinlogRequest)(nil),  // 2: FlushAndBackupBinlogRequest
	(*FlushAndBackupBinlogResponse)(nil), // 3: FlushAndBackupBinlogResponse
	(*FlushBinlogRequest)(nil),           // 4: FlushBinlogRequest
	(*FlushBinlogResponse)(nil),          // 5: FlushBinlogResponse
}
var file_agentrpc_proto_depIdxs = []int32{
	0, // 0: CloneService.Clone:input_type -> CloneRequest
	2, // 1: BackupBinlogService.FlushAndBackupBinlog:input_type -> FlushAndBackupBinlogRequest
	4, // 2: BackupBinlogService.FlushBinlog:input_type -> FlushBinlogRequest
	1, // 3: CloneService.Clone:output_type -> CloneResponse
	3, // 4: BackupBinlogService.FlushAndBackupBinlog:output_type -> FlushAndBackupBinlogResponse
	5, // 5: BackupBinlogService.FlushBinlog:output_type -> FlushBinlogResponse
	3, // [3:6] is the sub-list for method output_type
	0, // [0:3] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_agentrpc_proto_init() }
func file_agentrpc_proto_init() {
	if File_agentrpc_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_agentrpc_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CloneRequest); i {
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
		file_agentrpc_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CloneResponse); i {
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
		file_agentrpc_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FlushAndBackupBinlogRequest); i {
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
		file_agentrpc_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FlushAndBackupBinlogResponse); i {
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
		file_agentrpc_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FlushBinlogRequest); i {
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
		file_agentrpc_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FlushBinlogResponse); i {
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
			RawDescriptor: file_agentrpc_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   2,
		},
		GoTypes:           file_agentrpc_proto_goTypes,
		DependencyIndexes: file_agentrpc_proto_depIdxs,
		MessageInfos:      file_agentrpc_proto_msgTypes,
	}.Build()
	File_agentrpc_proto = out.File
	file_agentrpc_proto_rawDesc = nil
	file_agentrpc_proto_goTypes = nil
	file_agentrpc_proto_depIdxs = nil
}